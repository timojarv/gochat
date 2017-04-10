import React from 'react'

import MessageContainer from './message_container'
import SendField from './send_field'
import UsernameWidget from './username_widget'

export default function Chat(props) {
    return (
        <div id="chat">
            <header>
                <strong className="brand">RatChat</strong>
                <UsernameWidget username={props.username} />
            </header>
            <MessageContainer />
            <footer>
                <SendField />
            </footer>
        </div>
    );
}