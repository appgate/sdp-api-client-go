/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// AllocatedIp struct for AllocatedIp
type AllocatedIp struct {
	// ID of the IP Pool that has allocated this address.
	PoolId *string `json:"poolId,omitempty"`
	// Distinguished name of a device. Format: \"CN=,CN=,OU=\"
	DistinguishedName *string `json:"distinguishedName,omitempty"`
	// IP address either version 4 or 6 assigned to the tunnel.
	IpAddress *string `json:"ipAddress,omitempty"`
	// When the IP was allocated.
	AllocationTime *time.Time `json:"allocationTime,omitempty"`
	// When the IP allocation will be expired. It's equal to the last Entitlement Token's expiration date. Note that the allocation will still be reserved for the device & user according to the IP Pool settings.
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
}

// NewAllocatedIp instantiates a new AllocatedIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocatedIp() *AllocatedIp {
	this := AllocatedIp{}
	return &this
}

// NewAllocatedIpWithDefaults instantiates a new AllocatedIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocatedIpWithDefaults() *AllocatedIp {
	this := AllocatedIp{}
	return &this
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *AllocatedIp) GetPoolId() string {
	if o == nil || o.PoolId == nil {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedIp) GetPoolIdOk() (*string, bool) {
	if o == nil || o.PoolId == nil {
		return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *AllocatedIp) HasPoolId() bool {
	if o != nil && o.PoolId != nil {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *AllocatedIp) SetPoolId(v string) {
	o.PoolId = &v
}

// GetDistinguishedName returns the DistinguishedName field value if set, zero value otherwise.
func (o *AllocatedIp) GetDistinguishedName() string {
	if o == nil || o.DistinguishedName == nil {
		var ret string
		return ret
	}
	return *o.DistinguishedName
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedIp) GetDistinguishedNameOk() (*string, bool) {
	if o == nil || o.DistinguishedName == nil {
		return nil, false
	}
	return o.DistinguishedName, true
}

// HasDistinguishedName returns a boolean if a field has been set.
func (o *AllocatedIp) HasDistinguishedName() bool {
	if o != nil && o.DistinguishedName != nil {
		return true
	}

	return false
}

// SetDistinguishedName gets a reference to the given string and assigns it to the DistinguishedName field.
func (o *AllocatedIp) SetDistinguishedName(v string) {
	o.DistinguishedName = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *AllocatedIp) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedIp) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *AllocatedIp) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *AllocatedIp) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetAllocationTime returns the AllocationTime field value if set, zero value otherwise.
func (o *AllocatedIp) GetAllocationTime() time.Time {
	if o == nil || o.AllocationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AllocationTime
}

// GetAllocationTimeOk returns a tuple with the AllocationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedIp) GetAllocationTimeOk() (*time.Time, bool) {
	if o == nil || o.AllocationTime == nil {
		return nil, false
	}
	return o.AllocationTime, true
}

// HasAllocationTime returns a boolean if a field has been set.
func (o *AllocatedIp) HasAllocationTime() bool {
	if o != nil && o.AllocationTime != nil {
		return true
	}

	return false
}

// SetAllocationTime gets a reference to the given time.Time and assigns it to the AllocationTime field.
func (o *AllocatedIp) SetAllocationTime(v time.Time) {
	o.AllocationTime = &v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *AllocatedIp) GetExpirationTime() time.Time {
	if o == nil || o.ExpirationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedIp) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationTime == nil {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *AllocatedIp) HasExpirationTime() bool {
	if o != nil && o.ExpirationTime != nil {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *AllocatedIp) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

func (o AllocatedIp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PoolId != nil {
		toSerialize["poolId"] = o.PoolId
	}
	if o.DistinguishedName != nil {
		toSerialize["distinguishedName"] = o.DistinguishedName
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.AllocationTime != nil {
		toSerialize["allocationTime"] = o.AllocationTime
	}
	if o.ExpirationTime != nil {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	return json.Marshal(toSerialize)
}

type NullableAllocatedIp struct {
	value *AllocatedIp
	isSet bool
}

func (v NullableAllocatedIp) Get() *AllocatedIp {
	return v.value
}

func (v *NullableAllocatedIp) Set(val *AllocatedIp) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocatedIp) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocatedIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocatedIp(val *AllocatedIp) *NullableAllocatedIp {
	return &NullableAllocatedIp{value: val, isSet: true}
}

func (v NullableAllocatedIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocatedIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
