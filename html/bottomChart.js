var ctx = document.getElementById('bottomChart').getContext('2d');
var bottomChart = new Chart(ctx, {
  type: 'line',
  data: {
    // labels: [''],
    datasets: [{
      label: 'Sensor Packages per Day',
      data: [
        {x: 0 , y: 3},
        {x: 1 , y: 8}
      ],
      backgroundColor: [
        'rgba(54, 162, 235, 0.2)'
      ],
      borderColor: [
        'rgba(54, 162, 235, 1)'
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
