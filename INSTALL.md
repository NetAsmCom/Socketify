# Installation Guides

You need to install both [`Extension`](#extension) and [`Messenger`](#messenger) to have [`Socketify` API](API.md) available.

## Contents

- [Extension](#extension)
  - [Chrome](#chrome)
  - [Firefox](#firefox)
  - [Safari](#safari)
- [Messenger](#messenger)

## Extension

### Chrome

![Load Chrome Extension](Installer/Chrome.gif)

1. Go to `chrome://extensions`
2. Turn `Developer mode` switch on
3. Click `Load unpacked` button
4. Navigate and select `Chrome` extension directory
5. Note extension's `ID` because you will need it while [installing `Messenger`](#messenger)

### Firefox

![Load Firefox Extension](Installer/Firefox.gif)

1. Go to `about:debugging`
2. Click `Load Temporary Add-on` button
3. Navigate and open `manifest.json` under `Firefox` directory
4. Note `Extension ID` because you will need it while [installing `Messenger`](#messenger)

### Safari

> TODO: Planned

## Messenger

![Build and Install Messenger Host App](Installer/Messenger.gif)

1. [Download and install Go](https://golang.org)
2. Open terminal/console
3. Go to `Messenger` directory
4. _(only for Windows)_ Get [`registry`](https://godoc.org/golang.org/x/sys/windows/registry) package
    ```console
    go get -u golang.org/x/sys/windows/registry
    ```
5. Build the app
    ```console
    go build
    ```
6. Install by specifying extension `ID`s
    ```console
    "Messenger.exe" -install -chromeExtID=<ID> -firefoxExtID=<ID>

    ./Messenger -install -chromeExtID=<ID> -firefoxExtID=<ID>
    ```
