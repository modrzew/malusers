import React, { Component } from 'react';
import ErrorStyles from './Error.css';

export class Error extends Component {
  render() {
    return <div className={ErrorStyles.Error}>Invalid user name</div>;
  }
}
