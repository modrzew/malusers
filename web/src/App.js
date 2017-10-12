import React, { Component } from 'react';
import './App.css';
import {Switch, Route} from 'react-router-dom';
import {Search} from './Search';
import {GetUser} from './GetUser';


class App extends Component {
  render () {
    return (
        <Switch>
            <Route exact path='/' component={Search}/>
            <Route path='/user/:user' component={GetUser}/>
        </Switch>
    )
  }
}

export default App;