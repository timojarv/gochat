require("./main.scss");
require("milligram");

import React from 'react';
import ReactDOM from 'react-dom';

import { Provider } from 'react-redux';
import { createStore, applyMiddleware, compose } from 'redux';
import thunk from 'redux-thunk';

import reducers from './reducers';

import { authenticate } from './actions';

import App from './components/app';

const composeEnchancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const store = createStore(reducers, composeEnchancers(
	applyMiddleware(thunk)
));

// Restore old messages
//const storedMessages = localStorage.getItem("messages");
//if(storedMessages) store.dispatch(restoreMessages(storedMessages));

// Restore username
//const username = localStorage.getItem("username");
//if(username) store.dispatch(setUsername(username));

let token = localStorage.getItem("webtoken");
if(token) store.dispatch(authenticate(token));

ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>,
    document.getElementById("root")
);