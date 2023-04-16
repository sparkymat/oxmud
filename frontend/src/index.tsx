import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';

import { store } from './store';
import App from './components/App';

const element = document.getElementById('oxmud-app');

if (element) {
  ReactDOM.render(
    <Provider store={store}>
      <App />
    </Provider>,
    element,
  );
}
