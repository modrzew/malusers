import React, { Component } from 'react';
import { Doughnut } from 'react-chartjs-2';

export class GenderChart extends Component {
  render() {
    const data = {
      labels: ['Male', 'Female', 'Non-Binary', 'Not Specified'],
      datasets: [
        {
          data: [
            this.props.data.M,
            this.props.data.F,
            this.props.data.X,
            this.props.data['']
          ],
          backgroundColor: [
            'rgba(79, 181, 255, 1)',
            'rgba(153, 196, 50, 1)',
            'rgba(255, 167, 50, 1)',
            'rgba(191, 76, 187, 1)'
          ],
          borderColor: ['rgba(255, 255, 255, 1)'],
          borderWidth: 2
        }
      ]
    };
    const options = {
      title: {
        display: true,
        text: this.props.title,
        position: 'top',
        fontSize: 20
      },
      responsive: true,
      maintainAspectRatio: false
    };
    return <Doughnut data={data} options={options} />;
  }
}
