// +build darwin

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func install(chromeExtID string, firefoxExtID string) bool {
	currentUser, error := user.Current()
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("install: cannot get current user\n"))
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
	chromeManifestString = strings.Replace(chromeManifestString, "BIN_PATH", filepath.ToSlash(userBinaryPath), 1)
	chromeManifestString = strings.Replace(chromeManifestString, "EXT_ID", chromeExtID, 1)

	firefoxManifestString := firefoxManifest
	firefoxManifestString = strings.Replace(firefoxManifestString, "BIN_PATH", filepath.ToSlash(userBinaryPath), 1)
	firefoxManifestString = strings.Replace(firefoxManifestString, "EXT_ID", firefoxExtID, 1)

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

	return true
}

func uninstall() {
	currentUser, error := user.Current()
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("uninstall: %s\n", error.Error())))
		os.Stdout.Write([]byte("uninstall: cannot get current user\n"))
		return
	}

	binaryPath, error := os.Executable()
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("uninstall: cannot get binary path\n"))
		return
	}

	binaryStat, error := os.Stat(binaryPath)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("uninstall: cannot get binary stat\n"))
		return
	}

	if binaryStat.Mode().IsRegular() != true {
		os.Stdout.Write([]byte("uninstall: binary is not a regular file\n"))
		return
	}

	userBinaryPath := filepath.Join(currentUser.HomeDir, "Socketify", binaryStat.Name())
	error = os.Remove(userBinaryPath)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("uninstall: %s\n", error.Error())))
		os.Stdout.Write([]byte("uninstall: cannot remove binary\n"))
	}

	chromeManifestPath := filepath.Join(currentUser.HomeDir,
		"Library", "Application Support",
		"Google", "Chrome", "NativeMessagingHosts",
		"net.socketify.messenger.json")
	error = os.Remove(chromeManifestPath)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("uninstall: %s\n", error.Error())))
		os.Stdout.Write([]byte("uninstall: cannot remove chrome manifest\n"))
	}

	firefoxManifestPath := filepath.Join(currentUser.HomeDir,
		"Library", "Application Support",
		"Mozilla", "NativeMessagingHosts",
		"net.socketify.messenger.json")
	error = os.Remove(firefoxManifestPath)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("uninstall: %s\n", error.Error())))
		os.Stdout.Write([]byte("uninstall: cannot remove firefox manifest\n"))
	}
}
