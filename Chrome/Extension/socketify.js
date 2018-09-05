window.uuidv4 = function () {
    return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function (c) {
        var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
};

window.socketify = {
    _sockets: {},
    _post: function (message) {
        window.postMessage({
            _tab: {
                dir: "socketify-outbound",
                _ext: message
            }
        }, "*");
    },
    _handle: function (message) {
        var id = message.id;
        if (!id) {
            return;
        }

        var socket = socketify._sockets[id];
        if (!socket) {
            return;
        }

        var msg = message._msg;
        if (!msg) {
            return;
        }

        switch (msg.event) {
            case "open": {
                var openHandler = socket._handlers.onOpen;
                if (openHandler) {
                    openHandler(msg.address);
                }
            } return;
            case "connect": {
                var connectHandler = socket._handlers.onConnect;
                if (connectHandler && socket._id[0] === 's') {
                    connectHandler(msg.address);
                }
            } return;
            case "receive": {
                var receiveHandler = socket._handlers.onReceive;
                if (receiveHandler) {
                    if (socket._id[0] === 'c') {
                        receiveHandler(msg.payload);
                    } else {
                        receiveHandler(msg.address, msg.payload);
                    }
                }
            } return;
            case "disconnect": {
                var disconnectHandler = socket._handlers.onDisconnect;
                if (disconnectHandler && socket._id[0] === 's') {
                    disconnectHandler(msg.address, msg.error);
                }
            } return;
            case "close": {
                var closeHandler = socket._handlers.onClose;
                if (closeHandler) {
                    closeHandler(msg.error);
                }
            } return;
        }
    },
    tcpServer: function (address, handlers) {
        var id = `s-${uuidv4()}`;
        var socket = {
            _id: id,
            _handlers: handlers, // onOpen(addr), onConnect(addr), onReceive(addr, msg), onDisconnect(addr, err), onClose(err)
            send: function (target, message) {
                socketify._post({
                    id: id,
                    _msg: {
                        event: "send",
                        address: target,
                        payload: message
                    }
                });
            },
            drop: function (target) {
                socketify._post({
                    id: id,
                    _msg: {
                        event: "drop",
                        address: target
                    }
                });
            },
            close: function () {
                socketify._post({
                    id: id,
                    _msg: {
                        event: "close"
                    }
                });
            }
        };

        socketify._sockets[id] = socket;
        socketify._post({
            id: id,
            _msg: {
                event: "open-tcpServer",
                address: address
            }
        });

        return socket;
    },
    tcpClient: function (address, handlers) {
        var id = `c-${uuidv4()}`;
        var socket = {
            _id: id,
            _handlers: handlers, // onOpen(addr), onReceive(msg), onClose(err)
            send: function (message) {
                socketify._post({
                    id: id,
                    _msg: {
                        event: "send",
                        payload: message
                    }
                });
            },
            close: function () {
                socketify._post({
                    id: id,
                    _msg: {
                        event: "close"
                    }
                });
            }
        };

        socketify._sockets[id] = socket;
        socketify._post({
            id: id,
            _msg: {
                event: "open-tcpClient",
                address: address
            }
        });

        return socket;
    },
    udpPeer: function (address, handlers) {
        var id = `p-${uuidv4()}`;
        var socket = {
            _id: id,
            _handlers: handlers, // onOpen(addr), onReceive(addr, msg), onClose(err)
            send: function (target, message) {
                socketify._post({
                    id: id,
                    _msg: {
                        event: "send",
                        address: target,
                        payload: message
                    }
                });
            },
            close: function () {
                socketify._post({
                    id: id,
                    _msg: {
                        event: "close"
                    }
                });
            }
        };

        socketify._sockets[id] = socket;
        socketify._post({
            id: id,
            _msg: {
                event: "open-udpPeer",
                address: address
            }
        });

        return socket;
    }
};

window.addEventListener("message", function (event) {
    if (event.source !== window || event.data._tab.dir !== "socketify-inbound") {
        return;
    }

    socketify._handle(event.data._tab._ext);
}, false);
