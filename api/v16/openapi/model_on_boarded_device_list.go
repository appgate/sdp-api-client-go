/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OnBoardedDeviceList struct for OnBoardedDeviceList
type OnBoardedDeviceList struct {
	// The query applied to the list.
	Query *string `json:"query,omitempty"`
	// 'The range applied to the list. Format: -/. 3-5/8 means, out of 8 count (query affects the total), the items between (including) the 3rd and the 5th are returned.'
	Range *string `json:"range,omitempty"`
	// The field name used to sort the list.
	OrderBy *string `json:"orderBy,omitempty"`
	// Whether the sorting is applied descending or ascending.
	Descending *bool `json:"descending,omitempty"`
	// The filters applied to the list.
	FilterBy *[]FilterBy `json:"filterBy,omitempty"`
	// List of On-Boarded Devices.
	Data *[]OnBoardedDevice `json:"data,omitempty"`
}

// NewOnBoardedDeviceList instantiates a new OnBoardedDeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnBoardedDeviceList() *OnBoardedDeviceList {
	this := OnBoardedDeviceList{}
	return &this
}

// NewOnBoardedDeviceListWithDefaults instantiates a new OnBoardedDeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnBoardedDeviceListWithDefaults() *OnBoardedDeviceList {
	this := OnBoardedDeviceList{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *OnBoardedDeviceList) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDeviceList) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *OnBoardedDeviceList) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *OnBoardedDeviceList) SetQuery(v string) {
	o.Query = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *OnBoardedDeviceList) GetRange() string {
	if o == nil || o.Range == nil {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDeviceList) GetRangeOk() (*string, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *OnBoardedDeviceList) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *OnBoardedDeviceList) SetRange(v string) {
	o.Range = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *OnBoardedDeviceList) GetOrderBy() string {
	if o == nil || o.OrderBy == nil {
		var ret string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDeviceList) GetOrderByOk() (*string, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *OnBoardedDeviceList) HasOrderBy() bool {
	if o != nil && o.OrderBy != nil {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given string and assigns it to the OrderBy field.
func (o *OnBoardedDeviceList) SetOrderBy(v string) {
	o.OrderBy = &v
}

// GetDescending returns the Descending field value if set, zero value otherwise.
func (o *OnBoardedDeviceList) GetDescending() bool {
	if o == nil || o.Descending == nil {
		var ret bool
		return ret
	}
	return *o.Descending
}

// GetDescendingOk returns a tuple with the Descending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDeviceList) GetDescendingOk() (*bool, bool) {
	if o == nil || o.Descending == nil {
		return nil, false
	}
	return o.Descending, true
}

// HasDescending returns a boolean if a field has been set.
func (o *OnBoardedDeviceList) HasDescending() bool {
	if o != nil && o.Descending != nil {
		return true
	}

	return false
}

// SetDescending gets a reference to the given bool and assigns it to the Descending field.
func (o *OnBoardedDeviceList) SetDescending(v bool) {
	o.Descending = &v
}

// GetFilterBy returns the FilterBy field value if set, zero value otherwise.
func (o *OnBoardedDeviceList) GetFilterBy() []FilterBy {
	if o == nil || o.FilterBy == nil {
		var ret []FilterBy
		return ret
	}
	return *o.FilterBy
}

// GetFilterByOk returns a tuple with the FilterBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDeviceList) GetFilterByOk() (*[]FilterBy, bool) {
	if o == nil || o.FilterBy == nil {
		return nil, false
	}
	return o.FilterBy, true
}

// HasFilterBy returns a boolean if a field has been set.
func (o *OnBoardedDeviceList) HasFilterBy() bool {
	if o != nil && o.FilterBy != nil {
		return true
	}

	return false
}

// SetFilterBy gets a reference to the given []FilterBy and assigns it to the FilterBy field.
func (o *OnBoardedDeviceList) SetFilterBy(v []FilterBy) {
	o.FilterBy = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OnBoardedDeviceList) GetData() []OnBoardedDevice {
	if o == nil || o.Data == nil {
		var ret []OnBoardedDevice
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnBoardedDeviceList) GetDataOk() (*[]OnBoardedDevice, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OnBoardedDeviceList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []OnBoardedDevice and assigns it to the Data field.
func (o *OnBoardedDeviceList) SetData(v []OnBoardedDevice) {
	o.Data = &v
}

func (o OnBoardedDeviceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableOnBoardedDeviceList struct {
	value *OnBoardedDeviceList
	isSet bool
}

func (v NullableOnBoardedDeviceList) Get() *OnBoardedDeviceList {
	return v.value
}

func (v *NullableOnBoardedDeviceList) Set(val *OnBoardedDeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullableOnBoardedDeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullableOnBoardedDeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnBoardedDeviceList(val *OnBoardedDeviceList) *NullableOnBoardedDeviceList {
	return &NullableOnBoardedDeviceList{value: val, isSet: true}
}

func (v NullableOnBoardedDeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnBoardedDeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
