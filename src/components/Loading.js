import React from 'react';
import '../styles/styles.css';

const style = {
  margin:'0 25% 0 25%'
};


/**
 * Usage
 *  const textObj = {text:'Loading projects'};
    return <Loading {...textObj} />
 */
class Loading extends React.Component {
  render() {
    return (<div style={style}>
        <h3>{this.props.text}</h3>
        <div id="p2" className="mdl-progress mdl-js-progress mdl-progress__indeterminate"></div>
      </div>)
  }
}

Loading.defaultProps = {};

export default Loading;
