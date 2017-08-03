import React from 'react';
//import YeomanImage from './YeomanImage';
import Header from './Header';
import SideBar from './SideBar';
import MainContent from './Content';
import '../styles/styles.css';

class AppComponent extends React.Component {

  render() {
    return (
      <div className="demo-layout mdl-layout mdl-js-layout mdl-layout--fixed-drawer mdl-layout--fixed-header">
        <Header/>
        <SideBar/>
        <MainContent {...this.props}/>
      </div>
    );
  }
}

AppComponent.defaultProps = {
};

export default AppComponent;
