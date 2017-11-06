import React, { Component } from 'react';
import { Loader } from './Loader';
import { ChartsAnime } from './ChartsAnime';

const API_URL_ANIME = 'https://api.mal.modriv.net/stats/anime/gender';
const API_URL_MANGA = 'https://api.mal.modriv.net/stats/manga';

export class ChartsGetter extends Component {
  constructor(props) {
    super(props);
    this.state = {
      result: null,
      isLoading: false
    };
  }
  componentDidMount() {
    this.setState({ isLoading: true });
    fetch(`${API_URL_ANIME}`)
      .then(resp => resp.json())
      .then(data => {
        this.setState({ result: data, isLoading: false });
      });
  }

  render() {
    const { result, isLoading } = this.state;
    if (isLoading) {
      return <Loader />;
    } else if (result !== null) {
      return <ChartsAnime result={this.state.result} />;
    } else {
      return <Loader />;
    }
  }
}