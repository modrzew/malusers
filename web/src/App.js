import React, { Component } from 'react';
import './App.css';
import { Search } from './Search';
import { GetUser } from './GetUser';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      userName: null,
      error: null
    };
  }

  handleChange = event => {
    this.setState({ value: event.target.value });
  };

  handleSubmit = userName => {
    this.setState({ userName: userName, error: null });
  };

  handleError = error => {
    this.setState({ error: error });
  };

  render() {
    const userName = this.state.userName;
    const error = this.state.error;
    if (userName && !error) {
      return <GetUser userName={userName} onError={this.handleError} />;
    } else {
      return <Search onSubmit={this.handleSubmit} showError={error} />;
    }
  }
}

export default App;
