function entries(state = [], action) {
  switch(action.type) {
    case 'RECEIVE_ENTRIES':
      return action.entryList;
    case 'ADD_ENTRY':
      return [
        ...state,
        action.entry,
      ];
    default:
      return state;
  }
}

export default entries;