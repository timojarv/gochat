import ReconnectingWebSocket from 'reconnectingwebsocket';
import { store } from 'redux';
import { handleIncoming } from '../actions';

const conn = new ReconnectingWebSocket(`ws://${location.host}/ws`);

const ws = {
    queue: [],
    token: false,
    authenticated: false,
    send(data) {
        if(conn.readyState == 1) {
            conn.send(JSON.stringify(data));
        } else if(ws.authenticated) {
            ws.queue.push(data);
        }

        // Otherwise data is discarded
    },
    setToken(token) {
        if(ws.authenticated) return;
        ws.token = token;
        conn.refresh();
    },
    removeToken() {
        ws.token = false;
        conn.refresh();
    },
    registerHandler(func) {
        conn.addEventListener("message", e => func(JSON.parse(e.data)));
    } 
};

conn.addEventListener("open", function(e) {
    if(!ws.token) {
        ws.authenticated = false;
        return;
    }
    // Auth first
    conn.send(JSON.stringify({ body: ws.token }));
    ws.authenticated = true;
    ws.queue.map(
        data => conn.send(JSON.stringify(data))
    );
});

export default ws;