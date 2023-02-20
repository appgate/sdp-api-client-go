/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StatsAppliancesListAllOfDiskInfo Disk usage information.
type StatsAppliancesListAllOfDiskInfo struct {
	// Total disk space in bytes.
	Total *float32 `json:"total,omitempty"`
	// Used disk space in bytes.
	Used *float32 `json:"used,omitempty"`
	// Free disk space in bytes.
	Free *float32 `json:"free,omitempty"`
}

// NewStatsAppliancesListAllOfDiskInfo instantiates a new StatsAppliancesListAllOfDiskInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsAppliancesListAllOfDiskInfo() *StatsAppliancesListAllOfDiskInfo {
	this := StatsAppliancesListAllOfDiskInfo{}
	return &this
}

// NewStatsAppliancesListAllOfDiskInfoWithDefaults instantiates a new StatsAppliancesListAllOfDiskInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsAppliancesListAllOfDiskInfoWithDefaults() *StatsAppliancesListAllOfDiskInfo {
	this := StatsAppliancesListAllOfDiskInfo{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfDiskInfo) GetTotal() float32 {
	if o == nil || o.Total == nil {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfDiskInfo) GetTotalOk() (*float32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfDiskInfo) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *StatsAppliancesListAllOfDiskInfo) SetTotal(v float32) {
	o.Total = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfDiskInfo) GetUsed() float32 {
	if o == nil || o.Used == nil {
		var ret float32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfDiskInfo) GetUsedOk() (*float32, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfDiskInfo) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given float32 and assigns it to the Used field.
func (o *StatsAppliancesListAllOfDiskInfo) SetUsed(v float32) {
	o.Used = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfDiskInfo) GetFree() float32 {
	if o == nil || o.Free == nil {
		var ret float32
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfDiskInfo) GetFreeOk() (*float32, bool) {
	if o == nil || o.Free == nil {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfDiskInfo) HasFree() bool {
	if o != nil && o.Free != nil {
		return true
	}

	return false
}

// SetFree gets a reference to the given float32 and assigns it to the Free field.
func (o *StatsAppliancesListAllOfDiskInfo) SetFree(v float32) {
	o.Free = &v
}

func (o StatsAppliancesListAllOfDiskInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Used != nil {
		toSerialize["used"] = o.Used
	}
	if o.Free != nil {
		toSerialize["free"] = o.Free
	}
	return json.Marshal(toSerialize)
}

type NullableStatsAppliancesListAllOfDiskInfo struct {
	value *StatsAppliancesListAllOfDiskInfo
	isSet bool
}

func (v NullableStatsAppliancesListAllOfDiskInfo) Get() *StatsAppliancesListAllOfDiskInfo {
	return v.value
}

func (v *NullableStatsAppliancesListAllOfDiskInfo) Set(val *StatsAppliancesListAllOfDiskInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsAppliancesListAllOfDiskInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsAppliancesListAllOfDiskInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsAppliancesListAllOfDiskInfo(val *StatsAppliancesListAllOfDiskInfo) *NullableStatsAppliancesListAllOfDiskInfo {
	return &NullableStatsAppliancesListAllOfDiskInfo{value: val, isSet: true}
}

func (v NullableStatsAppliancesListAllOfDiskInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsAppliancesListAllOfDiskInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
