import React from 'react';
import ReactDOM from 'react-dom';
import Main from './Main';
import registerServiceWorker from './registerServiceWorker';

import { Router, Route } from 'react-router-dom';
import { Provider } from 'react-redux';
import store, { history } from './store';
ReactDOM.render(<Main />, document.getElementById('root'));
registerServiceWorker();
