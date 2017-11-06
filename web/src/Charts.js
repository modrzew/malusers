import React, { Component } from 'react';
import './Charts.css';
import { ChartsGetter } from './ChartsGetter';
import { Loader } from './Loader';

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

  handleAgeOnClick = sortBy => {
    this.setState({ sortBy: 'age' });
  };

  render() {
    let chart;
    const { category, sortBy } = this.state;

    if (category === 'anime' && sortBy === 'gender') {
      chart = <ChartsGetter />;
    } else if (category === 'anime' && sortBy === 'age') {
      chart = <Loader />;
    } else if (category === 'manga' && sortBy === 'gender') {
      chart = <Loader />;
    } else if (category === 'manga' && sortBy === 'age') {
      chart = <Loader />;
    }

    return (
      <div clasName="Charts">
        <div className="ChartsTitle">
          <span>
            Please select category and sorting method for chart display
          </span>
          <div className="ChartRow">
            <button className="AnimeButton" onClick={this.handleAnimeOnClick}>Anime</button>
            <button className="MangaButton" onClick={this.handleMangaOnClick}>Manga</button>
          </div>
          <div className="ChartRow">
            <button className="subButton" onClick={this.handleGenderOnClick}>by Gender</button>
            <button className="subButton" onClick={this.handleAgeOnClick}>by Age</button>
          </div>
        </div>
        <div className="ChartsContent">{chart}</div>
      </div>
    );
  }
}
