import {
  STATUS_REQUESTED,
  STATUS_RECEIVED
} from '../actions/const'

const message = '';
export default function  (state = {message}, action) {
  switch(action.type){
    case STATUS_REQUESTED:{
      console.log('hit '+STATUS_REQUESTED);
      state.message =  action.message;
    }
    break;
    case STATUS_RECEIVED:{
      state.message =  action.message;
    }
    break;
  }
  return state;
}
