import { MESSAGE_RECEIVE, MESSAGE_SEND, MESSAGE_RESTORE } from './types'
import ws from '../services/websocket'

export function sendMessage(message) {
    ws.send(JSON.stringify(message))
    return {
        type: MESSAGE_SEND
    };
}

export function handleMessage(message) {
    return {
        type: MESSAGE_RECEIVE,
        payload: message
    };
}

export function restoreMessages(stored) {
    return {
        type: MESSAGE_RESTORE,
        payload: JSON.parse(stored)
    }
}