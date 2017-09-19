import React from 'react';
import ReactDOM from 'react-dom';
import Main from './components/Main';
import registerServiceWorker from './registerServiceWorker';

import { Router, Route } from 'react-router-dom';
import { Provider } from 'react-redux';
import store, { history } from './store';

const router = (
  <Provider store={store}>
    <Router history={history}>
      <Route path="/" component={Main}/>
    </Router>
  </Provider>
)
ReactDOM.render(router, document.getElementById('root'));
registerServiceWorker();
