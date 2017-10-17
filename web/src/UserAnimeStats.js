import React, { Component } from 'react';
import './App.css';

export class UserAnimeStats extends Component {
  render() {
    return (
      <div className="User-row">
        <div className="User-col">
          <div className="anime">anime</div>
        </div>
        <div className="User-col">
          <div className="User-category">in progress</div>
          <div className="animeInProgress">22012</div>
        </div>
        <div className="User-col">
          <div className="User-category">completed</div>
          <div className="animeCompleted">22431</div>
        </div>
        <div className="User-col">
          <div className="User-category">on hold</div>
          <div className="animeOnHold">03330</div>
        </div>
        <div className="User-col">
          <div className="User-category">dropped</div>
          <div className="animeDropped">04333</div>
        </div>
        <div className="User-col">
          <div className="User-category">planned</div>
          <div className="animePlanned">03317</div>
        </div>
        <div className="User-col">
          <div className="User-category">rewatched</div>
          <div className="animeRewatch">04233</div>
        </div>
        <div className="User-col">
          <div className="User-category">mean score</div>
          <div className="animeMeanscore">6,39</div>
        </div>
      </div>
    );
  }
}
