import { createStore } from 'redux';
import { syncHistoryWithStore } from 'react-router-redux';
import createHistory from 'history/createBrowserHistory';
import rootReducer from './reducers/index';

const browserHistory = createHistory()

const defaultState = {
  entries: [],
}

const store = createStore(rootReducer, defaultState);


export const history = syncHistoryWithStore(browserHistory, store);

export default store;