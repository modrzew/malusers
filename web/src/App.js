import React, { Component } from 'react';
import './App.css';
import { Search } from './Search';
import { GetUser } from './GetUser';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      userName: null
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({ value: event.target.value });
  }

  handleSubmit = userName => {
    this.setState({ userName: userName });
  };

  render() {
    const userName = this.state.userName;
    if (userName) {
      return <GetUser userName={this.state.userName} />;
    } else {
      return <Search onSubmit={this.handleSubmit} />;
    }
  }
}

export default App;
