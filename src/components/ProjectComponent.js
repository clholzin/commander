import React from 'react';
import * as _ from 'underscore';
import {postCommand} from '../actions/projects';

class ProjectComponent extends React.Component {
  postCommandKey(key){
    let command = this.props.project.packagejson.scripts[key];
    let index = String(this.props.index);
    let dispatch = this.props.dispatch;
    postCommand(dispatch,{index,command})
  }
  render() {
    const project = this.props.project;
    const scripts = project.packagejson.scripts;
    const actionsBtnAvailable = _.chain(scripts)
      .keys()
      .filter((key)=>{
           return scripts[key] !== '';
        }).map(key=>{
          const obj = {key};
          obj['value'] = scripts[key];
          return obj;
        }).value();

    if(_.isEmpty(actionsBtnAvailable)){
      return (
        <div/>
      );
    }else{
      return (
        <div>
          <h4>{project.packagejson.name}</h4>
          {actionsBtnAvailable.map((obj,index) =>(
            <button onClick={(e)=>{
              this.postCommandKey(obj.key)
            }} key={index} className="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect" >{obj.key}</button>
          ))}
        </div>
      );
    }

  }
}

ProjectComponent.defaultProps = {};

export default ProjectComponent;


/*
 [{"folder":"C:\\nginx\\html\\BSP_Print_PDF","packagejson":{"name":"","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"","publish":"","start":"","serve":""}}},{"folder":"C:\\nginx\\html\\ChangeControlReport","packagejson":{"name":"ChangeControlReport","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"grunt build","publish":"","start":"","serve":"grunt serve"}}},{"folder":"C:\\nginx\\html\\ChangeControlReport2","packagejson":{"name":"ChangeControlReport_v2","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"grunt build","publish":"","start":"","serve":"grunt serve"}}},{"folder":"C:\\nginx\\html\\MPS-PDN-ULA","packagejson":{"name":"MSP ULA","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"sencha app build testing","publish":"sencha app build production","start":"","serve":""}}},{"folder":"C:\\nginx\\html\\MPS-PDN-ULA","packagejson":{"name":"MSP ULA","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"sencha app build testing","publish":"sencha app build production","start":"","serve":""}}},{"folder":"C:\\nginx\\html\\SIS-frontend","packagejson":{"name":"SIS-FrontEnd","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"","publish":"","start":"","serve":""}}},{"folder":"C:\\nginx\\html\\WBv2","packagejson":{"name":"Workbench v2","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"grunt build","publish":"","start":"","serve":"grunt serve"}}},{"folder":"C:\\nginx\\html\\WRA_CAA_Report","packagejson":{"name":"WRA_CAA_Report","scripts":{"eslint":"","stylelint":"","lint":"","test":"grunt test","build":"grunt build","publish":"","start":"grunt serve","serve":""}}},{"folder":"C:\\nginx\\html\\analytics","packagejson":{"name":"Analytics","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"node r.js -o single_js_file_build.js","publish":"","start":"","serve":""}}},{"folder":"C:\\nginx\\html\\home","packagejson":{"name":"commander","scripts":{"eslint":"","stylelint":"","lint":"eslint ./src","test":"cross-env NODE_ENV=test karma start","build":"","publish":"","start":"npm run serve:dev","serve":""}}},{"folder":"C:\\nginx\\html\\sspf_prototype","packagejson":{"name":"SSPF","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"","publish":"","start":"","serve":""}}}]
* */
