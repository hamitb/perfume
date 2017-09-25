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

export const addEntryAsync = entry => (dispatch, getState) => {
  Service.createEntry(entry).then( response => {
    dispatch(addEntry(response));
  }).catch( err => {
    console.log("actions/actionCreateors.js: Error:", err);
  });
}