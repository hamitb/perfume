import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as actionCreators from '../actions/actionCreators';

import Button from 'material-ui/Button';

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
    return (
      <div>
        <Button onClick={() => this.props.getEntryList()}>
          Hello World
        </Button>
      </div>
    );
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Main);
