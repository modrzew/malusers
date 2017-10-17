import React, { Component } from 'react';
import './App.css';
import './AppManga.css';
import './AppAnime.css';
import './AppUserRanking.css';
import { UserAnimeRanking } from './UserAnimeRanking';
import { UserMangaRanking } from './UserMangaRanking';
import { UserAnimeStats } from './UserAnimeStats';
import { UserMangaStats } from './UserMangaStats';

export class Users extends Component {
  render() {
    return (
      <div className="User-body">
        <div className="User-row">
          <div className="User-info">
            <div className="User-name">sweetmonia</div>
            <div className="User-update">last update: 2017-10-19</div>
            <UserAnimeRanking />
            <UserMangaRanking />
          </div>
        </div>
        <UserAnimeStats />
        <UserMangaStats />
      </div>
    );
  }
}
