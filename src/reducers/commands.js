import {
  COMMAND_POSTED,
  COMMAND_POST_DONE,
} from '../actions/const'

const command = {index:'',command:''};
export default function  (state = {command}, action) {
  switch(action.type){
    case COMMAND_POSTED:{
      console.log('hit '+COMMAND_POSTED);
      state.command =  action.command;
    }
    break;
    case COMMAND_POST_DONE:{
      state.command =  action.command;
    }
    break;
  }
  return state;
}
