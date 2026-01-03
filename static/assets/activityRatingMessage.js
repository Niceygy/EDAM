const activityRatings = ["Inactive", "People are asleep! Or on holiday", "About normal", "Active! Have you checked galnet?", "Very active! Best check on the thargoids"];

function setActivityMessage(rating) {
  var message = "";

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
