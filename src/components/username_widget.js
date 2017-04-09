import React from 'react'

class UsernameWidget extends React.Component {
    /*handleNameChange() {
        const newName = prompt("Enter new username:");
        this.props.setUsername(newName);
    }*/

    render() {
        return (
            <span className="usernameWidget">
                <em>{this.props.username}</em>
            </span>
        );
    }
}

export default UsernameWidget;