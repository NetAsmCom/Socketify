package main

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
