import React, { Component } from 'react';
import styles from './Loader.css';

export class Loader extends Component {
  render() {
    return (
      <div className={styles.loaderWrapper}>
        <div className={styles.loader} />
      </div>
    );
  }
}
