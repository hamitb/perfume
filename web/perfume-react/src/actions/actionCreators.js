import Service from '../components/Service';

// addEntry
export function addEntry(entry) {
  console.log('---addEntry action---');
  return {
    type: "ADD_ENTRY",
    entry: entry,
  };
}

// setEntryList
export const receiveEntries = (entryList) => {
  console.log('---receiveEntries action---');
  return {
    type: "RECEIVE_ENTRIES",
    entryList: entryList,
  }
}

export const getAllEntries = () => dispatch => {
  console.log("I'm here");
  Service.getEntryList().then(response => {
    const entryList = response.entries;

    dispatch(receiveEntries(entryList));
  })
}