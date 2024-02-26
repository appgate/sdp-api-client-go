/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ActiveSessions struct for ActiveSessions
type ActiveSessions struct {
	// User-friendly name for the stats.
	Name *string `json:"name,omitempty"`
	// The time the stats was generated.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// Recommended refresh interval in minutes.
	RefreshInterval *float32 `json:"refreshInterval,omitempty"`
	// The range applied to the list. Format: -/. 3-5/8 means, out of 8 count (query affects the total), the items between (including) the 3rd and the 5th are returned.
	Range *string `json:"range,omitempty"`
	// The field name used to sort the list.
	OrderBy *string `json:"orderBy,omitempty"`
	// Whether the sorting is applied descending or ascending.
	Descending *bool `json:"descending,omitempty"`
	// The queries applied to the list.
	Queries []string `json:"queries,omitempty"`
	// The total readable count of entities. Not influenced by the query.
	TotalCount *int32 `json:"totalCount,omitempty"`
	// The filters applied to the list.
	FilterBy []FilterBy `json:"filterBy,omitempty"`
	// The number of total Distinct Users currently active.
	DistinctUserCount *float32            `json:"distinctUserCount,omitempty"`
	Data              []BaseActiveSession `json:"data,omitempty"`
	GeolocationQuery  *GeolocationQuery   `json:"geolocationQuery,omitempty"`
}

// NewActiveSessions instantiates a new ActiveSessions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveSessions() *ActiveSessions {
	this := ActiveSessions{}
	return &this
}

// NewActiveSessionsWithDefaults instantiates a new ActiveSessions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveSessionsWithDefaults() *ActiveSessions {
	this := ActiveSessions{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ActiveSessions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ActiveSessions) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ActiveSessions) SetName(v string) {
	o.Name = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *ActiveSessions) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessions) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *ActiveSessions) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *ActiveSessions) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetRefreshInterval returns the RefreshInterval field value if set, zero value otherwise.
func (o *ActiveSessions) GetRefreshInterval() float32 {
	if o == nil || o.RefreshInterval == nil {
		var ret float32
		return ret
	}
	return *o.RefreshInterval
}

// GetRefreshIntervalOk returns a tuple with the RefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessions) GetRefreshIntervalOk() (*float32, bool) {
	if o == nil || o.RefreshInterval == nil {
		return nil, false
	}
	return o.RefreshInterval, true
}

// HasRefreshInterval returns a boolean if a field has been set.
func (o *ActiveSessions) HasRefreshInterval() bool {
	if o != nil && o.RefreshInterval != nil {
		return true
	}

	return false
}

// SetRefreshInterval gets a reference to the given float32 and assigns it to the RefreshInterval field.
func (o *ActiveSessions) SetRefreshInterval(v float32) {
	o.RefreshInterval = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *ActiveSessions) GetRange() string {
	if o == nil || o.Range == nil {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessions) GetRangeOk() (*string, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *ActiveSessions) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *ActiveSessions) SetRange(v string) {
	o.Range = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *ActiveSessions) GetOrderBy() string {
	if o == nil || o.OrderBy == nil {
		var ret string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessions) GetOrderByOk() (*string, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *ActiveSessions) HasOrderBy() bool {
	if o != nil && o.OrderBy != nil {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given string and assigns it to the OrderBy field.
func (o *ActiveSessions) SetOrderBy(v string) {
	o.OrderBy = &v
}

// GetDescending returns the Descending field value if set, zero value otherwise.
func (o *ActiveSessions) GetDescending() bool {
	if o == nil || o.Descending == nil {
		var ret bool
		return ret
	}
	return *o.Descending
}

// GetDescendingOk returns a tuple with the Descending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessions) GetDescendingOk() (*bool, bool) {
	if o == nil || o.Descending == nil {
		return nil, false
	}
	return o.Descending, true
}

// HasDescending returns a boolean if a field has been set.
func (o *ActiveSessions) HasDescending() bool {
	if o != nil && o.Descending != nil {
		return true
	}

	return false
}

// SetDescending gets a reference to the given bool and assigns it to the Descending field.
func (o *ActiveSessions) SetDescending(v bool) {
	o.Descending = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *ActiveSessions) GetQueries() []string {
	if o == nil || o.Queries == nil {
		var ret []string
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessions) GetQueriesOk() ([]string, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *ActiveSessions) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []string and assigns it to the Queries field.
func (o *ActiveSessions) SetQueries(v []string) {
	o.Queries = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ActiveSessions) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessions) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ActiveSessions) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ActiveSessions) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetFilterBy returns the FilterBy field value if set, zero value otherwise.
func (o *ActiveSessions) GetFilterBy() []FilterBy {
	if o == nil || o.FilterBy == nil {
		var ret []FilterBy
		return ret
	}
	return o.FilterBy
}

// GetFilterByOk returns a tuple with the FilterBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessions) GetFilterByOk() ([]FilterBy, bool) {
	if o == nil || o.FilterBy == nil {
		return nil, false
	}
	return o.FilterBy, true
}

// HasFilterBy returns a boolean if a field has been set.
func (o *ActiveSessions) HasFilterBy() bool {
	if o != nil && o.FilterBy != nil {
		return true
	}

	return false
}

// SetFilterBy gets a reference to the given []FilterBy and assigns it to the FilterBy field.
func (o *ActiveSessions) SetFilterBy(v []FilterBy) {
	o.FilterBy = v
}

// GetDistinctUserCount returns the DistinctUserCount field value if set, zero value otherwise.
func (o *ActiveSessions) GetDistinctUserCount() float32 {
	if o == nil || o.DistinctUserCount == nil {
		var ret float32
		return ret
	}
	return *o.DistinctUserCount
}

// GetDistinctUserCountOk returns a tuple with the DistinctUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessions) GetDistinctUserCountOk() (*float32, bool) {
	if o == nil || o.DistinctUserCount == nil {
		return nil, false
	}
	return o.DistinctUserCount, true
}

// HasDistinctUserCount returns a boolean if a field has been set.
func (o *ActiveSessions) HasDistinctUserCount() bool {
	if o != nil && o.DistinctUserCount != nil {
		return true
	}

	return false
}

// SetDistinctUserCount gets a reference to the given float32 and assigns it to the DistinctUserCount field.
func (o *ActiveSessions) SetDistinctUserCount(v float32) {
	o.DistinctUserCount = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ActiveSessions) GetData() []BaseActiveSession {
	if o == nil || o.Data == nil {
		var ret []BaseActiveSession
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessions) GetDataOk() ([]BaseActiveSession, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ActiveSessions) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []BaseActiveSession and assigns it to the Data field.
func (o *ActiveSessions) SetData(v []BaseActiveSession) {
	o.Data = v
}

// GetGeolocationQuery returns the GeolocationQuery field value if set, zero value otherwise.
func (o *ActiveSessions) GetGeolocationQuery() GeolocationQuery {
	if o == nil || o.GeolocationQuery == nil {
		var ret GeolocationQuery
		return ret
	}
	return *o.GeolocationQuery
}

// GetGeolocationQueryOk returns a tuple with the GeolocationQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveSessions) GetGeolocationQueryOk() (*GeolocationQuery, bool) {
	if o == nil || o.GeolocationQuery == nil {
		return nil, false
	}
	return o.GeolocationQuery, true
}

// HasGeolocationQuery returns a boolean if a field has been set.
func (o *ActiveSessions) HasGeolocationQuery() bool {
	if o != nil && o.GeolocationQuery != nil {
		return true
	}

	return false
}

// SetGeolocationQuery gets a reference to the given GeolocationQuery and assigns it to the GeolocationQuery field.
func (o *ActiveSessions) SetGeolocationQuery(v GeolocationQuery) {
	o.GeolocationQuery = &v
}

func (o ActiveSessions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.RefreshInterval != nil {
		toSerialize["refreshInterval"] = o.RefreshInterval
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
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.FilterBy != nil {
		toSerialize["filterBy"] = o.FilterBy
	}
	if o.DistinctUserCount != nil {
		toSerialize["distinctUserCount"] = o.DistinctUserCount
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.GeolocationQuery != nil {
		toSerialize["geolocationQuery"] = o.GeolocationQuery
	}
	return json.Marshal(toSerialize)
}

type NullableActiveSessions struct {
	value *ActiveSessions
	isSet bool
}

func (v NullableActiveSessions) Get() *ActiveSessions {
	return v.value
}

func (v *NullableActiveSessions) Set(val *ActiveSessions) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveSessions) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveSessions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveSessions(val *ActiveSessions) *NullableActiveSessions {
	return &NullableActiveSessions{value: val, isSet: true}
}

func (v NullableActiveSessions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveSessions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
