import React, { Component } from 'react';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {value: ''};
    
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleClick = this.handleClick.bind(this);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
  }

  handleSubmit(event) {
    alert('Please wait');
    event.preventDefault();
  }

      handleClick(event) {
      fetch('https://jsonplaceholder.typicode.com/users')
      .then((resp) => resp.json())
      .then((value) => {
        const object = JSON.stringify(value);
        alert(object);
      })
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
          <input type="submit" value="Submit" onClick={this.handleClick}/>
          </form>
        </div>
      </div>
    );
  }
}

export default App;
