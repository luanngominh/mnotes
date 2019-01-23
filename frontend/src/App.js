import React, { Component } from 'react';
import { Home } from "./page/home/Home"

class App extends Component {
  
  render() {
    if (this.loggedIn) {
      return "LoggedIN"
    } else {
      return <Home/>
    }
  }
}

export default App;
