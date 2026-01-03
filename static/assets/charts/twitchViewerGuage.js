var twitchViewerKnob = pureknob.createKnob(250, 250);
twitchViewerKnob.setProperty("valMin", 0);
twitchViewerKnob.setProperty("valMax", 9000);
twitchViewerKnob.setProperty("angleStart", -0.75 * Math.PI);
twitchViewerKnob.setProperty("angleEnd", 0.75 * Math.PI);
twitchViewerKnob.setProperty("readonly", true);
var node = twitchViewerKnob.node();
var elem = document.getElementById("twitchViewerGauge");
elem.appendChild(node);
fetch("/data/twitchcount")
  .then((response) => console.log(response.status) || response) // output the status and return response
  .then((response) => response.text())
  .then((response) => twitchViewerKnob.setValue(response));
