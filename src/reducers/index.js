import { combineReducers } from 'redux'

import messages from './messages'
import username from './username'

export default combineReducers({
    messages,
    username
});