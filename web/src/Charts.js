import React, { Component } from 'react';
import ChartStyles from './Charts.css';
import { ChartsGetter } from './ChartsGetter';

export class Charts extends Component {
  constructor(props) {
    super(props);
    this.state = {
      category: null,
      sortBy: null
    };
  }

  handleAnimeOnClick = category => {
    this.setState({ category: 'anime' });
  };

  handleMangaOnClick = category => {
    this.setState({ category: 'manga' });
  };

  handleGenderOnClick = sortBy => {
    this.setState({ sortBy: 'gender' });
  };

  handleYearOnClick = sortBy => {
    this.setState({ sortBy: 'year' });
  };

  render() {
    let chart;
    const { category, sortBy } = this.state;

    if (category === 'anime' && sortBy === 'gender') {
      chart = (
        <ChartsGetter cat={this.state.category} subcat={this.state.sortBy} />
      );
    } else if (category === 'anime' && sortBy === 'year') {
      chart = (
        <ChartsGetter cat={this.state.category} subcat={this.state.sortBy} />
      );
    } else if (category === 'manga' && sortBy === 'gender') {
      chart = (
        <ChartsGetter cat={this.state.category} subcat={this.state.sortBy} />
      );
    } else if (category === 'manga' && sortBy === 'year') {
      chart = (
        <ChartsGetter cat={this.state.category} subcat={this.state.sortBy} />
      );
    }

    return (
      <div className={ChartStyles.Charts}>
        <div className={ChartStyles.ChartsTitle}>
          <span>
            Please select category and sorting method for chart display
          </span>
          <div className={ChartStyles.ChartRow}>
            <button
              className={ChartStyles.CategoryButton}
              onClick={this.handleAnimeOnClick}
            >
              Anime
            </button>
            <button
              className={ChartStyles.CategoryButton}
              onClick={this.handleMangaOnClick}
            >
              Manga
            </button>
          </div>
          <div className={ChartStyles.ChartRow}>
            <button
              className={ChartStyles.subButton}
              onClick={this.handleGenderOnClick}
            >
              by Gender
            </button>
            <button
              className={ChartStyles.subButton}
              onClick={this.handleYearOnClick}
            >
              by Year
            </button>
          </div>
        </div>
        <div className={ChartStyles.ChartsContent}>{chart}</div>
      </div>
    );
  }
}
