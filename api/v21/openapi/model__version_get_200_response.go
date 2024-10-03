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

// VersionGet200Response struct for VersionGet200Response
type VersionGet200Response struct {
	// The current peer version.
	PeerVersion *int32 `json:"peerVersion,omitempty"`
	// The lowest accepted peer version.
	LowestPeerVersionSupported *int32 `json:"lowestPeerVersionSupported,omitempty"`
}

// NewVersionGet200Response instantiates a new VersionGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionGet200Response() *VersionGet200Response {
	this := VersionGet200Response{}
	return &this
}

// NewVersionGet200ResponseWithDefaults instantiates a new VersionGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionGet200ResponseWithDefaults() *VersionGet200Response {
	this := VersionGet200Response{}
	return &this
}

// GetPeerVersion returns the PeerVersion field value if set, zero value otherwise.
func (o *VersionGet200Response) GetPeerVersion() int32 {
	if o == nil || o.PeerVersion == nil {
		var ret int32
		return ret
	}
	return *o.PeerVersion
}

// GetPeerVersionOk returns a tuple with the PeerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionGet200Response) GetPeerVersionOk() (*int32, bool) {
	if o == nil || o.PeerVersion == nil {
		return nil, false
	}
	return o.PeerVersion, true
}

// HasPeerVersion returns a boolean if a field has been set.
func (o *VersionGet200Response) HasPeerVersion() bool {
	if o != nil && o.PeerVersion != nil {
		return true
	}

	return false
}

// SetPeerVersion gets a reference to the given int32 and assigns it to the PeerVersion field.
func (o *VersionGet200Response) SetPeerVersion(v int32) {
	o.PeerVersion = &v
}

// GetLowestPeerVersionSupported returns the LowestPeerVersionSupported field value if set, zero value otherwise.
func (o *VersionGet200Response) GetLowestPeerVersionSupported() int32 {
	if o == nil || o.LowestPeerVersionSupported == nil {
		var ret int32
		return ret
	}
	return *o.LowestPeerVersionSupported
}

// GetLowestPeerVersionSupportedOk returns a tuple with the LowestPeerVersionSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionGet200Response) GetLowestPeerVersionSupportedOk() (*int32, bool) {
	if o == nil || o.LowestPeerVersionSupported == nil {
		return nil, false
	}
	return o.LowestPeerVersionSupported, true
}

// HasLowestPeerVersionSupported returns a boolean if a field has been set.
func (o *VersionGet200Response) HasLowestPeerVersionSupported() bool {
	if o != nil && o.LowestPeerVersionSupported != nil {
		return true
	}

	return false
}

// SetLowestPeerVersionSupported gets a reference to the given int32 and assigns it to the LowestPeerVersionSupported field.
func (o *VersionGet200Response) SetLowestPeerVersionSupported(v int32) {
	o.LowestPeerVersionSupported = &v
}

func (o VersionGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PeerVersion != nil {
		toSerialize["peerVersion"] = o.PeerVersion
	}
	if o.LowestPeerVersionSupported != nil {
		toSerialize["lowestPeerVersionSupported"] = o.LowestPeerVersionSupported
	}
	return json.Marshal(toSerialize)
}

type NullableVersionGet200Response struct {
	value *VersionGet200Response
	isSet bool
}

func (v NullableVersionGet200Response) Get() *VersionGet200Response {
	return v.value
}

func (v *NullableVersionGet200Response) Set(val *VersionGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionGet200Response(val *VersionGet200Response) *NullableVersionGet200Response {
	return &NullableVersionGet200Response{value: val, isSet: true}
}

func (v NullableVersionGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}