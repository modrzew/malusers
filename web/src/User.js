import React, { Component } from 'react';
import { UserAnimeRanking } from './UserAnimeRanking';
import { UserMangaRanking } from './UserMangaRanking';
import { UserAnimeStats } from './UserAnimeStats';
import { UserMangaStats } from './UserMangaStats';
import moment from 'moment';
import UserStyles from './User.css';
import * as classnames from 'classnames';

export class User extends Component {
  render() {
    const date = moment(this.props.name.lastUpdate).format(
      'dddd, MMMM Do YYYY, h:mm a'
    );
    return (
      <div className={UserStyles.userBody}>
        <div className={UserStyles.userInfo}>
          <div
            className={classnames(UserStyles.userName, UserStyles.userStand)}
          >
            {this.props.name.username}
          </div>
          <div
            className={classnames(UserStyles.userUpdate, UserStyles.userStand)}
          >
            last update: {date}
          </div>
          <UserAnimeRanking
            userName={this.props.name.username}
            animeRanking={this.props.name.ranking.anime}
            animeStats={this.props.name.animeStats}
          />
          <UserMangaRanking
            userName={this.props.name.username}
            mangaRanking={this.props.name.ranking.manga}
            mangaStats={this.props.name.mangaStats}
          />
        </div>
        <UserAnimeStats animeStats={this.props.name.animeStats} />
        <UserMangaStats mangaStats={this.props.name.mangaStats} />
      </div>
    );
  }
}
