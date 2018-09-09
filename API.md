# API Documentation

## Contents

- [uuidv4]()
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

## `function` uuidv4

Takes no parameter and returns _UUIDv4_ `string` back.

```js
var coolID = uuidv4();
console.log(`my cool universally unique ID is ${coolID}`);
```

## `object` socketify

Some description goes here

### `function` udpPeer(address, handlers)

- `string` **address**

  local address to bind socket

- `object` **handlers**

  contains `onOpen`, `onReceive`, `onClose` event functions

**returns** `object` [udpPeer]()

> This function does things!

```js
var myPeer = socketify.udpPeer(":9696", {
    onOpen: function (address) {
        console.log(`peer bound to <${address}>`);
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

### `function` tcpClient(address, handlers)

- `string` **address**

  server address to connect

- `object` **handlers**

  contains `onOpen`, `onReceive`, `onClose` event functions

**returns** `object` [tcpClient]()

> This function does things!

```js
var myClient = socketify.tcpClient("127.0.0.1:9696", {
    onOpen: function (address) {
        console.log(`client bound to <${address}> and connected`);
    },
    onReceive: function (message) {
        console.log(`client received: ${message}`);
    },
    onClose: function (error) {
        if (error) {
            console.log(`client closed with error: ${error}`);
        } else {
            console.log(`client closed`);
        }
    }
});
```

### `function` tcpServer(address, handlers)

- `string` **address**

  local address to bind socket

- `object` **handlers**

  contains `onOpen`, `onConnect`, `onReceive`, `onDisconnect` `onClose` event functions

**returns** `object` [tcpServer]()

> This function does things!

```js
var myServer = socketify.tcpServer(":9696", {
    onOpen: function (address) {
        console.log(`server bound to <${address}> and listening`);
    },
    onConnect: function (address) {
        console.log(`server connected to <${address}>`);
    },
    onReceive: function (address, message) {
        console.log(`server received <${address}>: ${message}`);
    },
    onDisconnect: function (address, error) {
        if (error) {
            console.log(`server disconnected from <${address}> with error: ${error}`);
        } else {
            console.log(`server disconnected from <${address}>`);
        }
    },
    onClose: function (error) {
        if (error) {
            console.log(`server closed with error: ${error}`);
        } else {
            console.log(`server closed`);
        }
    }
});
```

## udpPeer

## tcpClient

## tcpServer
