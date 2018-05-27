import React, { Component } from "react";
import "../../node_modules/bootstrap/dist/css/bootstrap.min.css";

class StopSearchBar extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    const { onChange, value } = this.props;
    return (
      <div className="input-group mb-3">
        <input
          type="text"
          className="form-control text-center"
          onChange={onChange}
          value={value}
          type="text"
          id="time"
        />
        <div className="input-group-append">
          <button
            className="btn btn-outline-secondary"
            type="button"
            onClick={this.convertTimestamp}
          >
            Find Stop
          </button>
        </div>
      </div>
    );
  }
}

export default StopSearchBar;
