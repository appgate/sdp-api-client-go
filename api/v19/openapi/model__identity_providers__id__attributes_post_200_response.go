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

// IdentityProvidersIdAttributesPost200Response struct for IdentityProvidersIdAttributesPost200Response
type IdentityProvidersIdAttributesPost200Response struct {
	// The attributes received and unchanged by the Identity Provider.
	RawAttributes *map[string][]string `json:"rawAttributes,omitempty"`
	// The attributes received and mapped by the Identity Provider according to claimMappings.
	MappedAttributes *map[string]string `json:"mappedAttributes,omitempty"`
}

// NewIdentityProvidersIdAttributesPost200Response instantiates a new IdentityProvidersIdAttributesPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProvidersIdAttributesPost200Response() *IdentityProvidersIdAttributesPost200Response {
	this := IdentityProvidersIdAttributesPost200Response{}
	return &this
}

// NewIdentityProvidersIdAttributesPost200ResponseWithDefaults instantiates a new IdentityProvidersIdAttributesPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProvidersIdAttributesPost200ResponseWithDefaults() *IdentityProvidersIdAttributesPost200Response {
	this := IdentityProvidersIdAttributesPost200Response{}
	return &this
}

// GetRawAttributes returns the RawAttributes field value if set, zero value otherwise.
func (o *IdentityProvidersIdAttributesPost200Response) GetRawAttributes() map[string][]string {
	if o == nil || o.RawAttributes == nil {
		var ret map[string][]string
		return ret
	}
	return *o.RawAttributes
}

// GetRawAttributesOk returns a tuple with the RawAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersIdAttributesPost200Response) GetRawAttributesOk() (*map[string][]string, bool) {
	if o == nil || o.RawAttributes == nil {
		return nil, false
	}
	return o.RawAttributes, true
}

// HasRawAttributes returns a boolean if a field has been set.
func (o *IdentityProvidersIdAttributesPost200Response) HasRawAttributes() bool {
	if o != nil && o.RawAttributes != nil {
		return true
	}

	return false
}

// SetRawAttributes gets a reference to the given map[string][]string and assigns it to the RawAttributes field.
func (o *IdentityProvidersIdAttributesPost200Response) SetRawAttributes(v map[string][]string) {
	o.RawAttributes = &v
}

// GetMappedAttributes returns the MappedAttributes field value if set, zero value otherwise.
func (o *IdentityProvidersIdAttributesPost200Response) GetMappedAttributes() map[string]string {
	if o == nil || o.MappedAttributes == nil {
		var ret map[string]string
		return ret
	}
	return *o.MappedAttributes
}

// GetMappedAttributesOk returns a tuple with the MappedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersIdAttributesPost200Response) GetMappedAttributesOk() (*map[string]string, bool) {
	if o == nil || o.MappedAttributes == nil {
		return nil, false
	}
	return o.MappedAttributes, true
}

// HasMappedAttributes returns a boolean if a field has been set.
func (o *IdentityProvidersIdAttributesPost200Response) HasMappedAttributes() bool {
	if o != nil && o.MappedAttributes != nil {
		return true
	}

	return false
}

// SetMappedAttributes gets a reference to the given map[string]string and assigns it to the MappedAttributes field.
func (o *IdentityProvidersIdAttributesPost200Response) SetMappedAttributes(v map[string]string) {
	o.MappedAttributes = &v
}

func (o IdentityProvidersIdAttributesPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RawAttributes != nil {
		toSerialize["rawAttributes"] = o.RawAttributes
	}
	if o.MappedAttributes != nil {
		toSerialize["mappedAttributes"] = o.MappedAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityProvidersIdAttributesPost200Response struct {
	value *IdentityProvidersIdAttributesPost200Response
	isSet bool
}

func (v NullableIdentityProvidersIdAttributesPost200Response) Get() *IdentityProvidersIdAttributesPost200Response {
	return v.value
}

func (v *NullableIdentityProvidersIdAttributesPost200Response) Set(val *IdentityProvidersIdAttributesPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProvidersIdAttributesPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProvidersIdAttributesPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProvidersIdAttributesPost200Response(val *IdentityProvidersIdAttributesPost200Response) *NullableIdentityProvidersIdAttributesPost200Response {
	return &NullableIdentityProvidersIdAttributesPost200Response{value: val, isSet: true}
}

func (v NullableIdentityProvidersIdAttributesPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProvidersIdAttributesPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
