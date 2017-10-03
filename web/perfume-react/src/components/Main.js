import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as actionCreators from '../actions/actionCreators';
import { Button, Typography, TextField } from 'material-ui';

function mapStateToProps(state) {
  return {
    entries: state.entries,
  }
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators(actionCreators, dispatch);
}

class Main extends Component {
  handleSubmit = (e) => {
    e.preventDefault();
    console.log(this.refs);
  }

  render() {
    const { entries } = this.props;
    
    const exampleEntry = {
      "title": "JS: Promises",
      "link": "something.com",
      "labels": [ "label1", "label2" ],
    };
    
    return (
      <div>
        <Button color="primary" onClick={() => this.props.addEntryAsync(exampleEntry)}>
          ADD ENTRY
        </Button>
        { entries.map((entry) => {
          return (
            <div key={entry.id}>
              <div style={{ display: 'flex', flexDirection: 'column' }}>
                <Typography type="headline" onClick={() => this.props.getEntry(entry.id)}>{entry.title}</Typography>
                <form ref="testForm" onSubmit={this.handleSubmit}>
                  <input ref="title" placeholder="New Title"/>
                  <Button type="submit" color="accent">Update Title</Button>
                </form>
              </div>
              <ul>
                {entry.labels.map(label => <li key={label}>{label}</li>)}
              </ul>
            </div>
          );
        })}
      </div>
    );
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Main);
