import React, { Component } from 'react';
import Header from './components/Header/Header';
import './App.css';
import { connect, sendMsg } from './api';

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }
  
  send() {
    console.log("hello");
    sendMsg("hello");
  }


  render() {
    return (
      <div className="App">
        <Header />
          <button onClick={this.send}>Hit</button>  
      </div>
    );
  }
}

export default App;
