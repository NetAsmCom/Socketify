var script = document.createElement("script");
script.src = chrome.extension.getURL("socketify.js");
document.documentElement.appendChild(script);

chrome.runtime.onMessage.addListener(function (message, sender, sendResponse) {
    // TODO
});

window.addEventListener("message", function (event) {
    if (event.source !== window || event.data.type !== "socketify-outbound") {
        return;
    }

    // TODO
}, false);
