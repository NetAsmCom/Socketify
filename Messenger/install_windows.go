// +build windows

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
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

	chromeManifestPath := filepath.Join(currentUser.HomeDir, "Socketify", "net.socketify.messenger_chrome.json")
	error = ioutil.WriteFile(chromeManifestPath, []byte(chromeManifestString), os.ModePerm)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("install: cannot write chrome manifest\n"))
		return false
	}

	firefoxManifestPath := filepath.Join(currentUser.HomeDir, "Socketify", "net.socketify.messenger_firefox.json")
	error = ioutil.WriteFile(firefoxManifestPath, []byte(firefoxManifestString), os.ModePerm)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte([]byte("install: cannot write firefox manifest\n")))
		return false
	}

	key, _, error := registry.CreateKey(registry.CURRENT_USER, `SOFTWARE\Google\Chrome\NativeMessagingHosts\net.socketify.messenger`, registry.WRITE)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("install: cannot create chrome registry key\n"))
		return false
	}
	error = key.SetStringValue("", chromeManifestPath)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("install: cannot set chrome registry key\n"))
		return false
	}

	key, _, error = registry.CreateKey(registry.CURRENT_USER, `SOFTWARE\Mozilla\NativeMessagingHosts\net.socketify.messenger`, registry.WRITE)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("install: cannot create firefox registry key\n"))
		return false
	}
	error = key.SetStringValue("", firefoxManifestPath)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("install: %s\n", error.Error())))
		os.Stdout.Write([]byte("install: cannot set firefox registry key\n"))
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

	chromeManifestPath := filepath.Join(currentUser.HomeDir, "Socketify", "net.socketify.messenger_chrome.json")
	error = os.Remove(chromeManifestPath)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("uninstall: %s\n", error.Error())))
		os.Stdout.Write([]byte("uninstall: cannot remove chrome manifest\n"))
	}

	firefoxManifestPath := filepath.Join(currentUser.HomeDir, "Socketify", "net.socketify.messenger_firefox.json")
	error = os.Remove(firefoxManifestPath)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("uninstall: %s\n", error.Error())))
		os.Stdout.Write([]byte("uninstall: cannot remove firefox manifest\n"))
	}

	error = registry.DeleteKey(registry.CURRENT_USER, `SOFTWARE\Google\Chrome\NativeMessagingHosts\net.socketify.messenger`)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("uninstall: %s\n", error.Error())))
		os.Stdout.Write([]byte("uninstall: cannot delete chrome registry key\n"))
	}

	error = registry.DeleteKey(registry.CURRENT_USER, `SOFTWARE\Mozilla\NativeMessagingHosts\net.socketify.messenger`)
	if error != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("uninstall: %s\n", error.Error())))
		os.Stdout.Write([]byte("uninstall: cannot delete firefox registry key\n"))
	}
}
