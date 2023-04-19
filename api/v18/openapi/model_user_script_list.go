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

// UserScriptList struct for UserScriptList
type UserScriptList struct {
	// 'The range applied to the list. Format: -/. 3-5/8 means, out of 8 count (query affects the total), the items between (including) the 3rd and the 5th are returned.'
	Range *string `json:"range,omitempty"`
	// The field name used to sort the list.
	OrderBy *string `json:"orderBy,omitempty"`
	// Whether the sorting is applied descending or ascending.
	Descending *bool `json:"descending,omitempty"`
	// The queries applied to the list.
	Queries []string `json:"queries,omitempty"`
	// The filters applied to the list.
	FilterBy []FilterBy `json:"filterBy,omitempty"`
	// List of User Claim Scripts.
	Data []UserScript `json:"data,omitempty"`
}

// NewUserScriptList instantiates a new UserScriptList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserScriptList() *UserScriptList {
	this := UserScriptList{}
	return &this
}

// NewUserScriptListWithDefaults instantiates a new UserScriptList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserScriptListWithDefaults() *UserScriptList {
	this := UserScriptList{}
	return &this
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *UserScriptList) GetRange() string {
	if o == nil || o.Range == nil {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScriptList) GetRangeOk() (*string, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *UserScriptList) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *UserScriptList) SetRange(v string) {
	o.Range = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *UserScriptList) GetOrderBy() string {
	if o == nil || o.OrderBy == nil {
		var ret string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScriptList) GetOrderByOk() (*string, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *UserScriptList) HasOrderBy() bool {
	if o != nil && o.OrderBy != nil {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given string and assigns it to the OrderBy field.
func (o *UserScriptList) SetOrderBy(v string) {
	o.OrderBy = &v
}

// GetDescending returns the Descending field value if set, zero value otherwise.
func (o *UserScriptList) GetDescending() bool {
	if o == nil || o.Descending == nil {
		var ret bool
		return ret
	}
	return *o.Descending
}

// GetDescendingOk returns a tuple with the Descending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScriptList) GetDescendingOk() (*bool, bool) {
	if o == nil || o.Descending == nil {
		return nil, false
	}
	return o.Descending, true
}

// HasDescending returns a boolean if a field has been set.
func (o *UserScriptList) HasDescending() bool {
	if o != nil && o.Descending != nil {
		return true
	}

	return false
}

// SetDescending gets a reference to the given bool and assigns it to the Descending field.
func (o *UserScriptList) SetDescending(v bool) {
	o.Descending = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *UserScriptList) GetQueries() []string {
	if o == nil || o.Queries == nil {
		var ret []string
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScriptList) GetQueriesOk() ([]string, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *UserScriptList) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []string and assigns it to the Queries field.
func (o *UserScriptList) SetQueries(v []string) {
	o.Queries = v
}

// GetFilterBy returns the FilterBy field value if set, zero value otherwise.
func (o *UserScriptList) GetFilterBy() []FilterBy {
	if o == nil || o.FilterBy == nil {
		var ret []FilterBy
		return ret
	}
	return o.FilterBy
}

// GetFilterByOk returns a tuple with the FilterBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScriptList) GetFilterByOk() ([]FilterBy, bool) {
	if o == nil || o.FilterBy == nil {
		return nil, false
	}
	return o.FilterBy, true
}

// HasFilterBy returns a boolean if a field has been set.
func (o *UserScriptList) HasFilterBy() bool {
	if o != nil && o.FilterBy != nil {
		return true
	}

	return false
}

// SetFilterBy gets a reference to the given []FilterBy and assigns it to the FilterBy field.
func (o *UserScriptList) SetFilterBy(v []FilterBy) {
	o.FilterBy = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserScriptList) GetData() []UserScript {
	if o == nil || o.Data == nil {
		var ret []UserScript
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserScriptList) GetDataOk() ([]UserScript, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserScriptList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []UserScript and assigns it to the Data field.
func (o *UserScriptList) SetData(v []UserScript) {
	o.Data = v
}

func (o UserScriptList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if o.OrderBy != nil {
		toSerialize["orderBy"] = o.OrderBy
	}
	if o.Descending != nil {
		toSerialize["descending"] = o.Descending
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.FilterBy != nil {
		toSerialize["filterBy"] = o.FilterBy
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUserScriptList struct {
	value *UserScriptList
	isSet bool
}

func (v NullableUserScriptList) Get() *UserScriptList {
	return v.value
}

func (v *NullableUserScriptList) Set(val *UserScriptList) {
	v.value = val
	v.isSet = true
}

func (v NullableUserScriptList) IsSet() bool {
	return v.isSet
}

func (v *NullableUserScriptList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserScriptList(val *UserScriptList) *NullableUserScriptList {
	return &NullableUserScriptList{value: val, isSet: true}
}

func (v NullableUserScriptList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserScriptList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
