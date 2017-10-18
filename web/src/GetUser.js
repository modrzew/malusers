import React, { Component } from 'react';
import { User } from './User';

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

    fetch(`https://api.mal.modriv.net/user/${this.props.userName}`)
      .then(resp => resp.json())
      .then(data => {
        this.setState({ name: data, isLoading: false });
      });
  }

  render() {
    const { name, isLoading } = this.state;
    if (isLoading) {
      return <p>Loading...</p>;
    } else if (name != null) {
      return <User name={this.state.name} />;
    } else {
      return <p>Loading...</p>;
    }
  }
}
