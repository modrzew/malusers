import * as React from 'react';
import { Error } from './Error';
const SearchStyles = require('./Search.css');

type Props = {
  onSubmit(value: string): void;
  showError: string;
};
type State = {
  value: string;
};

export class Search extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = { value: '' };
  }

  handleChange = (event: any) => {
    this.setState({ value: event.target.value });
  };

  handleOnClick = (event: any) => {
    this.props.onSubmit(this.state.value);
  };

  render() {
    return (
      <div className={SearchStyles.Search}>
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
