import React, { Component } from 'react';
import './App.css';

export class Users extends Component {
  render() {
      return <div className="User-body">
        <div className="User-row">
          <div className="User-ranking">
            <div className="User-name">sweetmonia</div>
            <div className="User-update">last update: 2017-10-19</div>
            <div className="User-anime">sweetmonia is <span className="User-rancompleted">#305</span> ranked by number of <span className="User-completed">completed</span> anime and spent <span className="User-days">168,6</span> days of life to watch <span className="User-days">431</span> titles and <span className="User-days">10.174</span> episodes in total, making it <span className="User-randays">#203</span> place ranked by total lost days. sweetmonia decided that <span className="User-titlesdropped">43</span> titles are not worth of watching and give them <span className="User-dropped">dropped</span> status, with <span className="User-randropped">#217</span> place in this ranking.</div>
            <div className="User-manga"><span className="User-completed">Completed</span> mangas gave sweetmonia <span className="User-rancompleted">#857</span>place in ranking, which is <span className="User-days">795</span> chapters and <span className="User-days">631</span> volumes read. sweetmonia didn't finished enough mangas that is <span className="User-randropped">#556</span> in <span className="User-dropped">dropped</span> category.</div>
          </div>
        </div>
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
      </div>

       
  }
}