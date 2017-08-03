import React, {
  Component
} from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Main from '../components/App';

class App extends Component {
  render() {
    //const { actions,state } = this.props;
    return <Main {...this.props} />;
  }
}

App.propTypes = {
  actions: PropTypes.shape({})
};

function mapStateToProps(state) { // eslint-disable-line no-unused-vars
  const {projects} = state;
  const props = {projects};
  return props;
}
function mapDispatchToProps(dispatch) {
  const actions = {};
  const actionMap = { actions: bindActionCreators(actions, dispatch),dispatch};
  return actionMap;
}
export default connect(mapStateToProps, mapDispatchToProps)(App);
