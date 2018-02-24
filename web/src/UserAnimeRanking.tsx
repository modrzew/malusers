import * as React from 'react';
const CategoryRankingStyles = require('./CategoryRanking.css');
import * as classnames from 'classnames';
import * as numeral from 'numeral';

type Props = {
  animeRanking: any,
  animeStats: any,
  userName: string,
};

export class UserAnimeRanking extends React.Component<Props> {
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
      <div className={CategoryRankingStyles.categoryRanking}>
        {this.props.userName} is{' '}
        <span
          className={classnames(
            CategoryRankingStyles.ranking,
            CategoryRankingStyles.rankingCompleted
          )}
        >
          #{animeCompleted}
        </span>{' '}
        in{' '}
        <span
          className={classnames(
            CategoryRankingStyles.label,
            CategoryRankingStyles.labelCompleted
          )}
        >
          completed
        </span>{' '}
        anime ranking and spent{' '}
        <span
          className={classnames(
            CategoryRankingStyles.tag,
            CategoryRankingStyles.tagDays
          )}
        >
          {animeTagDays}
        </span>{' '}
        days of life to watch{' '}
        <span
          className={classnames(
            CategoryRankingStyles.tag,
            CategoryRankingStyles.tagTitles
          )}
        >
          {animeTagTitles}
        </span>{' '}
        titles and{' '}
        <span
          className={classnames(
            CategoryRankingStyles.tag,
            CategoryRankingStyles.tagEpisodes
          )}
        >
          {animeTagEpisodes}
        </span>{' '}
        episodes in total, making it{' '}
        <span
          className={classnames(
            CategoryRankingStyles.ranking,
            CategoryRankingStyles.rankingDays
          )}
        >
          #{animeDays}
        </span>{' '}
        place ranked by total{' '}
        <span
          className={classnames(
            CategoryRankingStyles.label,
            CategoryRankingStyles.labelDays
          )}
        >
          lost days
        </span>. {this.props.userName} decided that{' '}
        <span
          className={classnames(
            CategoryRankingStyles.tag,
            CategoryRankingStyles.tagTitlesDropped
          )}
        >
          {animeTitlesDropped}
        </span>{' '}
        titles are not worth of watching and gave them{' '}
        <span
          className={classnames(
            CategoryRankingStyles.label,
            CategoryRankingStyles.labelDropped
          )}
        >
          dropped
        </span>{' '}
        status, being{' '}
        <span
          className={classnames(
            CategoryRankingStyles.ranking,
            CategoryRankingStyles.rankingDropped
          )}
        >
          #{animeDropped}
        </span>{' '}
        in this ranking.
      </div>
    );
  }
}
