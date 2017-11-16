import React, { Component } from 'react';
import { Error } from './Error';
import AppStyles from './App.css';

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

  handleButtonOnClick = showComponent => {
    this.props.onButtonSubmit();
  };

  render() {
    return (
      <div className={AppStyles.App}>
        <button
          className={AppStyles.ChartButton}
          onClick={this.handleButtonOnClick}
        >
          Charts
        </button>
        <div className={AppStyles.SearchWrapped}>
          <div className={AppStyles.AppTitle}>
            Hi there, welcome to malusers, the ultimate MAL scrapper.
          </div>
          <div className={AppStyles.AppInstructions}>
            Please enter user name
          </div>
          <input
            type="text"
            placeholder="User name"
            value={this.state.value}
            onChange={this.handleChange}
          />
          <button
            className={AppStyles.SearchButton}
            onClick={this.handleOnClick}
          >
            Submit
          </button>
          {this.props.showError && <Error />}
        </div>
      </div>
    );
  }
}
