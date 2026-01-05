const activityRatings = ["Server hamsters must be asleep", "Off-Peak", "About normal", "Active! Have you checked galnet?", "Very active! Best check on salvation..."];

function setActivityMessage(rating) {
  rating = Math.round(rating);

  if (rating > 99) {
    document.getElementById("activityRatingMessage").innerText = `${activityRatings[4]} (${rating}% active)`;
  } else if (rating > 75) {
    document.getElementById("activityRatingMessage").innerText = `${activityRatings[3]} (${rating}% active)`;
  } else if (rating > 50) {
    document.getElementById("activityRatingMessage").innerText = `${activityRatings[2]} (${rating}% active)`;
  } else if (rating > 25) {
    document.getElementById("activityRatingMessage").innerText = `${activityRatings[1]} (${rating}% active)`;
  } else {
    document.getElementById("activityRatingMessage").innerText = `${activityRatings[0]} (${rating}% active)`;
  }
}

fetch("/data/activityrating")
  .then((response) => console.log(response.status) || response) // output the status and return response
  .then((response) => response.text())
  .then((response) => setActivityMessage(response));
