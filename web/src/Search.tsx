import * as mobx from 'mobx';
import { observer } from 'mobx-react';
import * as React from 'react';
import { Error } from './Error';
const SearchStyles = require('./Search.css');

type Props = {
  onSubmit(value: string): void;
  showError: string;
};

@observer
export class Search extends React.Component<Props> {
  @mobx.observable value: string = '';

  handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    this.value = event.target.value;
  };

  handleOnClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    this.props.onSubmit(this.value);
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
            value={this.value}
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
