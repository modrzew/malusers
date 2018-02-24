import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { App } from './App';
import './index.css';
const registerServiceWorker = require('./registerServiceWorker.js').default;

ReactDOM.render(<App />, document.getElementById('root'));
registerServiceWorker();
