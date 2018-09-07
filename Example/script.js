var udpPeerElem = undefined;
var tcpClientElem = undefined;
var tcpServerElem = undefined;

var inputElem = undefined;

var logsElem = undefined;

function init() {
    console.log(`init()`);

    setTimeout(function () {
        if (window.socketify) {
            document.getElementById("page").style.display = "block";
        } else {
            document.getElementById("unavailable").style.display = "block";
        }
    }, 100);

    udpPeerElem = document.getElementById("udpPeer");
    tcpClientElem = document.getElementById("tcpClient");
    tcpServerElem = document.getElementById("tcpServer");

    inputElem = document.getElementById("input");

    logsElem = document.getElementById("logs");
}

var udpPeer = undefined;
var tcpClient = undefined;
var tcpServer = undefined;

var activeType = "udpPeer";

function switchTo(type) {
    console.log(`switchTo(${type}) - activeType: ${activeType}`);

    if (type === activeType) {
        return;
    }

    closeClick();

    switch (type) {
        case "udpPeer": {
            activeType = type;

            udpPeerElem.className = "active";
            tcpClientElem.className = "";
            tcpServerElem.className = "";
        } break;
        case "tcpClient": {
            activeType = type;

            udpPeerElem.className = "";
            tcpClientElem.className = "active";
            tcpServerElem.className = "";
        } break;
        case "tcpServer": {
            activeType = type;

            udpPeerElem.className = "";
            tcpClientElem.className = "";
            tcpServerElem.className = "active";
        } break;
    }

    console.log(`~switchTo(${type}) - activeType: ${activeType}`);
}

function getInput() {
    var input = inputElem.value;
    inputElem.value = "";
    return input;
}

var logIDs = [];
var lastLogID = 0;

function addLog(log) {
    if (!log) {
        return;
    }

    var isScrolledToBottom = logsElem.scrollHeight - logsElem.clientHeight <= logsElem.scrollTop + 1;

    var logID = lastLogID++;
    if (lastLogID > 1024) {
        lastLogID = 0;
    }

    var logElem = document.createElement("div");
    logElem.id = `log-${logID}`;
    {
        if (log.type) {
            var typeElem = document.createElement("span");
            typeElem.className = "type";
            typeElem.innerHTML = log.type;
            logElem.appendChild(typeElem);
        }

        if (log.outEvent) {
            var outEventElem = document.createElement("span");
            outEventElem.className = "outEvent";
            outEventElem.innerHTML = log.outEvent + " &rarr;";
            logElem.appendChild(outEventElem);
        }

        if (log.inEvent) {
            var inEventElem = document.createElement("span");
            inEventElem.className = "inEvent";
            inEventElem.innerHTML = log.inEvent + " &larr;";
            logElem.appendChild(inEventElem);
        }

        if (log.address) {
            var addressElem = document.createElement("span");
            addressElem.innerHTML = `address{<b>${log.address}</b>}`;
            logElem.appendChild(addressElem);
        }

        var payloadElem = undefined;
        if (log.message || log.error) {
            payloadElem = document.createElement("div");
            payloadElem.className = "payload";
        }

        if (log.message) {
            var messageElem = document.createElement("span");
            messageElem.innerHTML = `message{<b>${log.message.length} chars</b>}`;
            logElem.appendChild(messageElem);

            if (payloadElem) {
                var messagePayloadElem = document.createElement("div");
                messagePayloadElem.innerText = log.message;
                payloadElem.appendChild(messagePayloadElem);
            }
        }

        if (log.error) {
            var errorElem = document.createElement("span");
            errorElem.className = "error";
            errorElem.innerHTML = `error{<b>${log.error.length} chars</b>}`;
            logElem.appendChild(errorElem);

            if (payloadElem) {
                var errorPayloadElem = document.createElement("div");
                errorPayloadElem.className = "error";
                errorPayloadElem.innerText = log.error;
                payloadElem.appendChild(errorPayloadElem);
                console.log("error payload");
            }
        }

        if (log.usage) {
            var usageElem = document.createElement("span");
            usageElem.className = "error";
            usageElem.innerHTML = log.usage;
            logElem.appendChild(usageElem);
        }

        if (payloadElem) {
            logElem.appendChild(payloadElem);
        }
    }
    logsElem.appendChild(logElem);

    logIDs.push(logID);
    if (logIDs.length > 32) {
        var topLogID = logIDs.shift();
        if (topLogID !== undefined) {
            logsElem.removeChild(document.getElementById(`log-${topLogID}`));
        }
    }

    if (isScrolledToBottom) {
        logsElem.scrollTop = logsElem.scrollHeight - logsElem.clientHeight;
    }
}

function openClick(input) {
    switch (activeType) {
        case "udpPeer": {
            if (!udpPeer) {
                addLog({
                    type: "udpPeer",
                    outEvent: "open",
                    address: input
                });
                console.log(`[udpPeer] open: ${input}`);
                udpPeer = socketify.udpPeer(input, {
                    onOpen: function (address) {
                        addLog({
                            type: "udpPeer",
                            inEvent: "onOpen",
                            address: address
                        });
                        console.log(`[udpPeer] onOpen: ${address}`);
                    },
                    onReceive: function (address, message) {
                        addLog({
                            type: "udpPeer",
                            inEvent: "onReceive",
                            address: address,
                            message: message
                        });
                        console.log(`[udpPeer] onReceive <${address}>: ${message}`);
                    },
                    onClose: function (error) {
                        udpPeer = undefined;
                        addLog({
                            type: "udpPeer",
                            inEvent: "onClose",
                            error: error
                        });
                        console.log(`[udpPeer] onClose: ${error}`);
                    }
                });
            }
        } break;
        case "tcpClient": {
            if (!tcpClient) {
                addLog({
                    type: "tcpClient",
                    outEvent: "open",
                    address: input
                });
                console.log(`[tcpClient] open: ${input}`);
                tcpClient = socketify.tcpClient(input, {
                    onOpen: function (address) {
                        addLog({
                            type: "tcpClient",
                            inEvent: "onOpen",
                            address: address
                        });
                        console.log(`[tcpClient] onOpen: ${address}`);
                    },
                    onReceive: function (message) {
                        addLog({
                            type: "tcpClient",
                            inEvent: "onReceive",
                            message: message
                        });
                        console.log(`[tcpClient] onReceive: ${message}`);
                    },
                    onClose: function (error) {
                        tcpClient = undefined;
                        addLog({
                            type: "tcpClient",
                            inEvent: "onClose",
                            error: error
                        });
                        console.log(`[tcpClient] onClose: ${error}`);
                    }
                });
            }
        } break;
        case "tcpServer": {
            if (!tcpServer) {
                addLog({
                    type: "tcpServer",
                    outEvent: "open",
                    address: input
                });
                console.log(`[tcpServer] open: ${input}`);
                tcpServer = socketify.tcpServer(input, {
                    onOpen: function (address) {
                        addLog({
                            type: "tcpServer",
                            inEvent: "onOpen",
                            address: address
                        });
                        console.log(`[tcpServer] onOpen: ${address}`);
                    },
                    onConnect: function (address) {
                        addLog({
                            type: "tcpServer",
                            inEvent: "onConnect",
                            address: address
                        });
                        console.log(`[tcpServer] onConnect <${address}>`);
                    },
                    onReceive: function (address, message) {
                        addLog({
                            type: "tcpServer",
                            inEvent: "onReceive",
                            address: address,
                            message: message
                        });
                        console.log(`[tcpServer] onReceive <${address}>: ${message}`);
                    },
                    onDisconnect: function (address, error) {
                        addLog({
                            type: "tcpServer",
                            inEvent: "onDisconnect",
                            address: address,
                            error: error
                        });
                        console.log(`[tcpServer] onDisconnect <${address}>: ${error}`);
                    },
                    onClose: function (error) {
                        tcpServer = undefined;
                        addLog({
                            type: "tcpServer",
                            inEvent: "onClose",
                            error: error
                        });
                        console.log(`[tcpServer] onClose: ${error}`);
                    }
                });
            }
        } break;
    }
}

function sendClick(input) {
    console.log(`sendClick(${input})`);

    switch (activeType) {
        case "udpPeer": {
            if (udpPeer) {
                var blocks = input.split(' ');
                if (blocks.length > 1) {
                    var message = input.substring(blocks[0].length + 1);
                    if (message.length > 0) {
                        addLog({
                            type: "udpPeer",
                            outEvent: "send",
                            address: blocks[0],
                            message: message
                        });
                        console.log(`[udpPeer] send <${blocks[0]}>: ${message}`);
                        udpPeer.send(blocks[0], message);
                    }
                }
            }
        } break;
        case "tcpClient": {
            if (tcpClient) {
                if (input.length > 0) {
                    addLog({
                        type: "tcpClient",
                        outEvent: "send",
                        message: input
                    });
                    console.log(`[tcpClient] send: ${message}`);
                    tcpClient.send(input);
                }
            }
        } break;
        case "tcpServer": {
            if (tcpServer) {
                var blocks = input.split(' ');
                if (blocks.length > 1) {
                    var message = input.substring(blocks[0].length + 1);
                    if (message.length > 0) {
                        addLog({
                            type: "tcpServer",
                            outEvent: "send",
                            address: blocks[0],
                            message: message
                        });
                        console.log(`[tcpServer] send <${blocks[0]}>: ${message}`);
                        tcpServer.send(blocks[0], message);
                    }
                }
            }
        } break;
    }
}

function dropClick(input) {
    console.log(`dropClick(${input})`);

    switch (activeType) {
        case "udpPeer": {

        } break;
        case "tcpClient": {

        } break;
        case "tcpServer": {
            if (tcpServer) {
                if (input.length > 0) {
                    addLog({
                        type: "tcpServer",
                        outEvent: "drop",
                        address: input
                    });
                    console.log(`[tcpServer] drop: ${input}`);
                    tcpServer.drop(input);
                }
            }
        } break;
    }
}

function closeClick() {
    console.log(`closeClick()`);

    switch (activeType) {
        case "udpPeer": {
            if (udpPeer) {
                addLog({
                    type: "udpPeer",
                    outEvent: "close"
                });
                console.log(`[udpPeer] close`);

                udpPeer.close();
                udpPeer = undefined;
            }
        } break;
        case "tcpClient": {
            if (tcpClient) {
                addLog({
                    type: "tcpClient",
                    outEvent: "close"
                });
                console.log(`[tcpClient] close`);

                tcpClient.close();
                tcpClient = undefined;
            }
        } break;
        case "tcpServer": {
            if (tcpServer) {
                addLog({
                    type: "tcpServer",
                    outEvent: "close"
                });
                console.log(`[tcpServer] close`);

                tcpServer.close();
                tcpServer = undefined;
            }
        } break;
    }
}