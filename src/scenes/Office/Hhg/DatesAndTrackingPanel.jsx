import { get } from 'lodash';
import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { reduxForm, getFormValues, isValid } from 'redux-form';
import editablePanel from './editablePanel';

// import { updateOrders } from './ducks';

import { SwaggerField } from 'shared/JsonSchemaForm/JsonSchemaField';
import { PanelSwaggerField } from 'shared/EditablePanel';

const DatesAndTrackingDisplay = props => {
  const fieldProps = {
    schema: props.shipmentSchema,
    values: props.officeHHG,
  };

  return (
    <div className="editable-panel-column">
      <PanelSwaggerField
        title="Pickup Date"
        fieldName="pickup_date"
        {...fieldProps}
      />
    </div>
  );
};

const DatesAndTrackingEdit = props => {
  const { shipmentSchema } = props;
  return (
    <div className="editable-panel-column">
      <SwaggerField
        title="Pickup Date"
        fieldName="pickup_date"
        swagger={shipmentSchema}
        required
      />
    </div>
  );
};

const formName = 'office_move_info_accounting';

let DatesAndTrackingPanel = editablePanel(
  DatesAndTrackingDisplay,
  DatesAndTrackingEdit,
);
DatesAndTrackingPanel = reduxForm({ form: formName })(DatesAndTrackingPanel);

function mapStateToProps(state) {
  let orders = get(state, 'office.officeOrders', {});

  return {
    // reduxForm
    initialValues: state.office.officeOrders,

    // Wrapper
    shipmentSchema: get(state, 'swagger.spec.definitions.Shipment', {}),
    hasError:
      state.office.ordersHaveLoadError || state.office.ordersHaveUpdateError,
    errorMessage: state.office.error,

    shipment: shipment,
    isUpdating: state.office.shipmentIsUpdating,

    // editablePanel
    formIsValid: isValid(formName)(state),
    getUpdateArgs: function() {
      let values = getFormValues(formName)(state);
      return [shipment.id, values];
    },
  };
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators(
    {
      update: updateShipment,
    },
    dispatch,
  );
}

export default connect(mapStateToProps, mapDispatchToProps)(
  DatesAndTrackingPanel,
);
