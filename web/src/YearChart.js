import React, { Component } from 'react';
import { Line } from 'react-chartjs-2';

export class YearChart extends Component {
  render() {
    const data = {
      labels: this.props.data.labels,
      datasets: [
        {
          label: 'completed',
          data: this.props.data.completed,
          borderColor: ['rgba(153, 196, 50, 1)'],
          backgroundColor: ['rgba(153, 196, 50, 1)'],
          fill: false,
          borderWidth: 2
        },
        {
          label: 'dropped',
          data: this.props.data.dropped,
          borderColor: ['rgba(255, 167, 50, 1)'],
          backgroundColor: ['rgba(255, 167, 50, 1)'],
          fill: false,
          borderWidth: 2
        }
      ]
    };
    const options = {
      title: {
        display: true,
        text: this.props.title,
        position: 'top',
        fontSize: 20,
      },
      responsive: true,
      maintainAspectRatio: false,
      elements: { point: { radius: 2 } },
      scales: {
        xAxes: [
          {
            display: true
          }
        ],
        yAxes: [
          {
            display: true
          }
        ]
      }
    };
    return <Line data={data} options={options} />;
  }
}
