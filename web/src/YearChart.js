import React, { Component } from 'react';
import { Line } from 'react-chartjs-2';

export class YearChart extends Component {
  render() {
    const data = {
      labels: this.props.data.labels,
      datasets: [
        {
          label: 'completed',
          data: this.props.data.completed
        },
        {
          label: 'dropped',
          data: this.props.data.dropped
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
      maintainAspectRatio: true,
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
