// Inject Socketify Script
var script = document.createElement("script");
script.src = chrome.extension.getURL("socketify.js");
document.documentElement.appendChild(script);

// Handle Page Messages
window.addEventListener("message", function (event) {
    if (event.source !== window || event.data._type !== "socketify-out") {
        return;
    }

    // TODO
}, false);

// Handle Extension Messages
chrome.runtime.onMessage.addListener(function (message, sender, sendResponse) {
    // TODO
});
