import * as React from 'react';
const styles = require('./Loader.css');

export class Loader extends React.PureComponent {
  render() {
    return (
      <div className={styles.loaderWrapper}>
        <div className={styles.loader} />
      </div>
    );
  }
}
