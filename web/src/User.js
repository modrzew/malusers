import React, { Component } from 'react';
import './App.css';
import './User.css';
import { UserAnimeRanking } from './UserAnimeRanking';
import { UserMangaRanking } from './UserMangaRanking';
import { UserAnimeStats } from './UserAnimeStats';
import { UserMangaStats } from './UserMangaStats';
import moment from 'moment';

export class User extends Component {
  render() {
    const date = moment(this.props.name.lastUpdate).format(
      'dddd, MMMM Do YYYY, h:mm a'
    );
    return (
      <div className="userBody">
        <div className="userInfo">
          <div className="userName userStand">{this.props.name.username}</div>
          <div className="userUpdate userStand">last update: {date}</div>
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
