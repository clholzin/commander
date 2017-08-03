import { getJSON,postJSON,getWS} from './xhr'

const API = '/api/v1';
const WSAPI = '/api/ws';
const QUERY = '.json';//'$filter=Username%20eq%20%27NWARD%27%20and%20WorkbenchType%20eq%20%27BL%27&$format=json';
const protoCall = window.location.protocol;
const port = ':'+window.location.port;
const outSide = 'localhost';
export function fetchProjects(cb) {
  getJSON(`${API}/projects`, (error, res) => {
    cb(error, res)
  })
}
export function postCommandRequest(obj,cb) {
  postJSON(`${API}/process`,obj, (error, res) => {
    cb(error, res)
  })
}

export function getStatusWS(obj,cb) {
  getWS(`${outSide}${port}${WSAPI}/status`,obj, (error, res) => {
    cb(error, res)
  })
}
