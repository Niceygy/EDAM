fetch("data/messageCount.csv")
  .then((response) => response.text())
  .then((csv) => {
    const lines = csv.trim().split("\n");
    const labels = [];
    const data = [];
    for (let i = 1; i < lines.length; i++) {
      const [unixtime, count] = lines[i].split(",");
      // Convert unixtime to readable date
      const date = new Date(parseInt(unixtime, 10) * 1000);
      labels.push(date.toLocaleString());
      data.push(parseInt(count, 10));
    }
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
