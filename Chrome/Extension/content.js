// Inject Socketify API
var script = document.createElement("script");
script.src = chrome.extension.getURL("socketify.js");
document.documentElement.appendChild(script);
