import axios from 'axios';
import ws from '../services/websocket';
import { AUTH, DEAUTH } from './types';


// Validate token and authenticate
export function authenticate(token) {
    return dispatch => {
        axios.post("/validate", { token })
        .then(response => {
            if(!response.data.error) {
                ws.setToken(token);
                dispatch({
                    type: AUTH,
                    payload: token
                });
            }
        });
    };
}

export function login(username, password) {
    return dispatch => {
        axios.post("/login", { username, password })
            .then(response => {
                if(!response.data.error) {
                    ws.setToken(response.data.data);
                    dispatch({
                        type: AUTH,
                        payload: response.data.data
                    });
                }
            });
    };
}