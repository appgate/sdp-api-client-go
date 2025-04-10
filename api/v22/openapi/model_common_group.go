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

// CommonGroup struct for CommonGroup
type CommonGroup struct {
	// The name of the group.
	Name          *string  `json:"name,omitempty"`
	ProviderNames []string `json:"providerNames,omitempty"`
}

// NewCommonGroup instantiates a new CommonGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonGroup() *CommonGroup {
	this := CommonGroup{}
	return &this
}

// NewCommonGroupWithDefaults instantiates a new CommonGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonGroupWithDefaults() *CommonGroup {
	this := CommonGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CommonGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CommonGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CommonGroup) SetName(v string) {
	o.Name = &v
}

// GetProviderNames returns the ProviderNames field value if set, zero value otherwise.
func (o *CommonGroup) GetProviderNames() []string {
	if o == nil || o.ProviderNames == nil {
		var ret []string
		return ret
	}
	return o.ProviderNames
}

// GetProviderNamesOk returns a tuple with the ProviderNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonGroup) GetProviderNamesOk() ([]string, bool) {
	if o == nil || o.ProviderNames == nil {
		return nil, false
	}
	return o.ProviderNames, true
}

// HasProviderNames returns a boolean if a field has been set.
func (o *CommonGroup) HasProviderNames() bool {
	if o != nil && o.ProviderNames != nil {
		return true
	}

	return false
}

// SetProviderNames gets a reference to the given []string and assigns it to the ProviderNames field.
func (o *CommonGroup) SetProviderNames(v []string) {
	o.ProviderNames = v
}

func (o CommonGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProviderNames != nil {
		toSerialize["providerNames"] = o.ProviderNames
	}
	return json.Marshal(toSerialize)
}

type NullableCommonGroup struct {
	value *CommonGroup
	isSet bool
}

func (v NullableCommonGroup) Get() *CommonGroup {
	return v.value
}

func (v *NullableCommonGroup) Set(val *CommonGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonGroup(val *CommonGroup) *NullableCommonGroup {
	return &NullableCommonGroup{value: val, isSet: true}
}

func (v NullableCommonGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
