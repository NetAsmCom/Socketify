var elem = document.createElement("script");
elem.src = chrome.extension.getURL("socketify.js");
elem.async = false;
document.documentElement.appendChild(elem);