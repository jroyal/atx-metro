import React, { Component } from "react";
import PropTypes from "prop-types";
import logo from "./logo.svg";
import "./App.css";
import "../node_modules/bootstrap/dist/css/bootstrap.min.css";
import { debounce } from "underscore";
import "./components/StopSearchBar";
import StopSearchBar from "./components/StopSearchBar";
import StopCard from "./components/StopCard";

class App extends Component {
  constructor(props) {
    super(props);
    console.log(process.env);
    this.state = {
      stop: "",
      possibleStops: []
    };
    this.onChange = this.onChange.bind(this);
    this.searchStops = this.searchStops.bind(this);
  }

  searchStops = debounce(name => {
    fetch("api/stops?name=" + name)
      .then(result => result.json())
      .then(data => {
        // TODO: Handle error
        if (Array.isArray(data)) {
          this.setState({ possibleStops: data });
        }
      });
  }, 500);

  onChange = e => {
    let name = e.target.value;
    this.setState({ stop: name });
    this.searchStops(name);
  };

  render() {
    const { possibleStops, stop } = this.state;
    const { apiKey } = this.props;
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </header>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>

        <div className="container-fluid">
          <div className="row justify-content-center timeConverter">
            <div className="col-6">
              <StopSearchBar onChange={this.onChange} value={stop} />
            </div>
          </div>
          <div className="row justify-content-center timeConverter">
            <div className="col-6">
              {possibleStops.map(stop => {
                return <StopCard apiKey={apiKey} stop={stop} />;
              })}
            </div>
          </div>
        </div>
      </div>
    );
  }
}

App.propTypes = {
  apiKey: PropTypes.string
};

export default App;
