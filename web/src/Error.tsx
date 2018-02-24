import * as React from 'react';
const styles = require('./Error.css');

export class Error extends React.Component {
  render() {
    return <div className={styles.Error}>Invalid user name</div>;
  }
}
