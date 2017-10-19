import React, { Component } from 'react';
import './CategoryRanking.css';
import numeral from 'numeral';
import 'numeral/locales';

numeral.locale('pl');

export class UserAnimeRanking extends Component {
  render() {
    const animeCompleted = numeral(this.props.animeRanking.completed).format(
      '0,0'
    );
    const animeDropped = numeral(this.props.animeRanking.dropped).format('0,0');
    const animeDays = numeral(this.props.animeRanking.totalDays).format('0,0');
    const animeTagDays = numeral(this.props.animeStats.totalDays).format('0,0');
    const animeTagTitles = numeral(this.props.animeStats.completed).format(
      '0,0'
    );
    const animeTagEpisodes = numeral(
      this.props.animeStats.totalEpisodes
    ).format('0,0');
    const animeTitlesDropped = numeral(this.props.animeStats.dropped).format(
      '0,0'
    );
    return (
      <div className="categoryRanking">
        {this.props.userName} is&nbsp;
        <span className="ranking rankingCompleted">{animeCompleted}</span>
        &nbsp;number of&nbsp;
        <span className="label labelCompleted">completed</span>
        &nbsp;anime and spent&nbsp;
        <span className="tag tagDays">{animeTagDays}</span>
        &nbsp;days of life to watch&nbsp;
        <span className="tag tagTitles">{animeTagTitles}</span> titles&nbsp;
        and&nbsp;
        <span className="tag tagEpisodes">{animeTagEpisodes}</span>
        &nbsp;episodes in total, making it&nbsp;
        <span className="ranking rankingDays">{animeDays}</span>
        &nbsp;place ranked by total lost days. {this.props.userName} decided
        &nbsp;that&nbsp;
        <span className="tag tagTitlesDropped">{animeTitlesDropped}</span>
        &nbsp;titles are not worth of watching and gave them&nbsp;
        <span className="label labelDropped">dropped</span>
        &nbsp;status, with&nbsp;
        <span className="ranking rankingDropped">{animeDropped}</span>
        &nbsp;in this ranking.
      </div>
    );
  }
}
