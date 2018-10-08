window.socketify = {
    connect: function (encrypt) {
        window.postMessage({
            type: "socketify-outbound",
            action: "connect",
            encrypt: (typeof encrypt === typeof true) && encrypt
        }, "*");
    },
    send: function (reliable, message) {
        window.postMessage({
            type: "socketify-outbound",
            action: "send",
            reliable: (typeof reliable === typeof true) && reliable,
            message: message || {}
        }, "*");
    },
    disconnect: function () {
        window.postMessage({
            type: "socketify-outbound",
            action: "disconnect"
        }, "*");
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

    switch (event.data.action) {
        case "setProps":
            window.socketify.isConnected = event.data.isConnected;
            window.socketify.isEncrypted = event.data.isEncrypted;
            break;
        case "onConnect":
            window.socketify.onConnect();
            break;
        case "onReceive":
            window.socketify.onReceive(event.data.reliable, event.data.message);
            break;
        case "onDisconnect":
            window.socketify.onDisconnect(event.data.error);
            break;
    }
}, false);
