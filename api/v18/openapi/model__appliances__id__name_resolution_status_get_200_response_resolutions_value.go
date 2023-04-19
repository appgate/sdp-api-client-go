/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.5
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue Resolution result.
type AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue struct {
	// Whether the resolution is for the resource is complete or partial.
	Partial *bool `json:"partial,omitempty"`
	// Completely resolved IPs.
	Finals []string `json:"finals,omitempty"`
	// Partially resolved names.
	Partials []string `json:"partials,omitempty"`
	// Errors occurred during resolution.
	Errors []string `json:"errors,omitempty"`
}

// NewAppliancesIdNameResolutionStatusGet200ResponseResolutionsValue instantiates a new AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliancesIdNameResolutionStatusGet200ResponseResolutionsValue() *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue {
	this := AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue{}
	return &this
}

// NewAppliancesIdNameResolutionStatusGet200ResponseResolutionsValueWithDefaults instantiates a new AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppliancesIdNameResolutionStatusGet200ResponseResolutionsValueWithDefaults() *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue {
	this := AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue{}
	return &this
}

// GetPartial returns the Partial field value if set, zero value otherwise.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) GetPartial() bool {
	if o == nil || o.Partial == nil {
		var ret bool
		return ret
	}
	return *o.Partial
}

// GetPartialOk returns a tuple with the Partial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) GetPartialOk() (*bool, bool) {
	if o == nil || o.Partial == nil {
		return nil, false
	}
	return o.Partial, true
}

// HasPartial returns a boolean if a field has been set.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) HasPartial() bool {
	if o != nil && o.Partial != nil {
		return true
	}

	return false
}

// SetPartial gets a reference to the given bool and assigns it to the Partial field.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) SetPartial(v bool) {
	o.Partial = &v
}

// GetFinals returns the Finals field value if set, zero value otherwise.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) GetFinals() []string {
	if o == nil || o.Finals == nil {
		var ret []string
		return ret
	}
	return o.Finals
}

// GetFinalsOk returns a tuple with the Finals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) GetFinalsOk() ([]string, bool) {
	if o == nil || o.Finals == nil {
		return nil, false
	}
	return o.Finals, true
}

// HasFinals returns a boolean if a field has been set.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) HasFinals() bool {
	if o != nil && o.Finals != nil {
		return true
	}

	return false
}

// SetFinals gets a reference to the given []string and assigns it to the Finals field.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) SetFinals(v []string) {
	o.Finals = v
}

// GetPartials returns the Partials field value if set, zero value otherwise.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) GetPartials() []string {
	if o == nil || o.Partials == nil {
		var ret []string
		return ret
	}
	return o.Partials
}

// GetPartialsOk returns a tuple with the Partials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) GetPartialsOk() ([]string, bool) {
	if o == nil || o.Partials == nil {
		return nil, false
	}
	return o.Partials, true
}

// HasPartials returns a boolean if a field has been set.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) HasPartials() bool {
	if o != nil && o.Partials != nil {
		return true
	}

	return false
}

// SetPartials gets a reference to the given []string and assigns it to the Partials field.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) SetPartials(v []string) {
	o.Partials = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) GetErrorsOk() ([]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) SetErrors(v []string) {
	o.Errors = v
}

func (o AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Partial != nil {
		toSerialize["partial"] = o.Partial
	}
	if o.Finals != nil {
		toSerialize["finals"] = o.Finals
	}
	if o.Partials != nil {
		toSerialize["partials"] = o.Partials
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableAppliancesIdNameResolutionStatusGet200ResponseResolutionsValue struct {
	value *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue
	isSet bool
}

func (v NullableAppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) Get() *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue {
	return v.value
}

func (v *NullableAppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) Set(val *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliancesIdNameResolutionStatusGet200ResponseResolutionsValue(val *AppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) *NullableAppliancesIdNameResolutionStatusGet200ResponseResolutionsValue {
	return &NullableAppliancesIdNameResolutionStatusGet200ResponseResolutionsValue{value: val, isSet: true}
}

func (v NullableAppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliancesIdNameResolutionStatusGet200ResponseResolutionsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
