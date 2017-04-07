import React from 'react'
import { Emojione } from 'react-emoji-render'
import { connect } from 'react-redux'

class MessageContainer extends React.Component {
    componentDidUpdate() {
        document.body.scrollTop = document.body.scrollHeight;
    }

    componentDidMount() {
        document.body.scrollTop = document.body.scrollHeight;
    }

    renderMessage(message, i) {
        return (
            <div
                key={i}
                className={"message" + (message.username == this.props.username ? " own" : "")}
            >
                <strong className="sender">{message.username}: </strong>
                <Emojione svg>{message.message}</Emojione>
            </div>
        );
    }

    render() {
        return (
            <div className="messages">
                {this.props.messages.map(this.renderMessage.bind(this))}
            </div>
        );
    }
}

const mapStateToProps = state => ({ messages: state.messages, username: state.username })

export default connect(mapStateToProps)(MessageContainer)