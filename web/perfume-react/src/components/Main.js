import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as actionCreators from '../actions/actionCreators';
import { Button, Typography } from 'material-ui';

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
    
    const exampleEntry = {
      "title": "JS: Promises",
      "link": "something.com",
      "labels": [ "label1", "label2" ],
    };
    
    return (
      <div>
        <Button onClick={() => this.props.addEntryAsync(exampleEntry)}>
          ADD ENTRY
        </Button>
        { entries.map((entry) => {
          return (
            <Typography key={entry.id} type="headline">{entry.title}</Typography>
          );
        })}
      </div>
    );
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Main);
