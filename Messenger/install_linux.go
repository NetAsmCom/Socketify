// +build linux

package main

import "os"

func install(chromeExtID string, firefoxExtID string) bool {
	os.Stdout.Write([]byte("install: linux installation not implemented yet\n"))
	return false
}

func uninstall() {
	os.Stdout.Write([]byte("uninstall: linux uninstallation not implemented yet\n"))
}
