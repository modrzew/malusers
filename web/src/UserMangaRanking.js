import React, { Component } from 'react';
import './CategoryRanking.css';

export class UserMangaRanking extends Component {
  render() {
    return (
      <div className="categoryRanking">
        <span className="label labelCompleted">Completed</span> manga gave&nbsp;
        {this.props.userName}&nbsp;
        <span className="ranking rankingCompleted">
          #{this.props.mangaRanking.completed}
        </span>
        &nbsp;ranking, which is&nbsp;
        <span className="tag tagChapters">
          {this.props.mangaStats.totalChapters}
        </span>
        &nbsp;chapters and&nbsp;
        <span className="tag tagVolumes">
          {this.props.mangaStats.totalVolumes}
        </span>
        &nbsp;volumes read. {this.props.userName} didn't finish enough
        &nbsp;manga that is&nbsp;
        <span className="ranking rankingDropped">
          #{this.props.mangaRanking.dropped}
        </span>
        &nbsp;in <span className="label labelDropped">dropped</span> category.
      </div>
    );
  }
}
