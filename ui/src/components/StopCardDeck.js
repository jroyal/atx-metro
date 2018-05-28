import React, { Component } from "react";
import PropTypes from "prop-types";
import "../../node_modules/bootstrap/dist/css/bootstrap.min.css";
import StopCard from "./StopCard";

class StopCardDeck extends Component {
  render() {
    const { stops } = this.props;
    let cards = stops.map(stop => <StopCard stop={stop} />);
    var cardDeck = [];
    for (var i = 0; i < cards.length; i++) {
      if (i % 2 === 0 && i !== 0) {
        cardDeck.push(<div class="w-100 d-none d-sm-block d-md-none" />);
      }
      if (i % 3 === 0 && i !== 0) {
        cardDeck.push(
          <div class="w-100 d-none d-md-block d-lg-block d-xl-none" />
        );
      }
      if (i % 4 === 0 && i !== 0) {
        cardDeck.push(<div class="w-100 d-none d-xl-block" />);
      }
      cardDeck.push(cards[i]);
    }
    return <div className="card-deck">{cardDeck}</div>;
  }
}

StopCardDeck.propTypes = {
  stops: PropTypes.array
};

export default StopCardDeck;
