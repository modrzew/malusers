import React, { Component } from 'react';
import './App.css';

export class Search extends Component {
  constructor(props) {
    super(props);
    this.state = {value: ''};
     }

  handleChange = (event) => {
    this.setState({value: event.target.value});
  }

  handleOnClick = (event) => {
    this.props.onSubmit(this.state.value);
  }
  
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">welcome to malusers</h1>
        </header>
        <div className="App-intro">
          <input type="text" placeholder="User name" value={this.state.value} onChange={this.handleChange}/>
          <input type="submit" value="Submit" onClick={this.handleOnClick}/>
        </div>
      </div>
    );
  }
}
