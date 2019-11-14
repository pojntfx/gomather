//+build mage

package main

import (
	fswatch "github.com/andreaskoch/go-fswatch"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/mholt/archiver"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var gocmd = mg.GoCmd()
var tempdir = os.TempDir()
var protocOut = filepath.Join(tempdir, "usr", "local", "protoc")
var binDir = ".bin"

func ProtocInstallDependencies() error {
	platform := os.Getenv("PLATFORM")
	propPlatform := "linux"
	if platform == "darwin" {
		propPlatform = "osx"
	}
	architecture := os.Getenv("ARCHITECTURE")
	propArchitecture := "x86_64"
	if architecture == "amd64" {
		propArchitecture = "x86_64"
	}

	protocZip := "https://github.com/protocolbuffers/protobuf/releases/download/v3.10.0/protoc-3.10.0-" + propPlatform + "-" + propArchitecture + ".zip"
	protocZipOut := filepath.Join(tempdir, "protoc"+propPlatform+"-"+propArchitecture+".zip")

	res, err := http.Get(protocZip)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = os.RemoveAll(protocZipOut)
	if err != nil {
		return err
	}

	err = os.RemoveAll(protocOut)
	if err != nil {
		return err
	}

	out, err := os.Create(protocZipOut)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	err = archiver.Unarchive(protocZipOut, protocOut)
	if err != nil {
		return err
	}

	sh.RunV(gocmd, "get", "-u", "github.com/golang/protobuf/protoc-gen-go")
	return sh.RunV(gocmd, "get", "-u", "github.com/fiorix/protoc-gen-cobra")
}

func ProtocBuild() error {
	return sh.RunWith(map[string]string{
		"PATH": os.Getenv("PATH") + ":" + filepath.Join(protocOut, "bin"),
	}, gocmd, "generate", "./...")
}

func Build() error {
	return ProtocBuild()
}

func BinaryBuild() error {
	platform := os.Getenv("PLATFORM")
	architecture := os.Getenv("ARCHITECTURE")

	_, err := os.Stat(binDir)
	if os.IsExist(err) {
		os.Mkdir(binDir, 0755)
	}

	return sh.RunWith(map[string]string{
		"CGO_ENABLED": "0",
		"GOOS":        platform,
		"GOARCH":      architecture,
	}, gocmd, "build", "-o", filepath.Join(binDir, "gomather-server-"+platform+"-"+architecture), "github.com/pojntfx/gomather/cmd/gomather-server")
}

func BinaryInstall() error {
	platform := os.Getenv("PLATFORM")
	architecture := os.Getenv("ARCHITECTURE")

	from, _ := os.Open(filepath.Join(binDir, "gomather-server-"+platform+"-"+architecture))
	defer from.Close()

	to, _ := os.OpenFile(filepath.Join("/usr", "local", "bin", "gomather-server"), os.O_RDWR|os.O_CREATE, 755)
	defer to.Close()

	_, err := io.Copy(to, from)

	return err
}

func Clean() {
	binariesToRemove, _ := filepath.Glob("gomather-server-*-*")
	generatedFilesFromProtosToRemove, _ := filepath.Glob(filepath.Join("src", "proto", "generated", "proto", "*"))
	for _, fileToRemove := range append(binariesToRemove, generatedFilesFromProtosToRemove...) {
		os.Remove(fileToRemove)
	}
}

func Run() error {
	return sh.RunV(gocmd, "run", filepath.Join("cmd", "gomather-server", "gomather-server.go"), "start")
}

func Watch() error {
	platform := os.Getenv("PLATFORM")
	architecture := os.Getenv("ARCHITECTURE")

	first := make(chan struct{}, 1)
	var cmd *exec.Cmd
	first <- struct{}{}

	w := fswatch.NewFolderWatcher(".", true, func(path string) bool {
		return strings.HasSuffix(path, ".pb.go") || strings.HasPrefix(path, "gomather-server-")
	}, 1)

	w.Start()
	for w.IsRunning() {
		select {
		case <-first:
		case <-w.ChangeDetails():
		}

		if cmd != nil {
			cmd.Process.Kill()
		}

		BinaryBuild()

		cmd = exec.Command(filepath.Join(binDir, "gomather-server-"+platform+"-"+architecture), "start")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			return err
		}
	}

	return nil
}
