// addEntry
export function addEntry(entry) {
  console.log('---addEntry action---');
  return {
    type: "ADD_ENTRY",
    ...entry,
  };
}