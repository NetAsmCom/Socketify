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

        // TODO: handle tcpServer, tcpClient
        switch (msg.event) {
            case "open": {
                var openHandler = socket._handlers.onOpen;
                if (openHandler) {
                    openHandler(msg.address, msg.error);
                }
            } return;
            case "receive": {
                var receiveHandler = socket._handlers.onReceive;
                if (receiveHandler) {
                    receiveHandler(msg.address, msg.payload);
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
            _handlers: handlers, // onOpen, onConnect, onReceive, onDisconnect, onClose
            close: function () { }
        };

        socketify._sockets[id] = socket;

        // TODO: _post(open-tcpServer)

        return socket;
    },
    tcpClient: function (address, handlers) {
        var id = `c-${uuidv4()}`;
        var socket = {
            _id: id,
            _handlers: handlers, // onOpen, onReceive, onClose
            send: function (message) { },
            close: function () { }
        };

        socketify._sockets[id] = socket;

        // TODO: _post(open-tcpClient)

        return socket;
    },
    udpPeer: function (address, handlers) {
        var id = `p-${uuidv4()}`;
        var socket = {
            _id: id,
            _handlers: handlers, // onOpen, onReceive, onClose
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
