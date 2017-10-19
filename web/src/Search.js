import React, { Component } from 'react';
import { Error } from './Error';
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
        <div className="AppTitle">
          Hi there, welcome to malusers, the ultimate MAL scrapper.
        </div>
        <div className="AppInstructions">Please enter user name</div>
        <input
          type="text"
          placeholder="User name"
          value={this.state.value}
          onChange={this.handleChange}
        />
        <button onClick={this.handleOnClick}>Submit</button>
        {this.props.showError && <Error />}
      </div>
    );
  }
}
