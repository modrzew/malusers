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
        &nbsp;place in this ranking. It means that {this.props.userName}
        &nbsp;read&nbsp;
        <span className="tag tagChapters">
          {this.props.mangaStats.totalChapters}
        </span>
        &nbsp;chapters and&nbsp;
        <span className="tag tagVolumes">
          {this.props.mangaStats.totalVolumes}
        </span>
        &nbsp;manga volumes already, spending&nbsp;
        <span className="tag tagDays">{this.props.mangaStats.totalDays}</span>
        &nbsp;days to do it.&nbsp;
        <span className="ranking rankingDropped">
          #{this.props.mangaRanking.dropped}
        </span>
        &nbsp;place in <span className="label labelDropped">dropped</span>
        &nbsp;category means that&nbsp;
        <span className="tag tagTitlesDropped">
          {this.props.mangaStats.dropped}
        </span>
        &nbsp;unread manga is enough for this rank.
      </div>
    );
  }
}
