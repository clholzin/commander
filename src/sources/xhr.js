localStorage.token = localStorage.token || 'public';//(Date.now()*Math.random())

function setToken(req) {
  req.setRequestHeader('Content-Type', 'application/json');
  req.setRequestHeader('Access-Control-Allow-Origin', 'http://localhost');
}

export function getJSON(url, callback) {
  const req = new XMLHttpRequest();
  req.onload = function () {
    if (req.status > 300) {
      callback(new Error('not found'))
    } else {
      let json = JSON.parse(req.response);
      callback(null, json)
    }
  };
  req.open('GET', url);
  setToken(req);
  req.send();
}

export function postJSON(url, obj, callback) {
  const req = new XMLHttpRequest();
  req.onload = function () {
      if (req.status === 500) {
        callback(new Error(req.responseText))
      } else {
        callback(null, req.responseText)
      }
  };
  req.open('POST', url);
  req.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
  setToken(req);
  req.send(JSON.stringify(obj))
}

export function getWS(url, obj, callback){
  let conn = new WebSocket("ws://" + url);

  conn.onclose = function (evt) {
    callback(null,'Closed');
  };
  conn.onmessage = function (evt) {
    let messages = {message:evt.data.split('\n')};
    callback(null,messages);
  };
  // Handle any errors that occur.
  conn.onerror = function(error) {
    callback(error);
    console.log('WebSocket Error: ' + JSON.stringify(error));
  };

  conn.onopen = () => conn.send(JSON.stringify(obj));
}

export function deleteJSON(url, callback) {
  const req = new XMLHttpRequest();
  req.onload = function () {
    setTimeout(() => {
      if (req.status === 500) {
        callback(new Error(req.responseText))
      } else {
        callback(null, req.responseText)
      }
    }, Math.random() * 5000)
  };
  req.open('DELETE', url);
  setToken(req);
  req.send();
}
