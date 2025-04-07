/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IdentityProvidersTestPost200Response struct for IdentityProvidersTestPost200Response
type IdentityProvidersTestPost200Response struct {
	// Whether the connection succeeded or not.
	Success *bool `json:"success,omitempty"`
	// The error text if the connection fails.
	Error *bool `json:"error,omitempty"`
}

// NewIdentityProvidersTestPost200Response instantiates a new IdentityProvidersTestPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProvidersTestPost200Response() *IdentityProvidersTestPost200Response {
	this := IdentityProvidersTestPost200Response{}
	return &this
}

// NewIdentityProvidersTestPost200ResponseWithDefaults instantiates a new IdentityProvidersTestPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProvidersTestPost200ResponseWithDefaults() *IdentityProvidersTestPost200Response {
	this := IdentityProvidersTestPost200Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *IdentityProvidersTestPost200Response) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersTestPost200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *IdentityProvidersTestPost200Response) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *IdentityProvidersTestPost200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *IdentityProvidersTestPost200Response) GetError() bool {
	if o == nil || o.Error == nil {
		var ret bool
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersTestPost200Response) GetErrorOk() (*bool, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *IdentityProvidersTestPost200Response) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given bool and assigns it to the Error field.
func (o *IdentityProvidersTestPost200Response) SetError(v bool) {
	o.Error = &v
}

func (o IdentityProvidersTestPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityProvidersTestPost200Response struct {
	value *IdentityProvidersTestPost200Response
	isSet bool
}

func (v NullableIdentityProvidersTestPost200Response) Get() *IdentityProvidersTestPost200Response {
	return v.value
}

func (v *NullableIdentityProvidersTestPost200Response) Set(val *IdentityProvidersTestPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProvidersTestPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProvidersTestPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProvidersTestPost200Response(val *IdentityProvidersTestPost200Response) *NullableIdentityProvidersTestPost200Response {
	return &NullableIdentityProvidersTestPost200Response{value: val, isSet: true}
}

func (v NullableIdentityProvidersTestPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProvidersTestPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
