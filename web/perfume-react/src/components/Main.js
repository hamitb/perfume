import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as actionCreators from '../actions/actionCreators';
import Service from '../components/Service';
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
  
  componentWillMount() {
    this.fetchEntries();
  }

  async fetchEntries() {
    try {
      const { entries } = await Service.getEntryList()
      this.props.setEntryList(entries);
    } catch (err) {
      console.log("components/main.js: Error on fetching entries");
      this.props.setEntryList([]);
    }
  }

  render() {
    const { entries } = this.props;
    
    return (
      <div>
        <Button onClick={() => this.props.addEntry()}>
          Hello World
        </Button>
        { entries.map((entry) => {
          return (
            <h2 key={entry.id}>{entry.title}</h2>
          );
        })}
      </div>
    );
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Main);
