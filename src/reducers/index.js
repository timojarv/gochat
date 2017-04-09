import { combineReducers } from 'redux'

import messages from './messages'
import username from './username'
import auth from './auth'

export default combineReducers({
    messages,
    username,
    auth
});