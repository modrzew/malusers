import React, { Component } from 'react';
import {Users} from './Users';

export class GetUser extends Component  {
  constructor(props) {
    super(props);
    this.state = {
      name: null,
      isLoading: false,
    }
  }

componentDidMount() {
  this.setState({isLoading: true});

  fetch(`https://jsonplaceholder.typicode.com/users/${this.props.userName}`)
  .then((resp) => resp.json())
  .then((data) => {
    this.setState({name: data, isLoading: false})
    const object = JSON.stringify(data);
    alert(object);
  })
}

  render () {
    const {name, isLoading} = this.state;
    if(isLoading) {
      return <p>Loading...</p>
    }
    else if (name != null){
      return <Users name={this.state.name}/>
    }
    else {
      return <p>Loading...</p>
    }
  }
}