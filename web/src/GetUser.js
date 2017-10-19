import React, { Component } from 'react';
import { User } from './User';
import { Loader } from './Loader';

const API_URL = 'https://api.mal.modriv.net';

export class GetUser extends Component {
  constructor(props) {
    super(props);
    this.state = {
      name: null,
      isLoading: false
    };
  }

  componentDidMount() {
    this.setState({ isLoading: true });

    fetch(`${API_URL}/user/${this.props.userName}`)
      .then(resp => resp.json())
      .then(data => {
        this.setState({ name: data, isLoading: false });
      });
  }

  render() {
    const { name, isLoading } = this.state;
    if (isLoading) {
      return <Loader />;
    } else if (name !== null) {
      return <User name={this.state.name} />;
    } else {
      return <Loader />;
    }
  }
}
