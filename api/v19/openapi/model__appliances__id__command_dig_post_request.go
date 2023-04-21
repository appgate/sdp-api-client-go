/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v19+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 19.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AppliancesIdCommandDigPostRequest struct for AppliancesIdCommandDigPostRequest
type AppliancesIdCommandDigPostRequest struct {
	// The host to query.
	Host string `json:"host"`
	// The DNS server to use for querying.
	Server *string `json:"server,omitempty"`
	// Select DNS request type.
	Type *string `json:"type,omitempty"`
	// Select protocol.
	Protocol *string `json:"protocol,omitempty"`
}

// NewAppliancesIdCommandDigPostRequest instantiates a new AppliancesIdCommandDigPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliancesIdCommandDigPostRequest(host string) *AppliancesIdCommandDigPostRequest {
	this := AppliancesIdCommandDigPostRequest{}
	this.Host = host
	var type_ string = "A"
	this.Type = &type_
	var protocol string = "UDP"
	this.Protocol = &protocol
	return &this
}

// NewAppliancesIdCommandDigPostRequestWithDefaults instantiates a new AppliancesIdCommandDigPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppliancesIdCommandDigPostRequestWithDefaults() *AppliancesIdCommandDigPostRequest {
	this := AppliancesIdCommandDigPostRequest{}
	var type_ string = "A"
	this.Type = &type_
	var protocol string = "UDP"
	this.Protocol = &protocol
	return &this
}

// GetHost returns the Host field value
func (o *AppliancesIdCommandDigPostRequest) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *AppliancesIdCommandDigPostRequest) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *AppliancesIdCommandDigPostRequest) SetHost(v string) {
	o.Host = v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *AppliancesIdCommandDigPostRequest) GetServer() string {
	if o == nil || o.Server == nil {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdCommandDigPostRequest) GetServerOk() (*string, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *AppliancesIdCommandDigPostRequest) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *AppliancesIdCommandDigPostRequest) SetServer(v string) {
	o.Server = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AppliancesIdCommandDigPostRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdCommandDigPostRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AppliancesIdCommandDigPostRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AppliancesIdCommandDigPostRequest) SetType(v string) {
	o.Type = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *AppliancesIdCommandDigPostRequest) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdCommandDigPostRequest) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *AppliancesIdCommandDigPostRequest) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *AppliancesIdCommandDigPostRequest) SetProtocol(v string) {
	o.Protocol = &v
}

func (o AppliancesIdCommandDigPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if o.Server != nil {
		toSerialize["server"] = o.Server
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableAppliancesIdCommandDigPostRequest struct {
	value *AppliancesIdCommandDigPostRequest
	isSet bool
}

func (v NullableAppliancesIdCommandDigPostRequest) Get() *AppliancesIdCommandDigPostRequest {
	return v.value
}

func (v *NullableAppliancesIdCommandDigPostRequest) Set(val *AppliancesIdCommandDigPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliancesIdCommandDigPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliancesIdCommandDigPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliancesIdCommandDigPostRequest(val *AppliancesIdCommandDigPostRequest) *NullableAppliancesIdCommandDigPostRequest {
	return &NullableAppliancesIdCommandDigPostRequest{value: val, isSet: true}
}

func (v NullableAppliancesIdCommandDigPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliancesIdCommandDigPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
