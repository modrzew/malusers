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
      <div className="App">
        <h1 className="AppTitle">
          Hi there, welcome to malusers the ultimate mal scrapper.
        </h1>
        <h2 className="AppInstructions">Please enter user name</h2>
        <input
          type="text"
          placeholder="User name"
          value={this.state.value}
          onChange={this.handleChange}
        />
        <button onClick={this.handleOnClick}>Submit</button>
      </div>
    );
  }
}
