const activityRatings = ["Inactive", "People are asleep! Or on holiday", "About normal", "Active! Have you checked galnet?", "Very active! Best check on the thargoids"];

function setActivityMessage(rating) {
  var message = "";

  if (rating > 99) {
    document.querySelector("body > div.main-flex > div > div:nth-child(2) > h3.activityRatingMessage").innerText = activityRatings[4];
  } else if (rating > 75) {
    document.querySelector("body > div.main-flex > div > div:nth-child(2) > h3.activityRatingMessage").innerText = activityRatings[3];
  } else if (rating > 50) {
    document.querySelector("body > div.main-flex > div > div:nth-child(2) > h3.activityRatingMessage").innerText = activityRatings[2];
  } else if (rating > 25) {
    document.querySelector("body > div.main-flex > div > div:nth-child(2) > h3.activityRatingMessage").innerText = activityRatings[1];
  } else {
    document.querySelector("body > div.main-flex > div > div:nth-child(2) > h3.activityRatingMessage").innerText = activityRatings[0];
  }
}
