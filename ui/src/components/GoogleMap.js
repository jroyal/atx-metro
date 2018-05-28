import React, { Component } from "react";
import PropTypes from "prop-types";
import "../../node_modules/bootstrap/dist/css/bootstrap.min.css";

class GoogleMap extends Component {
  render() {
    const { latitude, longitude } = this.props;
    return (
      <iframe
        width="100%"
        height="450"
        frameBorder="0"
        style={{ border: 0 }}
        src={
          "https://www.google.com/maps/embed/v1/place?q=" +
          latitude +
          "," +
          longitude +
          "&key=" +
          process.env.REACT_APP_GOOGLE_MAPS_API_KEY
        }
      />
    );
  }
}
GoogleMap.propTypes = {
  latitude: PropTypes.string,
  longitude: PropTypes.string
};
export default GoogleMap;
