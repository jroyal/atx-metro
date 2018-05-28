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
    const { stop } = this.props;
    const { showMap } = this.state;
    return (
      <div className="card mb-4 ml-4 mr-4">
        <div className="card-body">
          <h5 className="card-title">{stop.Name}</h5>
          <h6 className="card-subtitle mb-2 text-muted">Stop ID: {stop.ID}</h6>
          <p className="card-text">{stop.Desc}</p>
          <div
            className="btn-toolbar d-inline-block center-block"
            role="toolbar"
          >
            <a
              href={stop.URL}
              target="_blank"
              className="btn btn-outline-info btn-sm mr-2"
            >
              Stop Page
            </a>

            <button
              onClick={this.onClick}
              className="btn btn-outline-primary btn-sm"
            >
              Show on map
            </button>
          </div>
          {showMap && (
            <GoogleMap latitude={stop.Latitude} longitude={stop.Longitude} />
          )}
        </div>
      </div>
    );
  }
}

StopCard.propTypes = {
  stop: PropTypes.object
};

export default StopCard;
