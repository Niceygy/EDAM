fetch("data/messageCount.csv")
  .then((response) => response.text())
  .then((csv) => {
    const lines = csv.trim().split("\n");
    var labels = [];
    var data = [];
    // for (let i = 1; i < lines.length; i++) {
    //   const [unixtime, count] = lines[i].split(",");
    //   // Convert unixtime to readable date
    //   const date = new Date(parseInt(unixtime, 10) * 1000);
    //   labels.push(date.toLocaleString());
    //   data.push(parseInt(count, 10));
    // }
    labels, (data = getDataForxDays(1, lines));
    const ctx = document.getElementById("eddnMessagesChart").getContext("2d");
    new Chart(ctx, {
      type: "line",
      data: {
        labels: labels,
        datasets: [
          {
            label: "EDDN Messages / Hour",
            data: data,
            borderColor: "rgb(230, 129, 14)",
            borderWidth: 2,
            fill: false,
            pointRadius: 0,
          },
        ],
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          y: { beginAtZero: true },
          x: { display: true, title: { display: true, text: "Time" } },
        },
      },
    });
  });

function getDataForxDays(days, csv) {
  var labels = [];
  var data = [];
  // Source - https://stackoverflow.com/a/1296374
  // Posted by Stephen Wrighton, modified by community. See post 'Timeline' for change history
  // Retrieved 2025-12-30, License - CC BY-SA 4.0

  var xDaysAgo = new Date();
  xDaysAgo.setDate(xDaysAgo.getDate() - days);

  for (let i = 1; i < csv.length; i++) {
    const [unixtime, count] = csv[i].split(",");
    // Convert unixtime to readable date
    const date = new Date(parseInt(unixtime, 10) * 1000);
    if (date > xDaysAgo) {
      labels.push(date.toLocaleString());
      data.push(parseInt(count, 10));
    }
  }

  return labels, data;
}
