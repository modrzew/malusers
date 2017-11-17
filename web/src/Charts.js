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

  handleCategoryChange = event => {
    this.setState({ category: event.target.value });
  };

  handleSortbyChange = event => {
    this.setState({ sortBy: event.target.value });
  };

  render() {
    const { category, sortBy } = this.state;

    return (
      <div className={ChartStyles.Charts}>
        <div className={ChartStyles.ChartsItems}>
          <div className={ChartStyles.ChartsTitle}>malUsers</div>
          <div className={ChartStyles.ChartsHint}>
            Please select category and sorting method
          </div>
          <div className={ChartStyles.ChartSelector}>
            <div className={ChartStyles.ChartCategory}>Category</div>
            <div className={ChartStyles.ChartRow}>
              <label class="container">
                <input
                  type="radio"
                  name="category"
                  value="anime"
                  checked={category === 'anime'}
                  onChange={this.handleCategoryChange}
                />
                <span className="checkmark" />
                Anime
              </label>
              <label className="container">
                <input
                  type="radio"
                  name="category"
                  value="manga"
                  checked={category === 'manga'}
                  onChange={this.handleCategoryChange}
                />
                <span className="checkmark" />
                Manga
              </label>
            </div>
            <div className={ChartStyles.ChartCategory}>Sorted by</div>
            <div className={ChartStyles.ChartRow}>
              <label className="container">
                <input
                  type="radio"
                  name="sortedBy"
                  value="gender"
                  checked={sortBy === 'gender'}
                  onChange={this.handleSortbyChange}
                />
                <span className="checkmark" />
                Gender
              </label>
              <label className="container">
                <input
                  type="radio"
                  name="sortedBy"
                  value="year"
                  checked={sortBy === 'year'}
                  onChange={this.handleSortbyChange}
                />
                <span className="checkmark" />
                Year
              </label>
            </div>
          </div>
        </div>
        <div className={ChartStyles.ChartsItems}>
          <div
            className={ChartStyles.ChartsContent}
            onSubmit={this.handleSubmit}
          >
            {category &&
              sortBy && (
                <ChartsGetter
                  cat={this.state.category}
                  subcat={this.state.sortBy}
                />
              )}
          </div>
        </div>
      </div>
    );
  }
}
