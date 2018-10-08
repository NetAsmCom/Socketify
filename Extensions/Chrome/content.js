var script = document.createElement("script");
script.src = chrome.extension.getURL("socketify.js");
document.documentElement.appendChild(script);

var isConnected = false;

chrome.runtime.onMessage.addListener(function (message, sender, sendResponse) {
    // TODO
});

window.addEventListener("message", function (event) {
    if (event.source !== window || event.data.type !== "socketify-outbound") {
        return;
    }

    switch (event.data.action) {
        case "connect": {
            if (isConnected) {
                return;
            }

            if (!confirm("Do you want to allow this page to open encrypted connection?")) {
                window.postMessage({
                    type: "socketify-inbound",
                    action: "error",
                    error: {
                        code: 1,
                        text: "user refused permission to open connection"
                    }
                }, "*");
                return;
            }

            chrome.runtime.sendMessage({
                action: "connect",
                encrypt: event.data.encrypt
            });
        } break;
        case "send": {
            if (!isConnected) {
                return;
            }

            chrome.runtime.sendMessage({
                action: "send",
                reliable: event.data.reliable,
                message: event.data.message
            });
        } break;
        case "disconnect": {
            if (!isConnected) {
                return;
            }

            chrome.runtime.sendMessage({
                action: "disconnect"
            });
        } break;
    }
}, false);
