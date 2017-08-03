import React from 'react';
import {loadProjects} from '../actions/projects';
import Loading from './Loading';
import ProjectComponent from './ProjectComponent';

class Project extends React.Component {
  constructor(props) {
    super(props);
    this.state = props
  }
  componentWillMount(){
    loadProjects(this.props.dispatch);
  }
  render() {
    const {projects,dispatch} = this.props;
    if(typeof projects === 'undefined' || projects.length === 0){
      const textObj = {text:'Loading projects'};
      return <Loading {...textObj} />
    }
    return (
      <div>
        <ul className="mdl-list">
          {projects.map((project,index) => (
                <li className="mdl-list__item" key={index}>
                  <ProjectComponent dispatch={dispatch} index={index} project={project}/>
                </li>
            ))
          }
        </ul>
      </div>
    );
  }
}

Project.defaultProps = {};

export default Project;


/*
 [{"folder":"C:\\nginx\\html\\BSP_Print_PDF","packagejson":{"name":"","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"","publish":"","start":"","serve":""}}},{"folder":"C:\\nginx\\html\\ChangeControlReport","packagejson":{"name":"ChangeControlReport","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"grunt build","publish":"","start":"","serve":"grunt serve"}}},{"folder":"C:\\nginx\\html\\ChangeControlReport2","packagejson":{"name":"ChangeControlReport_v2","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"grunt build","publish":"","start":"","serve":"grunt serve"}}},{"folder":"C:\\nginx\\html\\MPS-PDN-ULA","packagejson":{"name":"MSP ULA","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"sencha app build testing","publish":"sencha app build production","start":"","serve":""}}},{"folder":"C:\\nginx\\html\\MPS-PDN-ULA","packagejson":{"name":"MSP ULA","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"sencha app build testing","publish":"sencha app build production","start":"","serve":""}}},{"folder":"C:\\nginx\\html\\SIS-frontend","packagejson":{"name":"SIS-FrontEnd","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"","publish":"","start":"","serve":""}}},{"folder":"C:\\nginx\\html\\WBv2","packagejson":{"name":"Workbench v2","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"grunt build","publish":"","start":"","serve":"grunt serve"}}},{"folder":"C:\\nginx\\html\\WRA_CAA_Report","packagejson":{"name":"WRA_CAA_Report","scripts":{"eslint":"","stylelint":"","lint":"","test":"grunt test","build":"grunt build","publish":"","start":"grunt serve","serve":""}}},{"folder":"C:\\nginx\\html\\analytics","packagejson":{"name":"Analytics","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"node r.js -o single_js_file_build.js","publish":"","start":"","serve":""}}},{"folder":"C:\\nginx\\html\\home","packagejson":{"name":"commander","scripts":{"eslint":"","stylelint":"","lint":"eslint ./src","test":"cross-env NODE_ENV=test karma start","build":"","publish":"","start":"npm run serve:dev","serve":""}}},{"folder":"C:\\nginx\\html\\sspf_prototype","packagejson":{"name":"SSPF","scripts":{"eslint":"","stylelint":"","lint":"","test":"","build":"","publish":"","start":"","serve":""}}}]
* */
