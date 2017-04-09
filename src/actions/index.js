import { MESSAGE_RECEIVE, MESSAGE_SEND, MESSAGE_RESTORE, USERNAME_SET } from './types'
import ws from '../services/websocket'

export * from './auth'

export function sendMessage(message) {
    ws.send({ body: message });
    return {
        type: MESSAGE_SEND
    };
}

export function handleIncoming(message) {
    const action = {};
    switch (message.type) {
        case 0:
            action.type = MESSAGE_RECEIVE;
            action.payload = message;
            break;
        case 2:
            action.type = USERNAME_SET;
            action.payload = message.body;
            break;
    }
    if(message.error) {
        console.log(message.body);
        action.type = 0;
    }
    return action;
}

export function restoreMessages(stored) {
    return {
        type: MESSAGE_RESTORE,
        payload: JSON.parse(stored)
    }
}