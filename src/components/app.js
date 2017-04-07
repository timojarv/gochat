import React from 'react'

import MessageContainer from './messagecontainer'
import SendField from './sendfield'
import UsernameWidget from './username_widget'

export default function App(props) {
    return (
        <div id="chat">
            <header>
                <strong className="brand">GoChat</strong>
                <UsernameWidget />
            </header>
            <MessageContainer />
            <footer>
                <SendField />
            </footer>
        </div>
    );
}