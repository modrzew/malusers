import React, { Component } from 'react';
import './CategoryRanking.css';

export class UserAnimeRanking extends Component {
  render() {
    return (
      <div className="categoryRanking">
        {this.props.userName} is&nbsp;
        <span className="ranking rankingCompleted">
          #{this.props.animeRanking.completed}
        </span>
        &nbsp;number of&nbsp;
        <span className="label labelCompleted">completed</span>
        &nbsp;anime and spent&nbsp;
        <span className="tag tagDays">{this.props.animeStats.totalDays}</span>
        &nbsp;days of life to watch&nbsp;
        <span className="tag tagTitles">
          {this.props.animeStats.completed}
        </span>&nbsp;titles and&nbsp;
        <span className="tag tagEpisodes">
          {this.props.animeStats.totalEpisodes}
        </span>
        &nbsp;episodes in total, making it&nbsp;
        <span className="ranking rankingDays">
          #{this.props.animeRanking.totalDays}
        </span>
        &nbsp;place ranked by total lost days. {this.props.userName} decided
        &nbsp;that&nbsp;
        <span className="tag tagTitlesDropped">
          {this.props.animeStats.dropped}
        </span>
        &nbsp;titles are not worth of watching and gave them&nbsp;
        <span className="label labelDropped">dropped</span>
        &nbsp;status, with&nbsp;
        <span className="ranking rankingDropped">
          #{this.props.animeRanking.dropped}
        </span>
        &nbsp;in this ranking.
      </div>
    );
  }
}
