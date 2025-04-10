/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CaConfigAllOf CA details.
type CaConfigAllOf struct {
	// Name constraints extension details for permitted hostnames or IPs.
	NameConstraintsPermitted []string `json:"nameConstraintsPermitted,omitempty"`
	// Name constraints extension details for excluded hostnames or IPs.
	NameConstraintsExcluded []string `json:"nameConstraintsExcluded,omitempty"`
	// The URL where the CRL is hosted.
	CrlUrl *string `json:"crlUrl,omitempty"`
}

// NewCaConfigAllOf instantiates a new CaConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaConfigAllOf() *CaConfigAllOf {
	this := CaConfigAllOf{}
	return &this
}

// NewCaConfigAllOfWithDefaults instantiates a new CaConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaConfigAllOfWithDefaults() *CaConfigAllOf {
	this := CaConfigAllOf{}
	return &this
}

// GetNameConstraintsPermitted returns the NameConstraintsPermitted field value if set, zero value otherwise.
func (o *CaConfigAllOf) GetNameConstraintsPermitted() []string {
	if o == nil || o.NameConstraintsPermitted == nil {
		var ret []string
		return ret
	}
	return o.NameConstraintsPermitted
}

// GetNameConstraintsPermittedOk returns a tuple with the NameConstraintsPermitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfigAllOf) GetNameConstraintsPermittedOk() ([]string, bool) {
	if o == nil || o.NameConstraintsPermitted == nil {
		return nil, false
	}
	return o.NameConstraintsPermitted, true
}

// HasNameConstraintsPermitted returns a boolean if a field has been set.
func (o *CaConfigAllOf) HasNameConstraintsPermitted() bool {
	if o != nil && o.NameConstraintsPermitted != nil {
		return true
	}

	return false
}

// SetNameConstraintsPermitted gets a reference to the given []string and assigns it to the NameConstraintsPermitted field.
func (o *CaConfigAllOf) SetNameConstraintsPermitted(v []string) {
	o.NameConstraintsPermitted = v
}

// GetNameConstraintsExcluded returns the NameConstraintsExcluded field value if set, zero value otherwise.
func (o *CaConfigAllOf) GetNameConstraintsExcluded() []string {
	if o == nil || o.NameConstraintsExcluded == nil {
		var ret []string
		return ret
	}
	return o.NameConstraintsExcluded
}

// GetNameConstraintsExcludedOk returns a tuple with the NameConstraintsExcluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfigAllOf) GetNameConstraintsExcludedOk() ([]string, bool) {
	if o == nil || o.NameConstraintsExcluded == nil {
		return nil, false
	}
	return o.NameConstraintsExcluded, true
}

// HasNameConstraintsExcluded returns a boolean if a field has been set.
func (o *CaConfigAllOf) HasNameConstraintsExcluded() bool {
	if o != nil && o.NameConstraintsExcluded != nil {
		return true
	}

	return false
}

// SetNameConstraintsExcluded gets a reference to the given []string and assigns it to the NameConstraintsExcluded field.
func (o *CaConfigAllOf) SetNameConstraintsExcluded(v []string) {
	o.NameConstraintsExcluded = v
}

// GetCrlUrl returns the CrlUrl field value if set, zero value otherwise.
func (o *CaConfigAllOf) GetCrlUrl() string {
	if o == nil || o.CrlUrl == nil {
		var ret string
		return ret
	}
	return *o.CrlUrl
}

// GetCrlUrlOk returns a tuple with the CrlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaConfigAllOf) GetCrlUrlOk() (*string, bool) {
	if o == nil || o.CrlUrl == nil {
		return nil, false
	}
	return o.CrlUrl, true
}

// HasCrlUrl returns a boolean if a field has been set.
func (o *CaConfigAllOf) HasCrlUrl() bool {
	if o != nil && o.CrlUrl != nil {
		return true
	}

	return false
}

// SetCrlUrl gets a reference to the given string and assigns it to the CrlUrl field.
func (o *CaConfigAllOf) SetCrlUrl(v string) {
	o.CrlUrl = &v
}

func (o CaConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NameConstraintsPermitted != nil {
		toSerialize["nameConstraintsPermitted"] = o.NameConstraintsPermitted
	}
	if o.NameConstraintsExcluded != nil {
		toSerialize["nameConstraintsExcluded"] = o.NameConstraintsExcluded
	}
	if o.CrlUrl != nil {
		toSerialize["crlUrl"] = o.CrlUrl
	}
	return json.Marshal(toSerialize)
}

type NullableCaConfigAllOf struct {
	value *CaConfigAllOf
	isSet bool
}

func (v NullableCaConfigAllOf) Get() *CaConfigAllOf {
	return v.value
}

func (v *NullableCaConfigAllOf) Set(val *CaConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCaConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCaConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaConfigAllOf(val *CaConfigAllOf) *NullableCaConfigAllOf {
	return &NullableCaConfigAllOf{value: val, isSet: true}
}

func (v NullableCaConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
