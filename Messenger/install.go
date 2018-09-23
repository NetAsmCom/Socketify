package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	chromeManifest = `{
    "name": "net.socketify.messenger",
    "description": "Socketify - TCP and UDP Sockets API",
    "path": "BIN_PATH",
    "type": "stdio",
    "allowed_origins": [
        "chrome-extension://EXT_ID/"
	]
}`

	firefoxManifest = `{
	"name": "net.socketify.messenger",
	"description": "Socketify - TCP and UDP Sockets API",
	"path": "BIN_PATH",
	"type": "stdio",
	"allowed_extensions": [
		"EXT_ID"
	]
}`
)

func install(chromeExtID string, firefoxExtID string) bool {
	currentUser, error := user.Current()
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("intall: cannot get current user\n"))
		return false
	}
	_ = os.Mkdir(filepath.Join(currentUser.HomeDir, "Socketify"), os.ModePerm)

	binaryPath, error := os.Executable()
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("install: cannot get binary path\n"))
		return false
	}

	binaryStat, error := os.Stat(binaryPath)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("install: cannot get binary stat\n"))
		return false
	}

	if binaryStat.Mode().IsRegular() != true {
		os.Stdout.Write([]byte("install: binary is not a regular file\n"))
		return false
	}

	binaryBytes, error := ioutil.ReadFile(binaryPath)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("install: cannot read binary\n"))
		return false
	}

	userBinaryPath := filepath.Join(currentUser.HomeDir, "Socketify", binaryStat.Name())
	error = ioutil.WriteFile(userBinaryPath, binaryBytes, os.ModePerm)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("install: cannot write binary\n"))
		return false
	}

	chromeManifestString := chromeManifest
	chromeManifestString = strings.Replace(chromeManifestString, "BIN_PATH", userBinaryPath, 1)
	chromeManifestString = strings.Replace(chromeManifestString, "EXT_ID", chromeExtID, 1)

	firefoxManifestString := firefoxManifest
	firefoxManifestString = strings.Replace(firefoxManifestString, "BIN_PATH", userBinaryPath, 1)
	firefoxManifestString = strings.Replace(firefoxManifestString, "EXT_ID", firefoxExtID, 1)

	switch runtime.GOOS {
	case "darwin":
		chromeManifestPath := filepath.Join(currentUser.HomeDir,
			"Library", "Application Support",
			"Google", "Chrome", "NativeMessagingHosts",
			"net.socketify.messenger.json")
		error = ioutil.WriteFile(chromeManifestPath, []byte(chromeManifestString), os.ModePerm)
		if error != nil {
			os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
			os.Stdout.Write([]byte("install: cannot write chrome manifest\n"))
			return false
		}

		firefoxManifestPath := filepath.Join(currentUser.HomeDir,
			"Library", "Application Support",
			"Mozilla", "NativeMessagingHosts",
			"net.socketify.messenger.json")
		error = ioutil.WriteFile(firefoxManifestPath, []byte(firefoxManifestString), os.ModePerm)
		if error != nil {
			os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
			os.Stdout.Write([]byte("install: cannot write firefox manifest\n"))
			return false
		}
	case "windows":
	case "linux":
	default:
		os.Stdout.Write([]byte("install: unknown os platform\n"))
		return false
	}

	return true
}

func uninstall() bool {
	return true
}
