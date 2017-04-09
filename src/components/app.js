import React from 'react';
import { connect } from 'react-redux';

import * as actions from '../actions';

import ws from '../services/websocket';

import Chat from './chat';
import LoginDialog from './auth/login_dialog';

class App extends React.Component {
    componentWillMount() {
        ws.registerHandler(this.props.handleIncoming);
    }

    handleLogin(username, password) {
        this.props.login(username, password);
    }

    render() {
        return this.props.authenticated 
            ? <Chat username={this.props.username} token={this.props.token} /> 
            : <LoginDialog handleLogin={this.handleLogin.bind(this)} />
        ;
    }
}

const mapStateToProps = state => ({
    authenticated: state.auth.authenticated,
    token: state.auth.token,
    username: state.username
});

export default connect(mapStateToProps, actions)(App);