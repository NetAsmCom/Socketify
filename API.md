# API Documentation

## Contents

- [uuidv4](#function-uuidv4)
- [socketify](#object-socketify)
  - [udpPeer](#function-udppeeraddress-handlers)
  - [tcpClient](#function-tcpclientaddress-handlers)
  - [tcpServer](#function-tcpserveraddress-handlers)
- [udpPeer](#object-udppeer)
  - [id](#string-id)
  - [onOpen](#function-onopenaddress)
  - [onReceive](#function-onreceiveaddress-message)
  - [onClose](#function-oncloseerror)
  - [send](#function-sendaddress-message)
  - [close](#function-close)
- [tcpClient](#object-tcpclient)
  - [id](#string-id-1)
  - [onOpen](#function-onopenaddress-1)
  - [onReceive](#function-onreceivemessage)
  - [onClose](#function-oncloseerror-1)
  - [send](#function-sendmessage)
  - [close](#function-close-1)
- [tcpServer](#object-tcpserver)
  - [id](#string-id-2)
  - [onOpen](#function-onopenaddress-2)
  - [onConnect](#function-onconnectaddress)
  - [onReceive](#function-onreceiveaddress-message-1)
  - [onDisconnect](#function-ondisconnectaddress-error)
  - [onClose](#function-oncloseerror-2)
  - [send](#function-sendaddress-message-1)
  - [drop](#function-dropaddress)
  - [close](#function-close-2)

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

**returns** [`object` udpPeer](#object-udppeer)

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

**returns** [`object` tcpClient](#object-tcpclient)

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

**returns** [`object` tcpServer](#object-tcpserver)

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

## `object` udpPeer

### `string` id

### `function` onOpen(address)

### `function` onReceive(address, message)

### `function` onClose(error)

### `function` send(address, message)

### `function` close()

## `object` tcpClient

### `string` id

### `function` onOpen(address)

### `function` onReceive(message)

### `function` onClose(error)

### `function` send(message)

### `function` close()

## `object` tcpServer

### `string` id

### `function` onOpen(address)

### `function` onConnect(address)

### `function` onReceive(address, message)

### `function` onDisconnect(address, error)

### `function` onClose(error)

### `function` send(address, message)

### `function` drop(address)

### `function` close()
