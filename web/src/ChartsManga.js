import React, { Component } from 'react';
import { GenderChart } from './GenderChart';
import { YearChart } from './YearChart';
import ChartStyles from './Charts.css';

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

export class ChartsManga extends Component {
  render() {
    if (this.props.subcat === 'gender') {
      return (
        <div className={ChartStyles.Charts}>
          <div className={ChartStyles.ChartRow}>
            <div className={ChartStyles.ChartCompleted}>
              <GenderChart
                title="Completed Manga by Gender"
                data={getGenderData(this.props.result, 'completed', 'count')}
              />
            </div>
            <div className={ChartStyles.ChartCompleted}>
              <GenderChart
                title="Completed Manga Mean Score by Gender"
                data={getGenderData(this.props.result, 'completed', 'mean')}
              />
            </div>
          </div>
          <div className={ChartStyles.ChartRow}>
            <div className={ChartStyles.ChartDropped}>
              <GenderChart
                title="Dropped Manga by Gender"
                data={getGenderData(this.props.result, 'dropped', 'count')}
              />
            </div>
            <div className={ChartStyles.ChartDropped}>
              <GenderChart
                title="Dropped Manga Mean Score by Gender"
                data={getGenderData(this.props.result, 'dropped', 'mean')}
              />
            </div>
          </div>
        </div>
      );
    } else {
      return (
        <div className={ChartStyles.Charts}>
          <div className={ChartStyles.ChartRow}>
            <div className={ChartStyles.ChartCompleted}>
              <YearChart
                title="Completed and Dropped Manga by Year of Birth"
                data={getYearData(this.props.result, 'count')}
              />
            </div>
            <div className={ChartStyles.ChartCompleted}>
              <YearChart
                title="Completed and Dropped Manga Score by Year of Birth"
                data={getYearData(this.props.result, 'mean')}
              />
            </div>
          </div>
        </div>
      );
    }
  }
}
