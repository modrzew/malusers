import React, { Component } from 'react';
import LoaderStyles from './Loader.css';

export class Loader extends Component {
  render() {
    return (
      <div className={LoaderStyles.loaderWrapper}>
        <div className={LoaderStyles.loader} />
      </div>
    );
  }
}
