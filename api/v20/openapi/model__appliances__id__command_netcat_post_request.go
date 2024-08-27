/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.3
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AppliancesIdCommandNetcatPostRequest struct for AppliancesIdCommandNetcatPostRequest
type AppliancesIdCommandNetcatPostRequest struct {
	// The destination to connect. Can be numerical IP address or a symbolic hostname.
	Destination string `json:"destination"`
	// The port to run command.
	Port int32 `json:"port"`
	// Select IP version explicitly.
	Version *int32 `json:"version,omitempty"`
	// Select protocol.
	Protocol *string `json:"protocol,omitempty"`
	// The number of seconds to run/wait before timing out.
	ProcessTimeout *int32 `json:"processTimeout,omitempty"`
}

// NewAppliancesIdCommandNetcatPostRequest instantiates a new AppliancesIdCommandNetcatPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliancesIdCommandNetcatPostRequest(destination string, port int32) *AppliancesIdCommandNetcatPostRequest {
	this := AppliancesIdCommandNetcatPostRequest{}
	this.Destination = destination
	this.Port = port
	var protocol string = "TCP"
	this.Protocol = &protocol
	var processTimeout int32 = 10
	this.ProcessTimeout = &processTimeout
	return &this
}

// NewAppliancesIdCommandNetcatPostRequestWithDefaults instantiates a new AppliancesIdCommandNetcatPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppliancesIdCommandNetcatPostRequestWithDefaults() *AppliancesIdCommandNetcatPostRequest {
	this := AppliancesIdCommandNetcatPostRequest{}
	var protocol string = "TCP"
	this.Protocol = &protocol
	var processTimeout int32 = 10
	this.ProcessTimeout = &processTimeout
	return &this
}

// GetDestination returns the Destination field value
func (o *AppliancesIdCommandNetcatPostRequest) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *AppliancesIdCommandNetcatPostRequest) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *AppliancesIdCommandNetcatPostRequest) SetDestination(v string) {
	o.Destination = v
}

// GetPort returns the Port field value
func (o *AppliancesIdCommandNetcatPostRequest) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *AppliancesIdCommandNetcatPostRequest) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *AppliancesIdCommandNetcatPostRequest) SetPort(v int32) {
	o.Port = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AppliancesIdCommandNetcatPostRequest) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdCommandNetcatPostRequest) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AppliancesIdCommandNetcatPostRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *AppliancesIdCommandNetcatPostRequest) SetVersion(v int32) {
	o.Version = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *AppliancesIdCommandNetcatPostRequest) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdCommandNetcatPostRequest) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *AppliancesIdCommandNetcatPostRequest) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *AppliancesIdCommandNetcatPostRequest) SetProtocol(v string) {
	o.Protocol = &v
}

// GetProcessTimeout returns the ProcessTimeout field value if set, zero value otherwise.
func (o *AppliancesIdCommandNetcatPostRequest) GetProcessTimeout() int32 {
	if o == nil || o.ProcessTimeout == nil {
		var ret int32
		return ret
	}
	return *o.ProcessTimeout
}

// GetProcessTimeoutOk returns a tuple with the ProcessTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdCommandNetcatPostRequest) GetProcessTimeoutOk() (*int32, bool) {
	if o == nil || o.ProcessTimeout == nil {
		return nil, false
	}
	return o.ProcessTimeout, true
}

// HasProcessTimeout returns a boolean if a field has been set.
func (o *AppliancesIdCommandNetcatPostRequest) HasProcessTimeout() bool {
	if o != nil && o.ProcessTimeout != nil {
		return true
	}

	return false
}

// SetProcessTimeout gets a reference to the given int32 and assigns it to the ProcessTimeout field.
func (o *AppliancesIdCommandNetcatPostRequest) SetProcessTimeout(v int32) {
	o.ProcessTimeout = &v
}

func (o AppliancesIdCommandNetcatPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.ProcessTimeout != nil {
		toSerialize["processTimeout"] = o.ProcessTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableAppliancesIdCommandNetcatPostRequest struct {
	value *AppliancesIdCommandNetcatPostRequest
	isSet bool
}

func (v NullableAppliancesIdCommandNetcatPostRequest) Get() *AppliancesIdCommandNetcatPostRequest {
	return v.value
}

func (v *NullableAppliancesIdCommandNetcatPostRequest) Set(val *AppliancesIdCommandNetcatPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliancesIdCommandNetcatPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliancesIdCommandNetcatPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliancesIdCommandNetcatPostRequest(val *AppliancesIdCommandNetcatPostRequest) *NullableAppliancesIdCommandNetcatPostRequest {
	return &NullableAppliancesIdCommandNetcatPostRequest{value: val, isSet: true}
}

func (v NullableAppliancesIdCommandNetcatPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliancesIdCommandNetcatPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
