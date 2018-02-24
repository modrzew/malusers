import * as React from 'react';
import * as ReactDOM from 'react-dom';
import './index.css';
import App from './App';
const registerServiceWorker = require('./registerServiceWorker.js').default;

ReactDOM.render(<App />, document.getElementById('root'));
registerServiceWorker();
