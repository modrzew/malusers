import React, { Component } from 'react';
import ChartistGraph from 'react-chartist';

export class ChartsAnime extends Component {
  render() {
    const gender = {
      labels: ['completed', 'dropped', 'total days'],
      series: [1, 2, 4]
    };
    return <ChartistGraph data={gender} />;
  }
}
