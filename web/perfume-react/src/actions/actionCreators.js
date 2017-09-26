import Service from '../components/Service';

// addEntry
export function addEntry(entry) {
  return {
    type: "ADD_ENTRY",
    entry: entry,
  };
}

// setEntryList
export const receiveEntries = (entryList) => {
  return {
    type: "RECEIVE_ENTRIES",
    entryList: entryList,
  }
}

export const getAllEntries = () => dispatch => {
  Service.getEntryList().then(response => {
    const entryList = response.entries;

    dispatch(receiveEntries(entryList));
  })
}

export const getEntry = id => (dispatch) => {
  Service.getEntry(id).then(entry => {
    console.log("Entry found: ", entry);
    return entry;
  }).catch(err => {
    console.log("Entry not found: ", err);
    return {};
  })
}

export const addEntryAsync = entry => (dispatch) => {
  Service.createEntry(entry).then(response => {
    dispatch(addEntry(response));
  }).catch( err => {
    console.log("actions/actionCreateors.js: Error:", err);
  });
}