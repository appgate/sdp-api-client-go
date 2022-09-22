/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Fido2DeviceList struct for Fido2DeviceList
type Fido2DeviceList struct {
	// The queries applied to the list.
	Queries []string `json:"queries,omitempty"`
	// The first query applied to the list. Deprecated as of 6.0. Use queries instead.
	// Deprecated
	Query *string `json:"query,omitempty"`
	// 'The range applied to the list. Format: -/. 3-5/8 means, out of 8 count (query affects the total), the items between (including) the 3rd and the 5th are returned.'
	Range *string `json:"range,omitempty"`
	// The field name used to sort the list.
	OrderBy *string `json:"orderBy,omitempty"`
	// Whether the sorting is applied descending or ascending.
	Descending *bool `json:"descending,omitempty"`
	// The filters applied to the list.
	FilterBy []FilterBy `json:"filterBy,omitempty"`
	// List of registered FIDO2 Devices.
	Data []Fido2Device `json:"data,omitempty"`
}

// NewFido2DeviceList instantiates a new Fido2DeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFido2DeviceList() *Fido2DeviceList {
	this := Fido2DeviceList{}
	return &this
}

// NewFido2DeviceListWithDefaults instantiates a new Fido2DeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFido2DeviceListWithDefaults() *Fido2DeviceList {
	this := Fido2DeviceList{}
	return &this
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *Fido2DeviceList) GetQueries() []string {
	if o == nil || o.Queries == nil {
		var ret []string
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fido2DeviceList) GetQueriesOk() ([]string, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *Fido2DeviceList) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []string and assigns it to the Queries field.
func (o *Fido2DeviceList) SetQueries(v []string) {
	o.Queries = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
// Deprecated
func (o *Fido2DeviceList) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Fido2DeviceList) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *Fido2DeviceList) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
// Deprecated
func (o *Fido2DeviceList) SetQuery(v string) {
	o.Query = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *Fido2DeviceList) GetRange() string {
	if o == nil || o.Range == nil {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fido2DeviceList) GetRangeOk() (*string, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *Fido2DeviceList) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *Fido2DeviceList) SetRange(v string) {
	o.Range = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *Fido2DeviceList) GetOrderBy() string {
	if o == nil || o.OrderBy == nil {
		var ret string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fido2DeviceList) GetOrderByOk() (*string, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *Fido2DeviceList) HasOrderBy() bool {
	if o != nil && o.OrderBy != nil {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given string and assigns it to the OrderBy field.
func (o *Fido2DeviceList) SetOrderBy(v string) {
	o.OrderBy = &v
}

// GetDescending returns the Descending field value if set, zero value otherwise.
func (o *Fido2DeviceList) GetDescending() bool {
	if o == nil || o.Descending == nil {
		var ret bool
		return ret
	}
	return *o.Descending
}

// GetDescendingOk returns a tuple with the Descending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fido2DeviceList) GetDescendingOk() (*bool, bool) {
	if o == nil || o.Descending == nil {
		return nil, false
	}
	return o.Descending, true
}

// HasDescending returns a boolean if a field has been set.
func (o *Fido2DeviceList) HasDescending() bool {
	if o != nil && o.Descending != nil {
		return true
	}

	return false
}

// SetDescending gets a reference to the given bool and assigns it to the Descending field.
func (o *Fido2DeviceList) SetDescending(v bool) {
	o.Descending = &v
}

// GetFilterBy returns the FilterBy field value if set, zero value otherwise.
func (o *Fido2DeviceList) GetFilterBy() []FilterBy {
	if o == nil || o.FilterBy == nil {
		var ret []FilterBy
		return ret
	}
	return o.FilterBy
}

// GetFilterByOk returns a tuple with the FilterBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fido2DeviceList) GetFilterByOk() ([]FilterBy, bool) {
	if o == nil || o.FilterBy == nil {
		return nil, false
	}
	return o.FilterBy, true
}

// HasFilterBy returns a boolean if a field has been set.
func (o *Fido2DeviceList) HasFilterBy() bool {
	if o != nil && o.FilterBy != nil {
		return true
	}

	return false
}

// SetFilterBy gets a reference to the given []FilterBy and assigns it to the FilterBy field.
func (o *Fido2DeviceList) SetFilterBy(v []FilterBy) {
	o.FilterBy = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Fido2DeviceList) GetData() []Fido2Device {
	if o == nil || o.Data == nil {
		var ret []Fido2Device
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fido2DeviceList) GetDataOk() ([]Fido2Device, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Fido2DeviceList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Fido2Device and assigns it to the Data field.
func (o *Fido2DeviceList) SetData(v []Fido2Device) {
	o.Data = v
}

func (o Fido2DeviceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if o.OrderBy != nil {
		toSerialize["orderBy"] = o.OrderBy
	}
	if o.Descending != nil {
		toSerialize["descending"] = o.Descending
	}
	if o.FilterBy != nil {
		toSerialize["filterBy"] = o.FilterBy
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableFido2DeviceList struct {
	value *Fido2DeviceList
	isSet bool
}

func (v NullableFido2DeviceList) Get() *Fido2DeviceList {
	return v.value
}

func (v *NullableFido2DeviceList) Set(val *Fido2DeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullableFido2DeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullableFido2DeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFido2DeviceList(val *Fido2DeviceList) *NullableFido2DeviceList {
	return &NullableFido2DeviceList{value: val, isSet: true}
}

func (v NullableFido2DeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFido2DeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
