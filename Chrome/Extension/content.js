// Inject Socketify Script
var script = document.createElement("script");
script.src = chrome.extension.getURL("socketify.js");
document.documentElement.appendChild(script);

// Handle Page Messages
window.addEventListener("message", function (event) {
    if (event.source !== window || event.data._info.type !== "socketify-out") {
        return;
    }

    chrome.runtime.sendMessage(event.data);
}, false);

// Handle Extension Messages
chrome.runtime.onMessage.addListener(function (message, sender, sendResponse) {
    message._info.type = "socketify-in";
    window.postMessage(message, "*");
});
