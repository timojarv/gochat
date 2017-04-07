import { USERNAME_SET } from '../actions/types'

export default function(state = "guest", action) {
    switch (action.type) {
        case USERNAME_SET:
            localStorage.setItem("username", action.payload);
            return action.payload;
    }
    return state;
}