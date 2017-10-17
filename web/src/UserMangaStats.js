import React, { Component } from 'react';
import './App.css';

export class UserMangaStats extends Component {
  render() {
    return (
      <div className="User-row">
        <div className="User-col">
          <div className="manga">manga</div>
        </div>
        <div className="User-col">
          <div className="User-category">in progress</div>
          <div className="mangaInProgress">012</div>
        </div>
        <div className="User-col">
          <div className="User-category">completed</div>
          <div className="mangaCompleted">007</div>
        </div>
        <div className="User-col">
          <div className="User-category">on hold</div>
          <div className="mangaOnHold">03202</div>
        </div>
        <div className="User-col">
          <div className="User-category">dropped</div>
          <div className="mangaDropped">0020</div>
        </div>
        <div className="User-col">
          <div className="User-category">planned</div>
          <div className="mangaPlanned">017</div>
        </div>
        <div className="User-col">
          <div className="User-category">reread</div>
          <div className="mangaRewatch">0200</div>
        </div>
        <div className="User-col">
          <div className="User-category">mean score</div>
          <div className="mangaMeanscore">7,41</div>
        </div>
      </div>
    );
  }
}
