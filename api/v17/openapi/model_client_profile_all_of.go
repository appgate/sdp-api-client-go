/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.5
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ClientProfileAllOf Represents a Client Profile.
type ClientProfileAllOf struct {
	// A name to identify the client profile. It will appear on the client UI.
	Name string `json:"name"`
	// SPA key name to be used in the profile. Same key names in different profiles will have the same SPA key. SPA key is used by the client to connect to the controllers.
	SpaKeyName string `json:"spaKeyName"`
	// Name of the Identity Provider to be used to authenticate.
	IdentityProviderName string `json:"identityProviderName"`
	// Overrides the Profile Hostname in global settings for this specific profile. Generated URLs will use this hostname instead.
	Hostname *string `json:"hostname,omitempty"`
}

// NewClientProfileAllOf instantiates a new ClientProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientProfileAllOf(name string, spaKeyName string, identityProviderName string) *ClientProfileAllOf {
	this := ClientProfileAllOf{}
	this.Name = name
	this.SpaKeyName = spaKeyName
	this.IdentityProviderName = identityProviderName
	return &this
}

// NewClientProfileAllOfWithDefaults instantiates a new ClientProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientProfileAllOfWithDefaults() *ClientProfileAllOf {
	this := ClientProfileAllOf{}
	return &this
}

// GetName returns the Name field value
func (o *ClientProfileAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClientProfileAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClientProfileAllOf) SetName(v string) {
	o.Name = v
}

// GetSpaKeyName returns the SpaKeyName field value
func (o *ClientProfileAllOf) GetSpaKeyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaKeyName
}

// GetSpaKeyNameOk returns a tuple with the SpaKeyName field value
// and a boolean to check if the value has been set.
func (o *ClientProfileAllOf) GetSpaKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaKeyName, true
}

// SetSpaKeyName sets field value
func (o *ClientProfileAllOf) SetSpaKeyName(v string) {
	o.SpaKeyName = v
}

// GetIdentityProviderName returns the IdentityProviderName field value
func (o *ClientProfileAllOf) GetIdentityProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityProviderName
}

// GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field value
// and a boolean to check if the value has been set.
func (o *ClientProfileAllOf) GetIdentityProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProviderName, true
}

// SetIdentityProviderName sets field value
func (o *ClientProfileAllOf) SetIdentityProviderName(v string) {
	o.IdentityProviderName = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ClientProfileAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfileAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ClientProfileAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ClientProfileAllOf) SetHostname(v string) {
	o.Hostname = &v
}

func (o ClientProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["spaKeyName"] = o.SpaKeyName
	}
	if true {
		toSerialize["identityProviderName"] = o.IdentityProviderName
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	return json.Marshal(toSerialize)
}

type NullableClientProfileAllOf struct {
	value *ClientProfileAllOf
	isSet bool
}

func (v NullableClientProfileAllOf) Get() *ClientProfileAllOf {
	return v.value
}

func (v *NullableClientProfileAllOf) Set(val *ClientProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClientProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClientProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientProfileAllOf(val *ClientProfileAllOf) *NullableClientProfileAllOf {
	return &NullableClientProfileAllOf{value: val, isSet: true}
}

func (v NullableClientProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
