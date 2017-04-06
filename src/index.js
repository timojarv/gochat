require("./main.scss")
require("siimple")

import React from 'react'
import ReactDOM from 'react-dom'

import { Provider } from 'react-redux';
import { createStore, applyMiddleware, compose } from 'redux';
import thunk from 'redux-thunk';

import reducers from './reducers';

import ws from './services/websocket'
import { handleMessage, restoreMessages } from './actions'

import App from './components/app'

const composeEnchancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const store = createStore(reducers, composeEnchancers(
	applyMiddleware(thunk)
));

ws.addEventListener("message", e => {
    const message = JSON.parse(e.data)
    store.dispatch(handleMessage(message))
})

// Restore old messages
const storedMessages = localStorage.getItem("messages");
if(storedMessages) store.dispatch(restoreMessages(storedMessages));

ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>,
    document.getElementById("root")
);