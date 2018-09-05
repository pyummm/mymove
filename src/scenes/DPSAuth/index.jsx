import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import PropTypes from 'prop-types';
import { reduxForm } from 'redux-form';
import { SwaggerField } from 'shared/JsonSchemaForm/JsonSchemaField';
import { setDPSAuthCookie } from './ducks';

const schema = {
  properties: {
    cookie_name: {
      type: 'string',
      title: 'Cookie Name',
    },
    redirect_url: {
      type: 'string',
      title: 'Redirect URL',
    },
  },
};

export class DPSAuth extends Component {
  sendRequest = values => {
    console.log(values);
    this.props.setDPSAuthCookie(values.cookie_name, values.redirect_url);
  };

  render() {
    return (
      <div className="usa-grid">
        <h1 className="sm-heading">Redirect to DPS</h1>
        <form onSubmit={this.props.handleSubmit(this.sendRequest)}>
          <SwaggerField
            fieldName="cookie_name"
            swagger={this.props.schema}
            required
          />
          <SwaggerField
            fieldName="redirect_url"
            swagger={this.props.schema}
            required
          />
          <button type="submit">Submit</button>
        </form>
      </div>
    );
  }
}
DPSAuth.propTypes = {
  setDPSAuthCookie: PropTypes.func.isRequired,
  schema: PropTypes.object.isRequired,
};

function mapStateToProps(state) {
  return {
    schema,
  };
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators({ setDPSAuthCookie }, dispatch);
}

export default connect(mapStateToProps, mapDispatchToProps)(
  reduxForm({ form: 'dpsAuth' })(DPSAuth),
);
