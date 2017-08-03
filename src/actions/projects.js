import { fetchProjects,postCommandRequest,getStatusWS } from '../sources/projectsAPI'
import {LOAD_PROJECTS,
  PROJECTS_WERE_LOADED,
  COMMAND_POSTED,
  COMMAND_POST_DONE,
  STATUS_REQUESTED,
  STATUS_RECEIVED
} from './const';


export function loadProjects(dispatch) {
  const projects = [];
  dispatch({ type:
    LOAD_PROJECTS,
    projects
  });

  fetchProjects((error, projects) =>{
    dispatch({
      type: PROJECTS_WERE_LOADED,
      projects
    })
  })
}

export function postCommand(dispatch,command) {
  dispatch({ type:
  COMMAND_POSTED,
    command
  });

  postCommandRequest(command,(error, command) =>{
    console.log(command);
    dispatch({
      type: COMMAND_POST_DONE,
      command
    });
    getStatus(dispatch,JSON.parse(command))
  })
}

export function getStatus(dispatch,commandID) {
  const message = {message:[]};
  dispatch({ type:
  STATUS_REQUESTED,
    message
  });

  getStatusWS(commandID,(error, message) =>{
    dispatch({
      type: STATUS_RECEIVED,
      message
    })
  })
}
/*

 */
