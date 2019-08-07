//Timer Loop to pull data every

var ctx = document.getElementById('topChart').getContext('2d');
var topChartJS = new Chart(ctx, {
  type: 'line',
  data: {
    labels: ["no Data available yet"],
    datasets: [{
      label: 'Sensor Packages per Hour',
      data: [314],
      backgroundColor: [
        'rgba(255, 99, 132, 0.2)',
      ],
      borderColor: [
        'rgba(255, 99, 132, 1)',
      ],
      borderWidth: 1
    }]
  },
  options: {
    scales: {
      yAxes: [{
        ticks: {
          beginAtZero: true
        }
      }]
    }
  }
});

var ctx = document.getElementById('bottomChart').getContext('2d');
var bottomChartJS = new Chart(ctx, {
  type: 'line',
  data: {
    labels: ["no Data available yet"],
    datasets: [{
      label: 'Sensor Packages per Day',
      data: [314],
      backgroundColor: [
        'rgba(54, 162, 235, 0.2)'
      ],
      borderColor: [
        'rgba(54, 162, 235, 1)'
      ],
      fill: true,
      borderWidth: 1
    }]
  },
  options: {
    scales: {
      yAxes: [{
        ticks: {
          beginAtZero: true
        }
      }]
    }
  }
});




setInterval(timerLoop, 3000);

function timerLoop()
{
  // console.log("PullTimerLoop-File exec");
  updateWindowsVar = JSON.parse(httpGet("updatewindows"));
  // console.log(updateWindowsVar);
  updateWindows(updateWindowsVar);
  updateTopChartVar = JSON.parse(httpGet("updatetopchart"));
  updateTopChart(updateTopChartVar);
  updateBottomChartVar = JSON.parse(httpGet("updatebottomchart"));
  updateBottomChart("bottomChart", updateBottomChartVar);
}


function httpGet(theUrl)
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", theUrl, false ); // false for synchronous request
    xmlHttp.send( null );
    return xmlHttp.responseText;
}

function updateWindows(WindowStatus)
{
    statBLi = WindowStatus.BakerStrFensterLi;
    console.log(statBLi);
    statBRe = WindowStatus.BakerStrFensterRe;
    console.log(statBRe);
    statKLi = WindowStatus.KuecheFensterLi;
    console.log(statKLi);
    statKRe = WindowStatus.KuecheFensterRe;
    console.log(statKRe);
    if (statKLi) {
      document.getElementById("KF-Li-Offen").style.opacity = 1;
    } else {
      document.getElementById("KF-Li-Offen").style.opacity = 0;
    }
    if (statKRe) {
      document.getElementById("KF-Re-Offen").style.opacity = 1;
    } else {
      document.getElementById("KF-Re-Offen").style.opacity = 0;
    }
    if (statBLi) {
      document.getElementById("BSF-Li-Offen").style.opacity = 1;
    } else {
      document.getElementById("BSF-Li-Offen").style.opacity = 0;
    }
    if (statBRe) {
      document.getElementById("BSF-Re-Offen").style.opacity = 1;
    } else {
      document.getElementById("BSF-Re-Offen").style.opacity = 0;
    }
    return;
}

function updateTopChart(updateTopChart)
{
  // chart = document.getElementById("topChart").topChartJS.data.labels;
  // console.log(chart)
  topChartJS.data.labels = updateTopChart.HourMatrix;
  topChartJS.data.datasets[0].data = updateTopChart.FlowMatrix;
  topChartJS.update();
}

function updateBottomChart(chart, updateBottomChart)
{
  bottomChartJS.data.labels = updateBottomChart.DayMatrix;
  bottomChartJS.data.datasets[0].data = updateBottomChart.FlowMatrix;
  bottomChartJS.update();
}

function addData(chart, label, data) {
    chart.data.labels.push(label);
    chart.data.datasets.forEach((dataset) => {
        dataset.data.push(data);
    });
    chart.update();
}

function removeData(chart) {
    chart.data.labels.pop();
    chart.data.datasets.forEach((dataset) => {
        dataset.data.pop();
    });
    chart.update();
}
