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
        <span className="label labelCompleted">Completed</span> manga gave{''}
        {this.props.userName}&nbsp;
        <span className="ranking rankingCompleted">#{mangaCompleted}</span>
        &nbsp;place in this ranking. It means that {this.props.userName}
        &nbsp;read&nbsp;
        <span className="tag tagChapters">{mangaChapters}</span>
        &nbsp;chapters and&nbsp;
        <span className="tag tagVolumes">{mangaVolumes}</span>
        &nbsp;manga volumes already, spending&nbsp;
        <span className="tag tagDays">{mangaTagDays}</span>&nbsp;
        <span className="label labelDays">days</span> to do it.&nbsp;
        <span className="ranking rankingDropped">#{mangaDropped}</span>
        &nbsp;place in <span className="label labelDropped">dropped</span>
        &nbsp;category means that&nbsp;
        <span className="tag tagTitlesDropped">{mangaTitlesDropped}</span>
        &nbsp;unread manga is enough for this rank.
      </div>
    );
  }
}
