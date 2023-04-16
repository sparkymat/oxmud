import React from 'react';

const App = () => (
  <div className="uk-flex uk-flex-column" style={{ height: '100%' }}>
    <div className="uk-background-muted uk-flex-1 uk-flex uk-flex-column uk-padding-small">
      <div className="uk-background-secondary uk-flex-1 uk-padding-small">
        <p>output</p>
      </div>
      <div className="uk-margin-top">
        <input type="text" className="uk-input uk-text-muted" autofocus="" />
      </div>
    </div>
  </div>
);

export default App;
