import React, { Component } from "react";
import PropTypes from "prop-types";
import "../../node_modules/bootstrap/dist/css/bootstrap.min.css";
import GoogleMap from "./GoogleMap";

class StopCard extends Component {
  constructor(props) {
    super(props);
    this.state = {
      showMap: false
    };
    this.onClick = this.onClick.bind(this);
  }

  onClick = () => {
    this.setState({ showMap: !this.state.showMap });
  };

  render() {
    const { apiKey, stop } = this.props;
    const { showMap } = this.state;
    console.log(stop);
    return (
      <div className="card">
        <div className="card-body">
          <h5 className="card-title">{stop.Name}</h5>
          <h6 className="card-subtitle mb-2 text-muted">Stop ID: {stop.ID}</h6>
          <p className="card-text">{stop.Desc}</p>
          <button onClick={this.onClick} className="btn btn-primary">
            Show on map
          </button>
          {showMap && (
            <GoogleMap
              apiKey={apiKey}
              latitude={stop.Latitude}
              longitude={stop.Longitude}
            />
          )}
        </div>
      </div>
    );
  }
}

StopCard.propTypes = {
  stop: PropTypes.object,
  apiKey: PropTypes.string
};

export default StopCard;
