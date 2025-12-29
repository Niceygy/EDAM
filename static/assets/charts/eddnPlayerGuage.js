var eddnPlayerKnob = pureknob.createKnob(300, 300);
eddnPlayerKnob.setProperty("valMin", 0);
eddnPlayerKnob.setProperty("valMax", 30000);
eddnPlayerKnob.setProperty("angleStart", -0.75 * Math.PI);
eddnPlayerKnob.setProperty("angleEnd", 0.75 * Math.PI);
eddnPlayerKnob.setProperty("readonly", true);
var node = eddnPlayerKnob.node();
var elem = document.getElementById("eddnNowGauge");
elem.appendChild(node);
fetch("/data/eddncount")
  .then((response) => console.log(response.status) || response) // output the status and return response
  .then((response) => response.text())
  .then((response) => eddnPlayerKnob.setValue(response));
