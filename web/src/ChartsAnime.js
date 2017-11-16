import React, { Component } from 'react';
import { GenderChart } from './GenderChart';
import { YearChart } from './YearChart';
import './Charts.css';

const getYearData = (result, type) => {
  const list = [];
  const completed = [];
  const dropped = [];
  Object.keys(result).forEach(function(key) {
    list.push(key);
    completed.push(result[key]['completed'][type]);
    dropped.push(result[key]['dropped'][type]);
  });
  return { labels: list, completed: completed, dropped: dropped };
};

const getGenderData = (result, value, type) => {
  return {
    M: result.M[value][type],
    F: result.F[value][type],
    X: result.X[value][type],
    '': result[''][value][type]
  };
};

export class ChartsAnime extends Component {
  render() {
    if (this.props.subcat === 'gender') {
      return (
        <div className="Charts">
          <div className="ChartRow">
            <div className="ChartCompleted">
              <GenderChart
                title="Completed Anime by Gender"
                data={getGenderData(this.props.result, 'completed', 'count')}
              />
            </div>
            <div className="ChartCompleted">
              <GenderChart
                title="Completed Anime Mean Score by Gender"
                data={getGenderData(this.props.result, 'completed', 'mean')}
              />
            </div>
          </div>
          <div className="ChartRow">
            <div className="ChartDropped">
              <GenderChart
                title="Dropped Anime by Gender"
                data={getGenderData(this.props.result, 'dropped', 'count')}
              />
            </div>
            <div className="ChartDropped">
              <GenderChart
                title="Dropped Anime Mean Score by Gender"
                data={getGenderData(this.props.result, 'dropped', 'mean')}
              />
            </div>
          </div>
        </div>
      );
    } else {
      return (
        <div className="Charts">
          <div className="ChartRow">
            <div className="ChartCompleted">
              <YearChart
                title="Completed and Dropped Anime by Year of Birth"
                data={getYearData(this.props.result, 'count')}
              />
            </div>
            <div className="ChartCompleted">
              <YearChart
                title="Completed and Dropped Mean Score by Year of Birth"
                data={getYearData(this.props.result, 'mean')}
              />
            </div>
          </div>
        </div>
      );
    }
  }
}
