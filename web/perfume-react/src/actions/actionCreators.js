// addEntry
export function addEntry(entry) {
  console.log('---addEntry action---');
  return {
    type: "ADD_ENTRY",
    ...entry,
  };
}

// getEntryList
export function getEntryList() {
  console.log('---getEntryList action---');
  return {
    type: "GET_ENTRY_LIST",
  }
}