import React, { Component } from 'react';
import { Error } from './Error';
import SearchStyles from './Search.css';

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
      <div className={SearchStyles.Search}>
        <button
          className={SearchStyles.chartButton}
          onClick={this.handleButtonOnClick}
        >
          Charts
        </button>
        <div className={SearchStyles.SearchWrapped}>
          <div className={SearchStyles.SearchTitle}>
            Hi there, welcome to malusers, the ultimate MAL scrapper.
          </div>
          <div className={SearchStyles.SearchHint}>Please enter user name</div>
          <input
            className={SearchStyles.textInput}
            type="text"
            placeholder="User name"
            value={this.state.value}
            onChange={this.handleChange}
          />
          <button
            className={SearchStyles.searchButton}
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
