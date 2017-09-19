import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as actionCreators from '../actions/actionCreators';

function mapStateToProps(state) {
  return {
    entries: state.entries,
  }
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators(actionCreators, dispatch);
}

class Main extends Component {
  render() {
    const { entries } = this.props;
    return (
      <div>
        {entries.map( (entry) => {
          return(
            <h2 key={entry.id}>{entry.title}</h2>
          )
        })}
      </div>
    );
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Main);
