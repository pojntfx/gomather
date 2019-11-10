//+build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/mholt/archiver"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var gocmd = mg.GoCmd()
var tempdir = os.TempDir()
var protocOut = filepath.Join(tempdir, "usr", "local", "protoc")

func ProtocInstallDependencies() error {
	platform := os.Getenv("PLATFORM")
	architecture := os.Getenv("ARCHITECTURE")

	propPlatform := "linux"
	if platform == "darwin" {
		propPlatform = "osx"
	}
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
