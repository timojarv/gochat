import { AUTH, DEAUTH } from '../actions/types'

export default function(state = {authenticated: false}, action) {
    switch (action.type) {
        case AUTH: 
            return { authenticated: true, token: action.payload };
        case DEAUTH:
            return { authenticated: false, token: undefined };
    }
    return state;
}