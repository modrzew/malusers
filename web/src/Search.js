import React, { Component } from 'react';
import './App.css';

export class Search extends Component {
  constructor(props) {
    super(props);
    this.state = { value: '' };
  }

  handleChange = event => {
    this.setState({ value: event.target.value });
  };

  handleOnClick = event => {
    this.props.onSubmit(this.state.value);
  };

  render() {
    return (
      <body className="body">
        <div className="App">
          <h1 className="App-title">
            Hi there, welcome to malusers the ultimate mal scrapper.
          </h1>
          <h2 className="App-instructions">Please enter user name</h2>
          <input
            type="text"
            placeholder="User name"
            value={this.state.value}
            onChange={this.handleChange}
          />
          <input type="submit" value="Submit" onClick={this.handleOnClick} />
        </div>
      </body>
    );
  }
}
