# API Documentation

## Contents

- [uuidv4](#function-uuidv4)
- [socketify](#object-socketify)
  - [udpPeer](#function-udppeerbindaddress-handlers)
  - [tcpClient](#function-tcpclientserveraddress-handlers)
  - [tcpServer](#function-tcpserverlistenaddress-handlers)
- [udpPeer](#object-udppeer)
  - [id](#string-id)
  - [onOpen](#function-onopenbindaddress)
  - [onReceive](#function-onreceivepeeraddress-message)
  - [onClose](#function-oncloseerror)
  - [send](#function-sendpeeraddress-message)
  - [close](#function-close)
- [tcpClient](#object-tcpclient)
  - [id](#string-id-1)
  - [onOpen](#function-onopenlocaladdress)
  - [onReceive](#function-onreceivemessage)
  - [onClose](#function-oncloseerror-1)
  - [send](#function-sendmessage)
  - [close](#function-close-1)
- [tcpServer](#object-tcpserver)
  - [id](#string-id-2)
  - [onOpen](#function-onopenlistenaddress)
  - [onConnect](#function-onconnectclientaddress)
  - [onReceive](#function-onreceiveclientaddress-message)
  - [onDisconnect](#function-ondisconnectclientaddress-error)
  - [onClose](#function-oncloseerror-2)
  - [send](#function-sendclientaddress-message)
  - [drop](#function-dropclientaddress)
  - [close](#function-close-2)

## `function` uuidv4

Takes no parameter, generates and returns [UUIDv4&nearr;](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) `string` back.

```js
var coolID = uuidv4();
console.log(`my cool universally unique ID is ${coolID}`);
```

## `object` socketify

This object being injected by content script at document start event and exposes socket creation functions to window.

### `function` udpPeer(bindAddress, handlers)

- `string` **bindAddress**

  local address to bind socket

- `object` **handlers**

  contains `onOpen`, `onReceive`, `onClose` event handling functions

  **returns** [`object` udpPeer](#object-udppeer)

Creates UDP socket, binds to specified address and fires `onOpen` or `onClose` event afterwards.

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

### `function` tcpClient(serverAddress, handlers)

- `string` **serverAddress**

  server address to connect

- `object` **handlers**

  contains `onOpen`, `onReceive`, `onClose` event functions

  **returns** [`object` tcpClient](#object-tcpclient)

Creates TCP socket, opens connection to server with specified address and fires `onOpen` or `onClose` event afterwards.

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

### `function` tcpServer(listenAddress, handlers)

- `string` **listenAddress**

  local address to listen from

- `object` **handlers**

  contains `onOpen`, `onConnect`, `onReceive`, `onDisconnect` `onClose` event functions

  **returns** [`object` tcpServer](#object-tcpserver)

Creates TCP socket, starts listening on specified address and fires `onOpen` or `onClose` afterwards.

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

## `object` udpPeer

UDP socket instance on browser that bridges native calls.

### `string` id

Unique socket id assigned by window to specify instance, primarly used to dispatch calls.

```js
// Sample code
```

### `function` onOpen(bindAddress)

- `string` **bindAddress**

  local bound address

> This function does things!

```js
// Sample code
```

### `function` onReceive(peerAddress, message)

- `string` **peerAddress**

  sender peer address

- `object` **message**

  message received from peer

> This function does things!

```js
// Sample code
```

### `function` onClose(error)

- `string` **error** _(optional)_

  error message, will be `undefined` if socket closed with [`function` close()](#function-close)

> This function does things!

```js
// Sample code
```

### `function` send(peerAddress, message)

- `string` **peerAddress**

  target peer address

- `object` **message**

  message to send to peer

> This function does things!

```js
// Sample code
```

### `function` close()

> This function does things!

```js
// Sample code
```

## `object` tcpClient

TCP client socket instance on browser that bridges native calls.

### `string` id

Unique socket id assigned by window to specify instance, primarly used to dispatch calls.

```js
// Sample code
```

### `function` onOpen(localAddress)

- `string` **localAddress**

  local bound address

> This function does things!

```js
// Sample code
```

### `function` onReceive(message)

- `object` **message**

  message received from server

> This function does things!

```js
// Sample code
```

### `function` onClose(error)

- `string` **error** _(optional)_

  error message, will be `undefined` if socket closed with [`function` close()](#function-close-1)

> This function does things!

```js
// Sample code
```

### `function` send(message)

- `object` **message**

  message to send to server

> This function does things!

```js
// Sample code
```

### `function` close()

> This function does things!

```js
// Sample code
```

## `object` tcpServer

TCP server socket instance on browser that bridges native calls.

### `string` id

Unique socket id assigned by window to specify instance, primarly used to dispatch calls.

```js
// Sample code
```

### `function` onOpen(listenAddress)

- `string` **listenAddress**

  local bound address

> This function does things!

```js
// Sample code
```

### `function` onConnect(clientAddress)

- `string` **clientAddress**

  connected client address

> This function does things!

```js
// Sample code
```

### `function` onReceive(clientAddress, message)

- `string` **clientAddress**

  sender client address

- `object` **message**

  message received from client

> This function does things!

```js
// Sample code
```

### `function` onDisconnect(clientAddress, error)

- `string` **clientAddress**

  disconnected client address

- `string` **error** _(optional)_

  error message, will be `undefined` if connection closed with [`function` drop()](#function-dropclientaddress)

> This function does things!

```js
// Sample code
```

### `function` onClose(error)

- `string` **error** _(optional)_

  error message, will be `undefined` if socket closed with [`function` close()](#function-close-2)

> This function does things!

```js
// Sample code
```

### `function` send(clientAddress, message)

- `string` **clientAddress**

  target client address

- `object` **message**

  message to send to client

> This function does things!

```js
// Sample code
```

### `function` drop(clientAddress)

- `string` **clientAddress**

  target client address

> This function does things!

```js
// Sample code
```

### `function` close()

> This function does things!

```js
// Sample code
```
