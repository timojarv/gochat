import React from 'react'
import { connect } from 'react-redux'

import * as actions from '../actions'

class UsernameWidget extends React.Component {
    handleNameChange() {
        const newName = prompt("Enter new username:");
        this.props.setUsername(newName);
    }

    render() {
        return (
            <span className="usernameWidget">
                <em>{this.props.username}</em>
                <button className="button button-clear" onClick={this.handleNameChange.bind(this)}>Change</button>
            </span>
        );
    }
}

const mapStateToProps = state => ({ username: state.username });

export default connect(mapStateToProps, actions)(UsernameWidget);