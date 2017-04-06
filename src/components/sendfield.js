import React from 'react'
import { connect } from 'react-redux'

import * as actions from '../actions'

class SendField extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            message: ""
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSend = this.handleSend.bind(this);
    }

    handleChange(e) {
        const message = e.target.value;
        this.setState(state => ({ ...state, message }));
    }

    handleSend(e) {
        e.preventDefault()
        if(!this.state.message.length) return;
        this.props.sendMessage({
            message: this.state.message,
            username: "reactor"
        })
        this.setState(state => ({ ...state, message: "" }));
        this.input.focus();
    }

    render() {
        return (
            <form onSubmit={this.handleSend} className="sendField">
                <input
                    className="siimple-input"
                    type="text"
                    onChange={this.handleChange}
                    value={this.state.message}
                    ref={input => this.input = input}
                />
                <button className="siimple-btn siimple-btn--blue">Send</button>
            </form>
        );
    }
}

export default connect(null, actions)(SendField)