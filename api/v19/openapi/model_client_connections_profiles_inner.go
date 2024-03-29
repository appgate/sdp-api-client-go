/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v19+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 19.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ClientConnectionsProfilesInner struct for ClientConnectionsProfilesInner
type ClientConnectionsProfilesInner struct {
	// A name to identify the client profile. It will appear on the client UI.
	Name string `json:"name"`
	// SPA key name to be used in the profile. Same key names in different profiles will have the same SPA key. SPA key is used by the client to connect to the controllers.
	SpaKeyName string `json:"spaKeyName"`
	// Name of the Identity Provider to be used to authenticate.
	IdentityProviderName string `json:"identityProviderName"`
	// Connection URL for the profile.
	Url *string `json:"url,omitempty"`
}

// NewClientConnectionsProfilesInner instantiates a new ClientConnectionsProfilesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientConnectionsProfilesInner(name string, spaKeyName string, identityProviderName string) *ClientConnectionsProfilesInner {
	this := ClientConnectionsProfilesInner{}
	this.Name = name
	this.SpaKeyName = spaKeyName
	this.IdentityProviderName = identityProviderName
	return &this
}

// NewClientConnectionsProfilesInnerWithDefaults instantiates a new ClientConnectionsProfilesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientConnectionsProfilesInnerWithDefaults() *ClientConnectionsProfilesInner {
	this := ClientConnectionsProfilesInner{}
	return &this
}

// GetName returns the Name field value
func (o *ClientConnectionsProfilesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClientConnectionsProfilesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClientConnectionsProfilesInner) SetName(v string) {
	o.Name = v
}

// GetSpaKeyName returns the SpaKeyName field value
func (o *ClientConnectionsProfilesInner) GetSpaKeyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaKeyName
}

// GetSpaKeyNameOk returns a tuple with the SpaKeyName field value
// and a boolean to check if the value has been set.
func (o *ClientConnectionsProfilesInner) GetSpaKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaKeyName, true
}

// SetSpaKeyName sets field value
func (o *ClientConnectionsProfilesInner) SetSpaKeyName(v string) {
	o.SpaKeyName = v
}

// GetIdentityProviderName returns the IdentityProviderName field value
func (o *ClientConnectionsProfilesInner) GetIdentityProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityProviderName
}

// GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field value
// and a boolean to check if the value has been set.
func (o *ClientConnectionsProfilesInner) GetIdentityProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProviderName, true
}

// SetIdentityProviderName sets field value
func (o *ClientConnectionsProfilesInner) SetIdentityProviderName(v string) {
	o.IdentityProviderName = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ClientConnectionsProfilesInner) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConnectionsProfilesInner) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ClientConnectionsProfilesInner) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ClientConnectionsProfilesInner) SetUrl(v string) {
	o.Url = &v
}

func (o ClientConnectionsProfilesInner) MarshalJSON() ([]byte, error) {
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
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableClientConnectionsProfilesInner struct {
	value *ClientConnectionsProfilesInner
	isSet bool
}

func (v NullableClientConnectionsProfilesInner) Get() *ClientConnectionsProfilesInner {
	return v.value
}

func (v *NullableClientConnectionsProfilesInner) Set(val *ClientConnectionsProfilesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableClientConnectionsProfilesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableClientConnectionsProfilesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientConnectionsProfilesInner(val *ClientConnectionsProfilesInner) *NullableClientConnectionsProfilesInner {
	return &NullableClientConnectionsProfilesInner{value: val, isSet: true}
}

func (v NullableClientConnectionsProfilesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientConnectionsProfilesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
