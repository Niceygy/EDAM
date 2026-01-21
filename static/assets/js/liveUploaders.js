let isLiveUploaderCountActive = false;
let message = 0;

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

    websocket.addEventListener("message", (_) => {
      if (isLiveUploaderCountActive) {
        message++;
        document.getElementById("liveUploaderText").innerText = `--: ${message}`;
      }
    });
  } else {
    document.getElementById("liveUploaderBtn").innerText = "Start live EDDN Count";
    document.getElementById("liveUploaderText").innerText = `Live Uploaders: --`;
    message = 0;

    websocket.close();
  }
}
