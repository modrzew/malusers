import React, { Component } from 'react';
import './App.css';

export class UserMangaRanking extends Component {
  render() {
    return (
      <div className="User-manga">
        <span className="User-completed">Completed</span> mangas gave&nbsp;
        sweetmonia <span className="User-rancompleted">#857</span> place in
        ranking, which is <span className="User-days">795</span>&nbsp;chapters
        and <span className="User-days">631</span> volumes read. sweetmonia
        didn't finished enough mangas that is&nbsp;
        <span className="User-randropped">#556</span> in&nbsp;
        <span className="User-dropped">dropped</span> category.
      </div>
    );
  }
}
