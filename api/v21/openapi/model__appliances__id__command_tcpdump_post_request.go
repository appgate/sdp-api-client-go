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
)

// AppliancesIdCommandTcpdumpPostRequest struct for AppliancesIdCommandTcpdumpPostRequest
type AppliancesIdCommandTcpdumpPostRequest struct {
	// The expression to filter tcpdump.
	Expression *string `json:"expression,omitempty"`
	// The network interface to tcpdump on.
	Interface *string `json:"interface,omitempty"`
	// The number of seconds to run/wait before timing out.
	ProcessTimeout *int32 `json:"processTimeout,omitempty"`
}

// NewAppliancesIdCommandTcpdumpPostRequest instantiates a new AppliancesIdCommandTcpdumpPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliancesIdCommandTcpdumpPostRequest() *AppliancesIdCommandTcpdumpPostRequest {
	this := AppliancesIdCommandTcpdumpPostRequest{}
	var processTimeout int32 = 30
	this.ProcessTimeout = &processTimeout
	return &this
}

// NewAppliancesIdCommandTcpdumpPostRequestWithDefaults instantiates a new AppliancesIdCommandTcpdumpPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppliancesIdCommandTcpdumpPostRequestWithDefaults() *AppliancesIdCommandTcpdumpPostRequest {
	this := AppliancesIdCommandTcpdumpPostRequest{}
	var processTimeout int32 = 30
	this.ProcessTimeout = &processTimeout
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *AppliancesIdCommandTcpdumpPostRequest) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdCommandTcpdumpPostRequest) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *AppliancesIdCommandTcpdumpPostRequest) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *AppliancesIdCommandTcpdumpPostRequest) SetExpression(v string) {
	o.Expression = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *AppliancesIdCommandTcpdumpPostRequest) GetInterface() string {
	if o == nil || o.Interface == nil {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdCommandTcpdumpPostRequest) GetInterfaceOk() (*string, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *AppliancesIdCommandTcpdumpPostRequest) HasInterface() bool {
	if o != nil && o.Interface != nil {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *AppliancesIdCommandTcpdumpPostRequest) SetInterface(v string) {
	o.Interface = &v
}

// GetProcessTimeout returns the ProcessTimeout field value if set, zero value otherwise.
func (o *AppliancesIdCommandTcpdumpPostRequest) GetProcessTimeout() int32 {
	if o == nil || o.ProcessTimeout == nil {
		var ret int32
		return ret
	}
	return *o.ProcessTimeout
}

// GetProcessTimeoutOk returns a tuple with the ProcessTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdCommandTcpdumpPostRequest) GetProcessTimeoutOk() (*int32, bool) {
	if o == nil || o.ProcessTimeout == nil {
		return nil, false
	}
	return o.ProcessTimeout, true
}

// HasProcessTimeout returns a boolean if a field has been set.
func (o *AppliancesIdCommandTcpdumpPostRequest) HasProcessTimeout() bool {
	if o != nil && o.ProcessTimeout != nil {
		return true
	}

	return false
}

// SetProcessTimeout gets a reference to the given int32 and assigns it to the ProcessTimeout field.
func (o *AppliancesIdCommandTcpdumpPostRequest) SetProcessTimeout(v int32) {
	o.ProcessTimeout = &v
}

func (o AppliancesIdCommandTcpdumpPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.Interface != nil {
		toSerialize["interface"] = o.Interface
	}
	if o.ProcessTimeout != nil {
		toSerialize["processTimeout"] = o.ProcessTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableAppliancesIdCommandTcpdumpPostRequest struct {
	value *AppliancesIdCommandTcpdumpPostRequest
	isSet bool
}

func (v NullableAppliancesIdCommandTcpdumpPostRequest) Get() *AppliancesIdCommandTcpdumpPostRequest {
	return v.value
}

func (v *NullableAppliancesIdCommandTcpdumpPostRequest) Set(val *AppliancesIdCommandTcpdumpPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliancesIdCommandTcpdumpPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliancesIdCommandTcpdumpPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliancesIdCommandTcpdumpPostRequest(val *AppliancesIdCommandTcpdumpPostRequest) *NullableAppliancesIdCommandTcpdumpPostRequest {
	return &NullableAppliancesIdCommandTcpdumpPostRequest{value: val, isSet: true}
}

func (v NullableAppliancesIdCommandTcpdumpPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliancesIdCommandTcpdumpPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
