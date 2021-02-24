/*
 * Appgate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the Integration chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin Interface (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-admin-interface-configure.html)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v14+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 14
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// InlineResponse2008 The name resolution status.
type InlineResponse2008 struct {
	// Dictionary of resource name and respective results.
	Resolutions *map[string]InlineResponse2008Resolutions `json:"resolutions,omitempty"`
}

// NewInlineResponse2008 instantiates a new InlineResponse2008 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2008() *InlineResponse2008 {
	this := InlineResponse2008{}
	return &this
}

// NewInlineResponse2008WithDefaults instantiates a new InlineResponse2008 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2008WithDefaults() *InlineResponse2008 {
	this := InlineResponse2008{}
	return &this
}

// GetResolutions returns the Resolutions field value if set, zero value otherwise.
func (o *InlineResponse2008) GetResolutions() map[string]InlineResponse2008Resolutions {
	if o == nil || o.Resolutions == nil {
		var ret map[string]InlineResponse2008Resolutions
		return ret
	}
	return *o.Resolutions
}

// GetResolutionsOk returns a tuple with the Resolutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008) GetResolutionsOk() (*map[string]InlineResponse2008Resolutions, bool) {
	if o == nil || o.Resolutions == nil {
		return nil, false
	}
	return o.Resolutions, true
}

// HasResolutions returns a boolean if a field has been set.
func (o *InlineResponse2008) HasResolutions() bool {
	if o != nil && o.Resolutions != nil {
		return true
	}

	return false
}

// SetResolutions gets a reference to the given map[string]InlineResponse2008Resolutions and assigns it to the Resolutions field.
func (o *InlineResponse2008) SetResolutions(v map[string]InlineResponse2008Resolutions) {
	o.Resolutions = &v
}

func (o InlineResponse2008) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Resolutions != nil {
		toSerialize["resolutions"] = o.Resolutions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2008 struct {
	value *InlineResponse2008
	isSet bool
}

func (v NullableInlineResponse2008) Get() *InlineResponse2008 {
	return v.value
}

func (v *NullableInlineResponse2008) Set(val *InlineResponse2008) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2008) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2008) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2008(val *InlineResponse2008) *NullableInlineResponse2008 {
	return &NullableInlineResponse2008{value: val, isSet: true}
}

func (v NullableInlineResponse2008) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2008) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
