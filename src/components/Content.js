import React from 'react';
import '../styles/styles.css';
import List from './List';
import Cards from './Cards';

class Content extends React.Component {
  render() {
    return (
      <main className="mdl-layout__content mdl-color--grey-100">
        <div className="mdl-grid demo-content">
          <List {...this.props}/>
          <Cards/>
        </div>
      </main>
    );
  }
}

Content.defaultProps = {};

export default Content;
