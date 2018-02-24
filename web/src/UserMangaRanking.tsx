import * as React from 'react';
const CategoryRankingStyles = require('./CategoryRanking.css');
import * as numeral from 'numeral';
import * as classnames from 'classnames';

type Props = {
  mangaRanking: any,
  mangaStats: any,
  userName: string,
}

export class UserMangaRanking extends React.Component<Props> {
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
      <div className={CategoryRankingStyles.categoryRanking}>
        <span
          className={classnames(
            CategoryRankingStyles.label,
            CategoryRankingStyles.labelCompleted
          )}
        >
          Completed
        </span>{' '}
        manga gave {this.props.userName}{' '}
        <span
          className={classnames(
            CategoryRankingStyles.ranking,
            CategoryRankingStyles.rankingCompleted
          )}
        >
          #{mangaCompleted}
        </span>{' '}
        place in this ranking. It means that {this.props.userName} read{' '}
        <span
          className={classnames(
            CategoryRankingStyles.tag,
            CategoryRankingStyles.tagChapters
          )}
        >
          {mangaChapters}
        </span>{' '}
        chapters and{' '}
        <span
          className={classnames(
            CategoryRankingStyles.tag,
            CategoryRankingStyles.tagVolumes
          )}
        >
          {mangaVolumes}
        </span>{' '}
        manga volumes already, spending{' '}
        <span
          className={classnames(
            CategoryRankingStyles.tag,
            CategoryRankingStyles.tagDays
          )}
        >
          {mangaTagDays}
        </span>{' '}
        <span
          className={classnames(
            CategoryRankingStyles.label,
            CategoryRankingStyles.labelDays
          )}
        >
          days
        </span>{' '}
        to do it.{' '}
        <span
          className={classnames(
            CategoryRankingStyles.ranking,
            CategoryRankingStyles.rankingDropped
          )}
        >
          #{mangaDropped}
        </span>{' '}
        place in{' '}
        <span
          className={classnames(
            CategoryRankingStyles.label,
            CategoryRankingStyles.labelDropped
          )}
        >
          dropped
        </span>{' '}
        category means that{' '}
        <span
          className={classnames(
            CategoryRankingStyles.tag,
            CategoryRankingStyles.tagTitlesDropped
          )}
        >
          {mangaTitlesDropped}
        </span>{' '}
        unread manga is enough for this rank.
      </div>
    );
  }
}
