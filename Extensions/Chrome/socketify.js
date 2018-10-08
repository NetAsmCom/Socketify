window.socketify = {
};

window.addEventListener("message", function (event) {
    if (event.source !== window || event.data.type !== "socketify-inbound") {
        return;
    }

    // TODO
}, false);
