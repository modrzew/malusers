import React, { Component } from 'react';
import { Loader } from './Loader';
import { ChartsAnime } from './ChartsAnime';
import { ChartsManga } from './ChartsManga';

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
    fetch(
      `https://api.mal.modriv.net/stats/${this.props.cat}/${this.props.subcat}`
    )
      .then(resp => resp.json())
      .then(data => {
        this.setState({ result: data, isLoading: false });
      });
  }

  render() {
    const { result, isLoading } = this.state;
    if (isLoading) {
      return <Loader />;
    } else if (result !== null && this.props.cat === 'anime') {
      return (
        <ChartsAnime result={this.state.result} subcat={this.props.subcat} />
      );
    } else if (result !== null && this.props.cat === 'manga') {
      return (
        <ChartsManga result={this.state.result} subcat={this.props.subcat} />
      );
    } else {
      return <Loader />;
    }
  }
}