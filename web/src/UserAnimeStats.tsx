import * as React from 'react';
const CategoryStyles = require('./Category.css');
const CategoryAnimeStyles = require('./CategoryAnime.css');
import * as classnames from 'classnames';

type Props = {
  animeStats: any,
};

export class UserAnimeStats extends React.Component<Props> {
  render() {
    return (
      <div className={CategoryStyles.CategoryRow}>
        <div className={CategoryStyles.CategoryCol}>
          <div
            className={classnames(
              CategoryStyles.category,
              CategoryStyles.animeBox
            )}
          >
            anime stats
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>in progress</div>
          <div
            className={classnames(
              CategoryAnimeStyles.anime,
              CategoryAnimeStyles.inProgress
            )}
          >
            {this.props.animeStats.inProgress}
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>completed</div>
          <div
            className={classnames(
              CategoryAnimeStyles.anime,
              CategoryAnimeStyles.completed
            )}
          >
            {this.props.animeStats.completed}
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>on hold</div>
          <div
            className={classnames(
              CategoryAnimeStyles.anime,
              CategoryAnimeStyles.onHold
            )}
          >
            {this.props.animeStats.onHold}
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>dropped</div>
          <div
            className={classnames(
              CategoryAnimeStyles.anime,
              CategoryAnimeStyles.dropped
            )}
          >
            {this.props.animeStats.dropped}
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>planned</div>
          <div
            className={classnames(
              CategoryAnimeStyles.anime,
              CategoryAnimeStyles.planned
            )}
          >
            {this.props.animeStats.planned}
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>rewatched</div>
          <div
            className={classnames(
              CategoryAnimeStyles.anime,
              CategoryAnimeStyles.rewatch
            )}
          >
            {this.props.animeStats.rewatched}
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>mean score</div>
          <div
            className={classnames(
              CategoryAnimeStyles.anime,
              CategoryAnimeStyles.meanScore
            )}
          >
            {this.props.animeStats.meanScore}
          </div>
        </div>
      </div>
    );
  }
}
