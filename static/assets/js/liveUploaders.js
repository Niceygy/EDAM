let isLiveUploaderCountActive = false;
let seenUploaders = [];

let protocol = "ws";
if (window.location.protocol == "https:") {
  protocol = "wss";
}

let websocket;

function toggleLiveUploaderCount() {
  isLiveUploaderCountActive = !isLiveUploaderCountActive;
  if (isLiveUploaderCountActive) {
    document.getElementById("liveUploaderBtn").innerText = "Stop live EDDN Count";
    websocket = new WebSocket(`${protocol}://${window.location.hostname}:${window.location.port}/ws`);

    websocket.addEventListener("message", (e) => {
      if (isLiveUploaderCountActive) {
        if (!seenUploaders.includes(e.data)) {
          seenUploaders.push(e.data);
        }

        document.getElementById("liveUploaderText").innerText = `Live Uploaders: ${seenUploaders.length}`;
      }
    });
  } else {
    document.getElementById("liveUploaderBtn").innerText = "Start live EDDN Count";
    document.getElementById("liveUploaderText").innerText = `Live Uploaders: --`;
    seenUploaders = [];

    websocket.close();
  }
}
