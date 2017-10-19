import React, { Component } from 'react';
import './Category.css';
import './CategoryManga.css';

export class UserMangaStats extends Component {
  render() {
    return (
      <div className="Category-row">
        <div className="Category-col">
          <div className="category mangaBox">manga stats</div>
        </div>
        <div className="Category-col">
          <div className="subcategory">in progress</div>
          <div className="manga inProgress">
            {this.props.mangaStats.inProgress}
          </div>
        </div>
        <div className="Category-col">
          <div className="subcategory">completed</div>
          <div className="manga completed">
            {this.props.mangaStats.completed}
          </div>
        </div>
        <div className="Category-col">
          <div className="subcategory">on hold</div>
          <div className="manga onHold">{this.props.mangaStats.onHold}</div>
        </div>
        <div className="Category-col">
          <div className="subcategory">dropped</div>
          <div className="manga dropped">{this.props.mangaStats.dropped}</div>
        </div>
        <div className="Category-col">
          <div className="subcategory">planned</div>
          <div className="manga planned">{this.props.mangaStats.planned}</div>
        </div>
        <div className="Category-col">
          <div className="subcategory">reread</div>
          <div className="manga rewatch">{this.props.mangaStats.rewatched}</div>
        </div>
        <div className="Category-col">
          <div className="subcategory">mean score</div>
          <div className="manga meanScore">
            {this.props.mangaStats.meanScore}
          </div>
        </div>
      </div>
    );
  }
}
