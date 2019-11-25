//+build mage

package main

import (
	fswatch "github.com/andreaskoch/go-fswatch"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/mholt/archiver"
	rz "gitlab.com/z0mbie42/rz-go/v2"
	"gitlab.com/z0mbie42/rz-go/v2/log"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var projectName = "gomather"
var dockerNamespace = "pojntfx"
var gocmd = mg.GoCmd()
var tempdir = os.TempDir()
var protocOut = filepath.Join(tempdir, "usr", "local", "protoc")
var binDir = ".bin"
var installPath = filepath.Join("/usr", "local", "bin", "gomather-server")
var baseProfiles = []string{
	projectName,
	projectName + "-dev",
}
var architectures = []string{
	"amd64",
	"arm64",
}

func ProtocDependencyInstall() error {
	platform := os.Getenv("PLATFORM")
	propPlatform := "linux"
	if platform == "darwin" {
		propPlatform = "osx"
	}
	architecture := os.Getenv("ARCHITECTURE")
	propArchitecture := "x86_64"
	if architecture == "amd64" {
		propArchitecture = "x86_64"
	} else {
		propArchitecture = "aarch_64"
	}

	protocZip := "https://github.com/protocolbuffers/protobuf/releases/download/v3.10.1/protoc-3.10.1-" + propPlatform + "-" + propArchitecture + ".zip"
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

	to, _ := os.OpenFile(installPath, os.O_RDWR|os.O_CREATE, 755)
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

func Start() error {
	return sh.RunV(gocmd, "run", filepath.Join("cmd", "gomather-server", "gomather-server.go"), "start")
}

func UnitTests() error {
	err := sh.RunV(gocmd, "test", "--tags", "unit", "./...")
	if err != nil {
		return err
	}
	log.Info("Passed")
	return nil
}

func IntegrationTests() error {
	err := sh.RunV(gocmd, "install", "./...")
	if err != nil {
		return err
	}

	err = sh.RunV("gomather-server", "--version")
	if err != nil {
		return err
	}

	installPathNonBinary, err := exec.LookPath("gomather-server")
	if err != nil {
		return err
	}
	os.Remove(installPathNonBinary)

	log.Info("Passed")
	return nil
}

func BinaryIntegrationTests() error {
	mg.SerialDeps(BinaryInstall)

	err := sh.RunV("gomather-server", "--version")
	if err != nil {
		return err
	}

	os.Remove(installPath)

	log.Info("Passed")
	return nil
}

func DockerMultiarchSetup() error {
	return sh.RunV("docker", "run", "--rm", "--privileged", "multiarch/qemu-user-static", "--reset", "-p", "yes")
}

func SkaffoldBuild() error {
	mg.SerialDeps(DockerMultiarchSetup)
	var profiles []string

	for _, architecture := range architectures {
		profiles = append(profiles, baseProfiles[0]+"-"+architecture)
	}

	sh.RunV("skaffold", "config", "unset", "--global", "default-repo")

	for _, profile := range profiles {
		sh.RunV("skaffold", "build", "-p", profile)
	}

	return nil
}

func DockerManifestBuild() error {
	var cmds []string

	manifestName := dockerNamespace + "/" + baseProfiles[0] + ":latest"

	cmds = append(cmds, "manifest", "create", "--amend", manifestName)

	for _, architecture := range architectures {
		cmds = append(cmds, dockerNamespace+"/"+baseProfiles[0]+":latest-"+architecture)
	}

	err := sh.RunWith(map[string]string{
		"DOCKER_CLI_EXPERIMENTAL": "enabled",
	}, "docker", cmds...)
	if err != nil {
		return err
	}

	return sh.RunV("docker", "manifest", "push", manifestName)
}

func watch(command []string, deps []interface{}) error {
	w := fswatch.NewFolderWatcher(".", true, func(path string) bool {
		return strings.HasSuffix(path, ".pb.go") || strings.HasPrefix(path, filepath.Join(".bin", "gomather-server-"))
	}, 1)
	w.Start()
	first := make(chan struct{}, 1)
	first <- struct{}{}

	var cmd *exec.Cmd

	for w.IsRunning() {
		select {
		case <-first:
		case <-w.ChangeDetails():
		}

		if cmd != nil {
			cmd.Process.Kill()
		}

		errorInDeps := false

		for _, dep := range deps {
			err := dep.(func() error)()
			if err != nil {
				log.Error("Error while running dependency", rz.Err(err))
				errorInDeps = true
				break
			}
			errorInDeps = false
		}

		if !errorInDeps {
			cmd = exec.Command(command[0], command[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Start()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Dev() error {
	platform := os.Getenv("PLATFORM")
	architecture := os.Getenv("ARCHITECTURE")

	command := []string{filepath.Join(binDir, "gomather-server-"+platform+"-"+architecture), "start"}
	return watch(command, []interface{}{Build, UnitTests, BinaryBuild})
}
