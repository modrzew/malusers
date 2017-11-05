import React, { Component } from 'react';
import ChartistGraph from 'react-chartist';
import './Charts.css';

export class ChartsAnime extends Component {
  render() {
    let gender = {
      labels: ['completed', 'dropped', 'total days'],
      series: [
        {
          'name': 'female',
          'data': [
            this.props.result.F.completed.count,
            this.props.result.F.dropped.count,
            this.props.result.F.total_days.count
          ]
        },
        [
          this.props.result.M.completed.count,
          this.props.result.M.dropped.count,
          this.props.result.M.total_days.count
        ],
        [
          this.props.result.X.completed.count,
          this.props.result.X.dropped.count,
          this.props.result.X.total_days.count
        ]
      ]
    };

    let options = {
      width: 800,
      height: 500,
      seriesBarDistance: 50,
      plugins: [Chartist.plugins.tooltip()],
      axisX: {
        position: 'end',
        showGridBackground: true
      },
      axisY: {
        offset: 80,
        position: 'start',
        labelOffset: {
          x: 10,
          y: 25
        }
      },
      chartPadding: {
        top: 20
      }
    };

    let type = 'Bar';

    return (
      <div className="chart">
        <div className="chartTitle">Gender</div> <ChartistGraph data={gender} type={type} options={options} />
      </div>
    );
  }
}
