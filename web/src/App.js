import React, { Component } from 'react';
import { Search } from './Search';
import { GetUser } from './GetUser';
import { Charts } from './Charts';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      userName: null,
      error: null,
      showComponent: false
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

  handleButtonSubmit = showComponent => {
    this.setState({ showComponent: true });
  };

  render() {
    const userName = this.state.userName;
    const error = this.state.error;
    const showComponent = this.state.showComponent;
    if (userName && !error) {
      return <GetUser userName={userName} onError={this.handleError} />;
    } else if (showComponent === true) return <Charts />;
    else {
      return (
        <Search
          onSubmit={this.handleSubmit}
          showError={error}
          onButtonSubmit={this.handleButtonSubmit}
        />
      );
    }
  }
}

export default App;
