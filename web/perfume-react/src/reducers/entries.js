import { Service } from '../components/Service';

function entries(state = [], action) {
  console.log("Entries reducer !");
  switch(action.type) {
    case 'GET_ENTRY_LIST':
      console.log("reducers->entries->get_entry_list");
    default:
      return state;
  }
  return state;
}

export default entries;