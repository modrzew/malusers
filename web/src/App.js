import React, { Component } from 'react';
import './App.css';
import { Search } from './Search';
import { GetUser } from './GetUser';
import { Charts } from './Charts';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      userName: null,
      error: null,
      showComponent: true
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
    const fakeData = {"":{"completed":{"count":2156859,"mean":177,"median":0},"dropped":{"count":131985,"mean":14,"median":0},"total_days":{"count":767194,"mean":62,"median":0}},"F":{"completed":{"count":11428568,"mean":122,"median":0},"dropped":{"count":724154,"mean":8,"median":0},"total_days":{"count":4402069,"mean":45,"median":0}},"M":{"completed":{"count":31745762,"mean":171,"median":0},"dropped":{"count":1718360,"mean":11,"median":0},"total_days":{"count":13033452,"mean":61,"median":0}},"X":{"completed":{"count":251502,"mean":136,"median":0},"dropped":{"count":16846,"mean":10,"median":0},"total_days":{"count":90929,"mean":47,"median":0}}};
    if (userName && !error) {
      return <GetUser userName={userName} onError={this.handleError} />;
    } else if (showComponent === true) return <Charts result={fakeData} />;
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