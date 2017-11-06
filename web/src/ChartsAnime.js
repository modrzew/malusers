import React, { Component } from 'react';
import { GenderChart } from './GenderChart';
import { YearChart } from './YearChart';
import './Charts.css';

const getGenderData = (result, key, type) => {
  return {
    M: result.M[key][type],
    F: result.F[key][type],
    '': result[''][key][type],
    X: result.X[key][type]
  };
};

const getYearData = (result, key, type) => {
  return {    
  }
}

export class ChartsAnime extends Component {
  render() {
    const Chart = this.props.subcat === 'gender' ? GenderChart : YearChart;
    return (
      <div className="Charts">
        <div className="ChartRow">
          <div className="ChartCompleted">
            <Chart
              title="Completed Anime by Gender"
              data={getGenderData(this.props.result, 'completed', 'count')}
            />
          </div>
          <div className="ChartCompleted">
            <Chart
              title="Completed Mean Score by Gender"
              data={getGenderData(this.props.result, 'completed', 'mean')}
            />
          </div>
        </div>
        <div className="ChartRow">
          <div className="ChartDropped">
            <Chart
              title="Dropped Anime by Gender"
              data={getGenderData(this.props.result, 'dropped', 'count')}
            />
          </div>
          <div className="ChartDropped">
            <Chart
              title="Dropped Mean Score by Gender"
              data={getGenderData(this.props.result, 'dropped', 'mean')}
            />
          </div>
        </div>
      </div>
    );
  }
}
