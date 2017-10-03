import React from 'react';
import ReactDOM from 'react-dom';
import Main from './components/Main';
import registerServiceWorker from './registerServiceWorker';

import { Router, Route } from 'react-router-dom';
import { Provider } from 'react-redux';
import store, { history } from './store';
import { getAllEntries } from './actions/actionCreators';

import { MuiThemeProvider, createMuiTheme } from 'material-ui/styles';
import { purple, green, red} from 'material-ui/colors';

const theme = createMuiTheme({
  palette: {
    primary: purple,
    secondary: {
      ...green,
      A400: '#00e677',
    },
    error: red,
  },
});

store.dispatch(getAllEntries());

const router = (
  <Provider store={store}>
    <MuiThemeProvider theme={theme}>
      <Router history={history}>
        <Route path="/" component={Main}/>
      </Router>
    </MuiThemeProvider>
  </Provider>
)

ReactDOM.render(router, document.getElementById('root'));
registerServiceWorker();
