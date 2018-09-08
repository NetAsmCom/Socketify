# API Documentation

## Contents

- [socketify]()
  - [udpPeer]()
  - [tcpClient]()
  - [tcpServer]()
- [udpPeer]()
  - [id]()
  - [onOpen]()
  - [onReceive]()
  - [onClose]()
  - [send]()
  - [close]()
- [tcpClient]()
  - [id]()
  - [onOpen]()
  - [onReceive]()
  - [onClose]()
  - [send]()
  - [close]()
- [tcpServer]()
  - [id]()
  - [onOpen]()
  - [onConnect]()
  - [onReceive]()
  - [onDisconnect]()
  - [onClose]()
  - [send]()
  - [drop]()
  - [close]()

## socketify

### `function` udpPeer(address, handlers)

- `string` address

  local address to bind socket

- `object` handlers

  contains `onOpen`, `onReceive`, `onClose` event functions

**returns** `object` [udpPeer]()

> This function does things!

```js
var myPeer = socketify.udpPeer(":9696", {
    onOpen: function (address) {
        console.log(`peer opened and bound to <${address}>`);
    },
    onReceive: function (address, message) {
        console.log(`peer received <${address}>: ${message}`);
    },
    onClose: function (error) {
        if (error) {
            console.log(`peer closed with error: ${error}`);
        } else {
            console.log(`peer closed`);
        }
    }
});
```

### tcpClient(address, handlers)

### tcpServer(address, handlers)

## udpPeer

## tcpClient

## tcpServer
