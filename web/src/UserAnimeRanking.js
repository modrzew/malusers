import React, { Component } from 'react';
import './CategoryRanking.css';
import numeral from 'numeral';

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
        {this.props.userName} is{' '}
        <span className="ranking rankingCompleted">#{animeCompleted}</span> in{' '}
        <span className="label labelCompleted">completed</span> anime ranking
        and spent <span className="tag tagDays">{animeTagDays}</span> days of{' '}
        life to watch <span className="tag tagTitles">
          {animeTagTitles}
        </span>{' '}
        titles and <span className="tag tagEpisodes">{animeTagEpisodes}</span> {' '}
        episodes in total, making it{' '}
        <span className="ranking rankingDays">#{animeDays}</span> place ranked
        by total <span className="label labelDays">lost days</span>.{' '}
        {this.props.userName} decided that{' '}
        <span className="tag tagTitlesDropped">{animeTitlesDropped}</span>{' '}
        titles are not worth of watching and gave them{' '}
        <span className="label labelDropped">dropped</span> status, being{' '}
        <span className="ranking rankingDropped">#{animeDropped}</span> in this{' '}
        ranking.
      </div>
    );
  }
}
