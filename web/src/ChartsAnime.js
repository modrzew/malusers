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
    const meanCompletedData = {
      M: this.props.result.M.completed.mean,
      F: this.props.result.F.completed.mean,
      '': this.props.result[''].completed.mean,
      X: this.props.result.X.completed.mean
    };
    const meanDroppedData = {
      M: this.props.result.M.dropped.mean,
      F: this.props.result.F.dropped.mean,
      '': this.props.result[''].dropped.mean,
      X: this.props.result.X.dropped.mean
    };
    return (
      <div className="Charts">
        <div className="ChartRow">
        <div className="ChartCompleted">
          <GenderChart title="Completed Anime by Gender" data={completedData} />
        </div>
        <div className="ChartCompleted">
          <GenderChart title="Completed Mean Score by Gender" data={meanCompletedData} />
        </div>
        </div>
        <div className="ChartRow">
        <div className="ChartDropped">
          <GenderChart title="Dropped Anime by Gender" data={droppedData} />
        </div>
        <div className="ChartDropped">
          <GenderChart title="Dropped Mean Score by Gender" data={meanDroppedData} />
        </div>
        </div>
      </div>
    );
  }
}
