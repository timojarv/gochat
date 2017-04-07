import React from 'react'
import emojione from '../services/emojione'
import { connect } from 'react-redux'

class MessageBox extends React.Component {
    componentDidUpdate() {
        this.elem.scrollTop = this.elem.scrollHeight;
    }

    componentDidMount() {
        this.elem.scrollTop = this.elem.scrollHeight;
    }

    renderMessage(message, i) {
        return (
            <div
                key={i}
                className={"siimple-alert message" + (message.username == "reactor" ? " own" : "")}
            >
                <strong className="sender">{message.username}: </strong>
                <span
                    className="messageBody"
                    dangerouslySetInnerHTML={{__html:  emojione.toImage(message.message)}}
                />
            </div>
        );
    }

    render() {
        return (
            <div ref={mb => this.elem = mb} className="messageBox">
                {this.props.messages.map(this.renderMessage)}
            </div>
        );
    }
}

const mapStateToProps = state => ({ messages: state.messages })

export default connect(mapStateToProps)(MessageBox)