var eddnPlayerKnob = pureknob.createKnob(300, 300);
eddnPlayerKnob.setProperty("valMin", 0);
eddnPlayerKnob.setProperty("valMax", 10000);
eddnPlayerKnob.setProperty("angleStart", -0.75 * Math.PI);
eddnPlayerKnob.setProperty("angleEnd", 0.75 * Math.PI);
var node = eddnPlayerKnob.node();
var elem = document.getElementById("eddnNowGauge");
elem.appendChild(node);
fetch("/data/steamcount")
  .then((response) => console.log(response.status) || response) // output the status and return response
  .then((response) => response.text())
  .then((response) => eddnPlayerKnob.setValue(response));
