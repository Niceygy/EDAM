var overallKnob = pureknob.createKnob(300, 300);
overallKnob.setProperty("valMin", 0);
overallKnob.setProperty("valMax", 100);
overallKnob.setProperty("angleStart", -0.75 * Math.PI);
overallKnob.setProperty("angleEnd", 0.75 * Math.PI);
overallKnob.setProperty("readonly", true);
var node = overallKnob.node();
var elem = document.getElementById("overallGauge");
elem.appendChild(node);
fetch("/data/activityrating")
  .then((response) => console.log(response.status) || response) // output the status and return response
  .then((response) => response.text())
  .then((response) => overallKnob.setValue(response) || response);
// .then((response) => overallKnob.setProperty("label", `${Math.round(response)}%`));
