/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IpPoolAllOf Represents an IP Pool.
type IpPoolAllOf struct {
	// Whether the IP pool is for v4 or v6.
	IpVersion6     *bool              `json:"ipVersion6,omitempty"`
	Ranges         []IpPoolRangeInner `json:"ranges,omitempty"`
	ExcludedRanges []IpPoolRangeInner `json:"excludedRanges,omitempty"`
	// Number of days Allocated IPs will be reserved for device&users before they are reclaimable by others.
	LeaseTimeDays *int32 `json:"leaseTimeDays,omitempty"`
	// The total size of the IP Pool.
	Total *BigInt `json:"total,omitempty"`
	// Number of IPs in the pool are currently in use by device&users.
	CurrentlyUsed *int64 `json:"currentlyUsed,omitempty"`
	// Number of IPs in the pool are not currently in use but reserved for device&users according to the \"leaseTimeDays\" setting.
	Reserved *int64 `json:"reserved,omitempty"`
}

// NewIpPoolAllOf instantiates a new IpPoolAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpPoolAllOf() *IpPoolAllOf {
	this := IpPoolAllOf{}
	var ipVersion6 bool = false
	this.IpVersion6 = &ipVersion6
	var leaseTimeDays int32 = 30
	this.LeaseTimeDays = &leaseTimeDays
	return &this
}

// NewIpPoolAllOfWithDefaults instantiates a new IpPoolAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpPoolAllOfWithDefaults() *IpPoolAllOf {
	this := IpPoolAllOf{}
	var ipVersion6 bool = false
	this.IpVersion6 = &ipVersion6
	var leaseTimeDays int32 = 30
	this.LeaseTimeDays = &leaseTimeDays
	return &this
}

// GetIpVersion6 returns the IpVersion6 field value if set, zero value otherwise.
func (o *IpPoolAllOf) GetIpVersion6() bool {
	if o == nil || o.IpVersion6 == nil {
		var ret bool
		return ret
	}
	return *o.IpVersion6
}

// GetIpVersion6Ok returns a tuple with the IpVersion6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolAllOf) GetIpVersion6Ok() (*bool, bool) {
	if o == nil || o.IpVersion6 == nil {
		return nil, false
	}
	return o.IpVersion6, true
}

// HasIpVersion6 returns a boolean if a field has been set.
func (o *IpPoolAllOf) HasIpVersion6() bool {
	if o != nil && o.IpVersion6 != nil {
		return true
	}

	return false
}

// SetIpVersion6 gets a reference to the given bool and assigns it to the IpVersion6 field.
func (o *IpPoolAllOf) SetIpVersion6(v bool) {
	o.IpVersion6 = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *IpPoolAllOf) GetRanges() []IpPoolRangeInner {
	if o == nil || o.Ranges == nil {
		var ret []IpPoolRangeInner
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolAllOf) GetRangesOk() ([]IpPoolRangeInner, bool) {
	if o == nil || o.Ranges == nil {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *IpPoolAllOf) HasRanges() bool {
	if o != nil && o.Ranges != nil {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []IpPoolRangeInner and assigns it to the Ranges field.
func (o *IpPoolAllOf) SetRanges(v []IpPoolRangeInner) {
	o.Ranges = v
}

// GetExcludedRanges returns the ExcludedRanges field value if set, zero value otherwise.
func (o *IpPoolAllOf) GetExcludedRanges() []IpPoolRangeInner {
	if o == nil || o.ExcludedRanges == nil {
		var ret []IpPoolRangeInner
		return ret
	}
	return o.ExcludedRanges
}

// GetExcludedRangesOk returns a tuple with the ExcludedRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolAllOf) GetExcludedRangesOk() ([]IpPoolRangeInner, bool) {
	if o == nil || o.ExcludedRanges == nil {
		return nil, false
	}
	return o.ExcludedRanges, true
}

// HasExcludedRanges returns a boolean if a field has been set.
func (o *IpPoolAllOf) HasExcludedRanges() bool {
	if o != nil && o.ExcludedRanges != nil {
		return true
	}

	return false
}

// SetExcludedRanges gets a reference to the given []IpPoolRangeInner and assigns it to the ExcludedRanges field.
func (o *IpPoolAllOf) SetExcludedRanges(v []IpPoolRangeInner) {
	o.ExcludedRanges = v
}

// GetLeaseTimeDays returns the LeaseTimeDays field value if set, zero value otherwise.
func (o *IpPoolAllOf) GetLeaseTimeDays() int32 {
	if o == nil || o.LeaseTimeDays == nil {
		var ret int32
		return ret
	}
	return *o.LeaseTimeDays
}

// GetLeaseTimeDaysOk returns a tuple with the LeaseTimeDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolAllOf) GetLeaseTimeDaysOk() (*int32, bool) {
	if o == nil || o.LeaseTimeDays == nil {
		return nil, false
	}
	return o.LeaseTimeDays, true
}

// HasLeaseTimeDays returns a boolean if a field has been set.
func (o *IpPoolAllOf) HasLeaseTimeDays() bool {
	if o != nil && o.LeaseTimeDays != nil {
		return true
	}

	return false
}

// SetLeaseTimeDays gets a reference to the given int32 and assigns it to the LeaseTimeDays field.
func (o *IpPoolAllOf) SetLeaseTimeDays(v int32) {
	o.LeaseTimeDays = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *IpPoolAllOf) GetTotal() BigInt {
	if o == nil || o.Total == nil {
		var ret BigInt
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolAllOf) GetTotalOk() (*BigInt, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *IpPoolAllOf) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *IpPoolAllOf) SetTotal(v BigInt) {
	o.Total = &v
}

// GetCurrentlyUsed returns the CurrentlyUsed field value if set, zero value otherwise.
func (o *IpPoolAllOf) GetCurrentlyUsed() int64 {
	if o == nil || o.CurrentlyUsed == nil {
		var ret int64
		return ret
	}
	return *o.CurrentlyUsed
}

// GetCurrentlyUsedOk returns a tuple with the CurrentlyUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolAllOf) GetCurrentlyUsedOk() (*int64, bool) {
	if o == nil || o.CurrentlyUsed == nil {
		return nil, false
	}
	return o.CurrentlyUsed, true
}

// HasCurrentlyUsed returns a boolean if a field has been set.
func (o *IpPoolAllOf) HasCurrentlyUsed() bool {
	if o != nil && o.CurrentlyUsed != nil {
		return true
	}

	return false
}

// SetCurrentlyUsed gets a reference to the given int64 and assigns it to the CurrentlyUsed field.
func (o *IpPoolAllOf) SetCurrentlyUsed(v int64) {
	o.CurrentlyUsed = &v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *IpPoolAllOf) GetReserved() int64 {
	if o == nil || o.Reserved == nil {
		var ret int64
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolAllOf) GetReservedOk() (*int64, bool) {
	if o == nil || o.Reserved == nil {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *IpPoolAllOf) HasReserved() bool {
	if o != nil && o.Reserved != nil {
		return true
	}

	return false
}

// SetReserved gets a reference to the given int64 and assigns it to the Reserved field.
func (o *IpPoolAllOf) SetReserved(v int64) {
	o.Reserved = &v
}

func (o IpPoolAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpVersion6 != nil {
		toSerialize["ipVersion6"] = o.IpVersion6
	}
	if o.Ranges != nil {
		toSerialize["ranges"] = o.Ranges
	}
	if o.ExcludedRanges != nil {
		toSerialize["excludedRanges"] = o.ExcludedRanges
	}
	if o.LeaseTimeDays != nil {
		toSerialize["leaseTimeDays"] = o.LeaseTimeDays
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.CurrentlyUsed != nil {
		toSerialize["currentlyUsed"] = o.CurrentlyUsed
	}
	if o.Reserved != nil {
		toSerialize["reserved"] = o.Reserved
	}
	return json.Marshal(toSerialize)
}

type NullableIpPoolAllOf struct {
	value *IpPoolAllOf
	isSet bool
}

func (v NullableIpPoolAllOf) Get() *IpPoolAllOf {
	return v.value
}

func (v *NullableIpPoolAllOf) Set(val *IpPoolAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIpPoolAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIpPoolAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpPoolAllOf(val *IpPoolAllOf) *NullableIpPoolAllOf {
	return &NullableIpPoolAllOf{value: val, isSet: true}
}

func (v NullableIpPoolAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpPoolAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
