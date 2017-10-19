import React, { Component } from 'react';
import './Category.css';
import './CategoryAnime.css';

export class UserAnimeStats extends Component {
  render() {
    return (
      <div className="Category-row">
        <div className="Category-col">
          <div className="category animeBox">anime stats</div>
        </div>
        <div className="Category-col">
          <div className="subcategory">in progress</div>
          <div className="anime inProgress">
            {this.props.animeStats.inProgress}
          </div>
        </div>
        <div className="Category-col">
          <div className="subcategory">completed</div>
          <div className="anime completed">
            {this.props.animeStats.completed}
          </div>
        </div>
        <div className="Category-col">
          <div className="subcategory">on hold</div>
          <div className="anime onHold">{this.props.animeStats.onHold}</div>
        </div>
        <div className="Category-col">
          <div className="subcategory">dropped</div>
          <div className="anime dropped">{this.props.animeStats.dropped}</div>
        </div>
        <div className="Category-col">
          <div className="subcategory">planned</div>
          <div className="anime planned">{this.props.animeStats.planned}</div>
        </div>
        <div className="Category-col">
          <div className="subcategory">rewatched</div>
          <div className="anime rewatch">{this.props.animeStats.rewatched}</div>
        </div>
        <div className="Category-col">
          <div className="subcategory">mean score</div>
          <div className="anime meanScore">
            {this.props.animeStats.meanScore}
          </div>
        </div>
      </div>
    );
  }
}
