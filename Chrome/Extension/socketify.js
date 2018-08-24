window.uuidv4 = function () {
    return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function (c) {
        var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
};

window.socketify = {
    _sockets: {},
    _sendMessage: function (message) {
        message.type = "socketify-out";
        window.postMessage(message, "*");
    },
    _onMessage: function (message) {
        // TODO
    },
    tcpServer: function (endPoint, callback) { },
    tcpClient: function (endPoint, callback) { },
    udpPeer: function (endPoint, callback) {
        var id = uuidv4();
        var socket = {
            type: "udpPeer",
            id: id,
            onCreate: callback
        };

        this._sockets[id] = socket;
    }
};

window.addEventListener("message", function (event) {
    if (event.source === window || event.data.type === "socketify-in") {
        window.socketify._onMessage(event.data);
    }
}, false);
