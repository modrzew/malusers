import React, { Component } from 'react';
import './CategoryRanking.css';
import numeral from 'numeral';

export class UserMangaRanking extends Component {
  render() {
    const mangaCompleted = numeral(this.props.mangaRanking.completed).format(
      '0,0'
    );
    const mangaDropped = numeral(this.props.mangaRanking.dropped).format('0,0');
    const mangaChapters = numeral(this.props.mangaStats.totalChapters).format(
      '0,0'
    );
    const mangaVolumes = numeral(this.props.mangaStats.totalVolumes).format(
      '0,0'
    );
    const mangaTagDays = numeral(this.props.mangaStats.totalDays).format('0,0');
    const mangaTitlesDropped = numeral(this.props.mangaStats.dropped).format(
      '0,0'
    );
    return (
      <div className="categoryRanking">
        <span className="label labelCompleted">Completed</span> manga gave{' '}
        {this.props.userName}{' '}
        <span className="ranking rankingCompleted">#{mangaCompleted}</span>{' '}
        place in this ranking. It means that {this.props.userName} read{' '}
        <span className="tag tagChapters">{mangaChapters}</span> chapters and{' '}
        <span className="tag tagVolumes">{mangaVolumes}</span> manga volumes
        already, spending <span className="tag tagDays">
          {mangaTagDays}
        </span>{' '}
        <span className="label labelDays">days</span> to do it.{' '}
        <span className="ranking rankingDropped">#{mangaDropped}</span> place in{' '}
        <span className="label labelDropped">dropped</span> category means that{' '}
        <span className="tag tagTitlesDropped">{mangaTitlesDropped}</span>{' '}
        unread manga is enough for this rank.
      </div>
    );
  }
}
