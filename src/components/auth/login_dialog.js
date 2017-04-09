import React from 'react'

class LoginDialog extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            username: "",
            password: ""
        };
    }

    handleLogin(e) {
        this.props.handleLogin(this.state.username, this.state.password);
        e.preventDefault();
    }

    render() {
        return (
            <form className="centered container" onSubmit={this.handleLogin.bind(this)}>
                <div className="spacer" />
                <h1>Log In</h1>
                <h5>If the account doesn't exist, it will be created automatically.</h5>
                <label>Username:</label>
                <input
                    className="narrow"
                    type="text"
                    value={this.state.username}
                    onChange={e => { let username = e.target.value; this.setState(s => ({ ...s, username })); }}
                />
                <label>Password:</label>
                <input
                    className="narrow"
                    type="password"
                    value={this.state.password}
                    onChange={e => { let password = e.target.value; this.setState(s => ({ ...s, password })); }}
                />

                <button>Log In or Register</button>
            </form>
        );
    }
}

export default LoginDialog;