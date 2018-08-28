window.uuidv4 = function () {
    return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function (c) {
        var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
};

window.socketify = {
    _sockets: {},
    _sendMessage: function (message) {
        message._type = "socketify-out";
        window.postMessage(message, "*");
    },
    _onMessage: function (message) {
        switch (message._info.command) {
            // TODO
        }
    },
    tcpServer: function (endPoint, callback) { },
    tcpClient: function (endPoint, callback) { },
    udpPeer: function (endPoint, callback) {
        var id = uuidv4();
        socketify._sockets[id] = {
            _id: id,
            _onOpen: callback,
            onMessage: function (sender, message) { },
            onClose: function () { },
            sendMessage: function (target, message) {
                message._info = {
                    command: "udpPeer-send",
                    id: id,
                    target: target
                };
                socketify._sendMessage(message);
            },
            close: function () {
                socketify._sendMessage({
                    _info: {
                        command: "udpPeer-close",
                        id: id
                    }
                });
            }
        };
        socketify._sendMessage({
            _info: {
                command: "udpPeer-open",
                id: id,
                endPoint: endPoint
            }
        });
    }
};

window.addEventListener("message", function (event) {
    if (event.source === window || event.data._type === "socketify-in") {
        window.socketify._onMessage(event.data);
    }
}, false);
