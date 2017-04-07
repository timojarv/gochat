import ReconnectingWebSocket from 'reconnecting-websocket'

const ws = new ReconnectingWebSocket(`ws://${location.host}/ws`);

export default ws