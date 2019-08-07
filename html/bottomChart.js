var ctx = document.getElementById('bottomChart').getContext('2d');
var bottomChartJS = new Chart(ctx, {
  type: 'line',
  data: {
    labels: [1500,1600,1700,1750,1800,1850,1900,1950,1999,2050],
    datasets: [{
      label: 'Sensor Packages per Day',
      data: [282,350,411,502,635,809,947,1402,3700,5267],
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
