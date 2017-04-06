import React from 'react'

import MessageBox from './messagebox'
import SendField from './sendfield'

export default function App(props) {
    return (
        <div id="chat">
            <MessageBox />
            <SendField />
        </div>
    );
}