import React, { Component } from 'react';
import './Loader.css';

export class Loader extends Component {
  render() {
    return (
      <div className="loader-wrapper">
        <div className="loader" />
      </div>
    );
  }
}
