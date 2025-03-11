/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v21+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 21.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DashboardSessionsDto struct for DashboardSessionsDto
type DashboardSessionsDto struct {
	// The total count of active sessions.
	TotalCount *int32 `json:"totalCount,omitempty"`
	// Geolocation data related to active sessions.
	GeoIpData []SessionGeoData                 `json:"geoIpData,omitempty"`
	Versions  *map[string]ClientVersionSupport `json:"versions,omitempty"`
}

// NewDashboardSessionsDto instantiates a new DashboardSessionsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardSessionsDto() *DashboardSessionsDto {
	this := DashboardSessionsDto{}
	return &this
}

// NewDashboardSessionsDtoWithDefaults instantiates a new DashboardSessionsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardSessionsDtoWithDefaults() *DashboardSessionsDto {
	this := DashboardSessionsDto{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *DashboardSessionsDto) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSessionsDto) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *DashboardSessionsDto) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *DashboardSessionsDto) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetGeoIpData returns the GeoIpData field value if set, zero value otherwise.
func (o *DashboardSessionsDto) GetGeoIpData() []SessionGeoData {
	if o == nil || o.GeoIpData == nil {
		var ret []SessionGeoData
		return ret
	}
	return o.GeoIpData
}

// GetGeoIpDataOk returns a tuple with the GeoIpData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSessionsDto) GetGeoIpDataOk() ([]SessionGeoData, bool) {
	if o == nil || o.GeoIpData == nil {
		return nil, false
	}
	return o.GeoIpData, true
}

// HasGeoIpData returns a boolean if a field has been set.
func (o *DashboardSessionsDto) HasGeoIpData() bool {
	if o != nil && o.GeoIpData != nil {
		return true
	}

	return false
}

// SetGeoIpData gets a reference to the given []SessionGeoData and assigns it to the GeoIpData field.
func (o *DashboardSessionsDto) SetGeoIpData(v []SessionGeoData) {
	o.GeoIpData = v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *DashboardSessionsDto) GetVersions() map[string]ClientVersionSupport {
	if o == nil || o.Versions == nil {
		var ret map[string]ClientVersionSupport
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSessionsDto) GetVersionsOk() (*map[string]ClientVersionSupport, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *DashboardSessionsDto) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given map[string]ClientVersionSupport and assigns it to the Versions field.
func (o *DashboardSessionsDto) SetVersions(v map[string]ClientVersionSupport) {
	o.Versions = &v
}

func (o DashboardSessionsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.GeoIpData != nil {
		toSerialize["geoIpData"] = o.GeoIpData
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardSessionsDto struct {
	value *DashboardSessionsDto
	isSet bool
}

func (v NullableDashboardSessionsDto) Get() *DashboardSessionsDto {
	return v.value
}

func (v *NullableDashboardSessionsDto) Set(val *DashboardSessionsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardSessionsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardSessionsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardSessionsDto(val *DashboardSessionsDto) *NullableDashboardSessionsDto {
	return &NullableDashboardSessionsDto{value: val, isSet: true}
}

func (v NullableDashboardSessionsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardSessionsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
