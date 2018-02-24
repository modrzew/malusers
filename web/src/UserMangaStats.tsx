import * as React from 'react';
const CategoryStyles = require('./Category.css');
const CategoryMangaStyles = require('./CategoryManga.css');
import * as classnames from 'classnames';

type Props = {
  mangaStats: any;
};

export class UserMangaStats extends React.Component<Props> {
  render() {
    return (
      <div className={CategoryStyles.CategoryRow}>
        <div className={CategoryStyles.CategoryCol}>
          <div
            className={classnames(
              CategoryStyles.category,
              CategoryStyles.mangaBox
            )}
          >
            manga stats
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>in progress</div>
          <div
            className={classnames(
              CategoryMangaStyles.manga,
              CategoryMangaStyles.inProgress
            )}
          >
            {this.props.mangaStats.inProgress}
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>completed</div>
          <div
            className={classnames(
              CategoryMangaStyles.manga,
              CategoryMangaStyles.completed
            )}
          >
            {this.props.mangaStats.completed}
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>on hold</div>
          <div
            className={classnames(
              CategoryMangaStyles.manga,
              CategoryMangaStyles.onHold
            )}
          >
            {this.props.mangaStats.onHold}
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>dropped</div>
          <div
            className={classnames(
              CategoryMangaStyles.manga,
              CategoryMangaStyles.dropped
            )}
          >
            {this.props.mangaStats.dropped}
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>planned</div>
          <div
            className={classnames(
              CategoryMangaStyles.manga,
              CategoryMangaStyles.planned
            )}
          >
            {this.props.mangaStats.planned}
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>reread</div>
          <div
            className={classnames(
              CategoryMangaStyles.manga,
              CategoryMangaStyles.rewatch
            )}
          >
            {this.props.mangaStats.rewatched}
          </div>
        </div>
        <div className={CategoryStyles.CategoryCol}>
          <div className={CategoryStyles.subcategory}>mean score</div>
          <div
            className={classnames(
              CategoryMangaStyles.manga,
              CategoryMangaStyles.meanScore
            )}
          >
            {this.props.mangaStats.meanScore}
          </div>
        </div>
      </div>
    );
  }
}
