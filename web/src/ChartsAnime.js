import React, { Component } from 'react';
import { GenderChart } from './GenderChart';
import './Charts.css';

export class ChartsAnime extends Component {
  render() {
    const completedData = {
      M: this.props.result.M.completed.count,
      F: this.props.result.F.completed.count,
      '': this.props.result[''].completed.count,
      X: this.props.result.X.completed.count
    };
    const droppedData = {
      M: this.props.result.M.dropped.count,
      F: this.props.result.F.dropped.count,
      '': this.props.result[''].dropped.count,
      X: this.props.result.X.dropped.count
    };
    return (
      <div className="Charts">
        <div className="ChartsTitle">Gender Charts</div>
        <div className="ChartsContent">
          <div className="ChartCompleted">Completed Anime by Gender
            <GenderChart label="Completed" data={completedData} />
          </div>
          <div className="ChartDropped">Dropped Anime by Gender
            <GenderChart label="Dropped" data={droppedData} />
          </div>
        </div>
      </div>
    );
  }
}
