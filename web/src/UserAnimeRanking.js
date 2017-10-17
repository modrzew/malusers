import React, { Component } from 'react';
import './App.css';

export class UserAnimeRanking extends Component {
  render() {
    return (
      <div className="User-anime">
        sweetmonia is <span className="User-rancompleted">#305</span> ranked by
        number of <span className="User-completed">completed</span> anime and
        spent <span className="User-days">168,6</span> days of life to
        watch&nbsp;
        <span className="User-days">431</span> titles and&nbsp;
        <span className="User-days">10.174</span> episodes in total,&nbsp;
        making it <span className="User-randays">#203</span> place ranked by
        total lost days. sweetmonia decided that&nbsp;
        <span className="User-titlesdropped">43</span> titles are not worth of
        watching and give them <span className="User-dropped">dropped</span>
        &nbsp; status, with <span className="User-randropped">#217</span>
        this ranking.
      </div>
    );
  }
}
