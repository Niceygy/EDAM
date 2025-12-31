var steamPlayerKnob = pureknob.createKnob(250, 250);
steamPlayerKnob.setProperty("valMin", 0);
steamPlayerKnob.setProperty("valMax", 20000);
steamPlayerKnob.setProperty("angleStart", -0.75 * Math.PI);
steamPlayerKnob.setProperty("angleEnd", 0.75 * Math.PI);
steamPlayerKnob.setProperty("readonly", true);
var node = steamPlayerKnob.node();
var elem = document.getElementById("steamPlayerGauge");
elem.appendChild(node);
fetch("/data/steamcount")
  .then((response) => console.log(response.status) || response) // output the status and return response
  .then((response) => response.text())
  .then((response) => steamPlayerKnob.setValue(response));
