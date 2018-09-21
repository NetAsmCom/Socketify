# Socketify

> TCP and UDP Sockets API on Chrome, Firefox and Safari desktop browsers with extensions via native messaging.

**What?** A cross-platform, cross-browser extension for desktop browsers that injects simple & easy-to-use `UdpPeer`, `TcpServer` and `TcpClient` sockets API into page window, available in plain JavaScript.

**Why?** I was prototyping a web-based multiplayer-online game then realized that WebSocket and WebRTC standard APIs are not flexible enough to achieve custom networking solutions when needed. After that I took the challenge and decided to provide raw UDP and TCP sockets with a simple API so that people can implement their own network transport layer on top. Especially for real-time games, you'd better use thin UDP transport layer to fight with network congestion!

**How?** Explain how this project achieve the solution.

## Getting Started

- [Installation Guides](INSTALL.md)
- [Example Test Page](Example/index.html)
- [API Documentation](API.md)

## TODO

- [x] Native Messaging Host
- [x] Socketify API
- [x] Chrome Extension
- [ ] Firefox Porting
- [ ] Installation Guides
- [ ] API Documentation
- [ ] Safari Extension
- [ ] Extension Popup Menu
- [ ] Host App Installer
- [ ] Encryption Support
- [ ] Unity Support
