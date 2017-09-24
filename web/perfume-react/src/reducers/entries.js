function entries(state = [], action) {
  switch(action.type) {
    case 'SET_ENTRY_LIST':
      return action.entryList;
    case 'ADD_ENTRY':
      return state;
    default:
      return state;
  }
}

export default entries;