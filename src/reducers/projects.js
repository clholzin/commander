import {
  LOAD_PROJECTS,
  PROJECTS_WERE_LOADED
} from '../actions/const'

const projects = [];
export default function  (state = projects, action) {
  switch(action.type){
    case LOAD_PROJECTS:{
      console.log('hit '+LOAD_PROJECTS);
      state =  action.projects;
    }
    break;
    case PROJECTS_WERE_LOADED:{
      state =  action.projects;
    }
    break;
  }
  return state;
}
