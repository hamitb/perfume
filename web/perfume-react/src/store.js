import { createStore, applyMiddleware, compose } from 'redux';
import { syncHistoryWithStore } from 'react-router-redux';
import thunk from 'redux-thunk';
import createHistory from 'history/createBrowserHistory';
import rootReducer from './reducers/index';

const browserHistory = createHistory()

const defaultState = {
  entries: [],
}

const composeEnhancer = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const store = createStore(
  rootReducer,
  defaultState,
  composeEnhancer(
    applyMiddleware(thunk),
  ),
);



export const history = syncHistoryWithStore(browserHistory, store);

export default store;