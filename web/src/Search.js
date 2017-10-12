import React, { Component } from 'react';
import './App.css';
import { Link } from 'react-router-dom'

export class Search extends Component {
  constructor(props) {
    super(props);
    this.state = {value: ''};
    
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
  }

  handleSubmit(event) {
    alert('Please wait');
    event.preventDefault();
  }
  
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">welcome to malusers</h1>
        </header>
        <div className="App-intro">
          <form onSubmit={this.handleSubmit} >
          <input type="text" placeholder="User name" value={this.state.value} onChange={this.handleChange}/>
          <Link to={`user/${this.state.value}`}>
            <input type="submit" value="Submit"/>
          </Link>
          </form>
        </div>
      </div>
    );
  }
}
