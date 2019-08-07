var ctx = document.getElementById('topChart').getContext('2d');
var topChart = new Chart(ctx, {
  type: 'line',
  data: {
    // labels: [''],
    datasets: [{
      label: 'Sensor Packages per Hour',
      data: [
        {x: 0 , y: 3},
        {x: 1 , y: 8}
      ],
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
