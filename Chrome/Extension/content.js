var script = document.createElement("script");
script.src = chrome.extension.getURL("socketify.js");
document.documentElement.appendChild(script);

window.addEventListener("message", function (event) {
    if (event.source !== window || event.data._tab.dir !== "socketify-outbound") {
        return;
    }

    chrome.runtime.sendMessage(event.data._tab._ext);
}, false);

chrome.runtime.onMessage.addListener(function (message, sender, sendResponse) {
    window.postMessage({
        _tab: {
            dir: "socketify-inbound",
            _ext: message
        }
    }, "*");
});
