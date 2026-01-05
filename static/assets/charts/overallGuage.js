fetch("/data/activityrating")
  .then((response) => console.log(response.status) || response) // output the status and return response
  .then((response) => response.text())
  .then((response) => setActivityMessage(response));
