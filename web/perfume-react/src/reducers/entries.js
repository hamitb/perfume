import Service from '../components/Service';

async function fetchEntries() {
  try {
    const entriesJSON = await Service.getEntryList();
    const entriesList = entriesJSON.entries;
    return entriesList;
  } catch (err) {
    console.error(err);
    return [];
  }
}

function entries(state = [], action) {
  console.log("Entries reducer !");
  switch(action.type) {
    case 'GET_ENTRY_LIST':
      console.log("reducers->entries->get_entry_list:", entry_list);
      const entry_list = fetchEntries();
    default:
      return state;
  }
  return state;
}

export default entries;