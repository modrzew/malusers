import React, { Component } from 'react';
import styles from './Error.css';

export class Error extends Component {
  render() {
    return <div className={styles.Error}>Invalid user name</div>;
  }
}
