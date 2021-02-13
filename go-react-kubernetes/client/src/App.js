import React from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';


class App extends React.Component{

  constructor(props){
    super(props);
    this.state = {
      ping: "Waiting"
    }
  }

  componentDidMount() {
    axios.get('/api/ping')
      .then((res, err) => {
        if(res.status === 200) {
          this.setState({
            ping: res.data.message
          })
        }
        else {
          console.log(err);
        }
      })
  }

  render(){
    
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <a
            className="App-link"
            target="_blank"
            rel="noopener noreferrer"
          >
            {this.state.ping}
          </a>
        </header>
      </div>
    );
  }
}

export default App;
