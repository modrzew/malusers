import React, { Component } from 'react';

export class GetUser extends Component  {
  constructor(props) {
    super(props);
    this.state = {
      user: null,
      isLoading: false,
    }
  }

componentDidMount() {
  this.setState({isLoading: true});

  fetch(`https://jsonplaceholder.typicode.com/users/${this.props.match.params.user}`)
  .then((resp) => resp.json())
  .then((data) => {
    this.setState({user: data.user, isLoading: false})
  })
}

  render () {
    const {user, isLoading} = this.state;
    if(isLoading) {
      return <p>Loading...</p>
    }
    else if (user != null){
      return <UserPage user={this.state.user}/>
    }
    else {
      return <p>Loading...</p>
    }
  }
}