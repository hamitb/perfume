// addEntry
export function addEntry(entry) {
  console.log('---addEntry action---');
  return {
    type: "ADD_ENTRY",
    entry: entry,
  };
}

// setEntryList
export function setEntryList(entryList) {
  console.log('---setEntryList action---');
  return {
    type: "SET_ENTRY_LIST",
    entryList: entryList,
  }
}