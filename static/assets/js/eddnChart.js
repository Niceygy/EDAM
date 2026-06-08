fetch("/data/eddncsv")
  .then((response) => response.text())
  .then((csv) => {
    const lines = csv.trim().split("\n");

    let dateRange = "";
    let url = window.location.href;
    if (url.includes("?range=")) {
      dateRange = url.split("?range=")[1].replace("d", "");
      if (dateRange == "all") {
        dateRange = "365";
      }

      document.getElementById("dateSelect").value = url.split("?range=")[1];
    } else {
      dateRange = "1";
    }

    let tmp = getDataForxDays(dateRange, lines);
    let labels = tmp[0];
    let data = tmp[1];

    console.log(`labels=${labels}`);
    console.log(`data=${data}`);

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
        plugins: {
          tooltip: {
            enabled: true,
            callbacks: {
              label: function (context) {
                return context.dataset.label + ": " + context.parsed.y;
              },
            },
          },
        },
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

    console.log(`count = ${count}, day = ${date.toDateString()}`);
    if (date > xDaysAgo) {
      labels.push(date.toLocaleString());
      data.push(parseInt(count, 10));
      if (parseInt(count, 10) > 100) {
        console.log("------------");
        console.log(`count=${count}`);
        console.log(`parseInt(count, 10)=${parseInt(count, 10)}`);
        console.log(`i=${i}`);
        console.log(`csv[i]=${csv[i]}`);
        console.log("------------");
      }
    }
  }

  labels.reverse();
  data.reverse();

  return [labels, data];
}
