// Expose UUIDv4 API
window.uuidv4 = function () {
    return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function (c) {
        var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
};

// Expose Socketify API
window.socketify = {
    _sockets: {},
    _sendMessage: function (info, payload) {
        info.type = "socketify-out";
        window.postMessage({
            _info: info,
            _payload: payload
        }, "*");
    },
    _onMessage: function (info, payload) {
        var id = info.id;
        var socket = socketify._sockets[id];
        switch (info.command) {
            // TODO: Handle other events?
            case "open": {
                if (info.result.success) {
                    socket.endPoint = info.endPoint;
                    socket._onOpen(socket, undefined);
                } else {
                    socket._onOpen(undefined, info.result.error);
                    delete socketify._sockets[id];
                }
            } break;
            case "close": {
                socket.onClose(info.result.error);
                delete socketify._sockets[id];
            } break;
        }
    },
    tcpServer: function (endPoint, callback) {
        // TODO
    },
    tcpClient: function (endPoint, callback) {
        // TODO
    },
    udpPeer: function (endPoint, callback) {
        var id = uuidv4();
        socketify._sockets[id] = {
            id: id,
            endPoint: endPoint,
            _onOpen: callback,
            onMessage: function (sender, message) { /* Unhandled - User should override! */ },
            onClose: function (error) { /* Unhandled - User should override! */ },
            sendMessage: function (target, message) {
                socketify._sendMessage({
                    command: "send",
                    id: id,
                    endPoint: target
                }, message);
            },
            close: function () {
                socketify._sendMessage({
                    command: "close",
                    id: id
                }, undefined);
            }
        };
        socketify._sendMessage({
            command: "udpPeer-open",
            id: id,
            endPoint: endPoint
        }, undefined);
    }
};

// Handle Content Messages
window.addEventListener("message", function (event) {
    if (event.source === window && event.data._info.type === "socketify-in") {
        window.socketify._onMessage(event.data._info, event.data._payload);
    }
}, false);
