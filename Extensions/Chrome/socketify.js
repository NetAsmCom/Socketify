window.socketify = {
    connect: function (encrypt) {
    },
    send: function (reliable, message) {
    },
    disconnect: function () {
    },
    isConnected: false,
    isEncrypted: false,
    onConnect: function () { /* unhandled */ },
    onReceive: function (reliable, message) { /* unhandled */ },
    onDisconnect: function (error) { /* unhandled */ }
};

window.addEventListener("message", function (event) {
    if (event.source !== window || event.data.type !== "socketify-inbound") {
        return;
    }

    // TODO
}, false);
