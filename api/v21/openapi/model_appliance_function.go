/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v21+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 21.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ApplianceFunction the model 'ApplianceFunction'
type ApplianceFunction string

// List of ApplianceFunction
const (
	CONTROLLER         ApplianceFunction = "Controller"
	GATEWAY            ApplianceFunction = "Gateway"
	LOG_SERVER         ApplianceFunction = "LogServer"
	LOG_FORWARDER      ApplianceFunction = "LogForwarder"
	CONNECTOR          ApplianceFunction = "Connector"
	PORTAL             ApplianceFunction = "Portal"
	METRICS_AGGREGATOR ApplianceFunction = "MetricsAggregator"
)

// All allowed values of ApplianceFunction enum
var AllowedApplianceFunctionEnumValues = []ApplianceFunction{
	"Controller",
	"Gateway",
	"LogServer",
	"LogForwarder",
	"Connector",
	"Portal",
	"MetricsAggregator",
}

func (v *ApplianceFunction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApplianceFunction(value)
	for _, existing := range AllowedApplianceFunctionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApplianceFunction", value)
}

// NewApplianceFunctionFromValue returns a pointer to a valid ApplianceFunction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApplianceFunctionFromValue(v string) (*ApplianceFunction, error) {
	ev := ApplianceFunction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApplianceFunction: valid values are %v", v, AllowedApplianceFunctionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApplianceFunction) IsValid() bool {
	for _, existing := range AllowedApplianceFunctionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplianceFunction value
func (v ApplianceFunction) Ptr() *ApplianceFunction {
	return &v
}

type NullableApplianceFunction struct {
	value *ApplianceFunction
	isSet bool
}

func (v NullableApplianceFunction) Get() *ApplianceFunction {
	return v.value
}

func (v *NullableApplianceFunction) Set(val *ApplianceFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceFunction(val *ApplianceFunction) *NullableApplianceFunction {
	return &NullableApplianceFunction{value: val, isSet: true}
}

func (v NullableApplianceFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
