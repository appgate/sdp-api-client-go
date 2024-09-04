/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.3
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceWithStatusAllOfDetails Appliance status details.
type ApplianceWithStatusAllOfDetails struct {
	// The Appliance build version.
	Version *string      `json:"version,omitempty"`
	Cpu     *SystemInfo  `json:"cpu,omitempty"`
	Memory  *SystemInfo  `json:"memory,omitempty"`
	Disk    *SystemInfo  `json:"disk,omitempty"`
	Network *NetworkInfo `json:"network,omitempty"`
	// Appliance Status.
	Status *string `json:"status,omitempty"`
	Roles  *Roles  `json:"roles,omitempty"`
	// State of the Appliance. For internal use.
	State *string `json:"state,omitempty"`
	// The volume number in use.
	VolumeNumber *int32                                  `json:"volumeNumber,omitempty"`
	Upgrade      *ApplianceWithStatusAllOfDetailsUpgrade `json:"upgrade,omitempty"`
}

// NewApplianceWithStatusAllOfDetails instantiates a new ApplianceWithStatusAllOfDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceWithStatusAllOfDetails() *ApplianceWithStatusAllOfDetails {
	this := ApplianceWithStatusAllOfDetails{}
	return &this
}

// NewApplianceWithStatusAllOfDetailsWithDefaults instantiates a new ApplianceWithStatusAllOfDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceWithStatusAllOfDetailsWithDefaults() *ApplianceWithStatusAllOfDetails {
	this := ApplianceWithStatusAllOfDetails{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOfDetails) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOfDetails) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOfDetails) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApplianceWithStatusAllOfDetails) SetVersion(v string) {
	o.Version = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOfDetails) GetCpu() SystemInfo {
	if o == nil || o.Cpu == nil {
		var ret SystemInfo
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOfDetails) GetCpuOk() (*SystemInfo, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOfDetails) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given SystemInfo and assigns it to the Cpu field.
func (o *ApplianceWithStatusAllOfDetails) SetCpu(v SystemInfo) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOfDetails) GetMemory() SystemInfo {
	if o == nil || o.Memory == nil {
		var ret SystemInfo
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOfDetails) GetMemoryOk() (*SystemInfo, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOfDetails) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given SystemInfo and assigns it to the Memory field.
func (o *ApplianceWithStatusAllOfDetails) SetMemory(v SystemInfo) {
	o.Memory = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOfDetails) GetDisk() SystemInfo {
	if o == nil || o.Disk == nil {
		var ret SystemInfo
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOfDetails) GetDiskOk() (*SystemInfo, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOfDetails) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given SystemInfo and assigns it to the Disk field.
func (o *ApplianceWithStatusAllOfDetails) SetDisk(v SystemInfo) {
	o.Disk = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOfDetails) GetNetwork() NetworkInfo {
	if o == nil || o.Network == nil {
		var ret NetworkInfo
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOfDetails) GetNetworkOk() (*NetworkInfo, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOfDetails) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given NetworkInfo and assigns it to the Network field.
func (o *ApplianceWithStatusAllOfDetails) SetNetwork(v NetworkInfo) {
	o.Network = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOfDetails) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOfDetails) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOfDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceWithStatusAllOfDetails) SetStatus(v string) {
	o.Status = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOfDetails) GetRoles() Roles {
	if o == nil || o.Roles == nil {
		var ret Roles
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOfDetails) GetRolesOk() (*Roles, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOfDetails) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given Roles and assigns it to the Roles field.
func (o *ApplianceWithStatusAllOfDetails) SetRoles(v Roles) {
	o.Roles = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOfDetails) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOfDetails) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOfDetails) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ApplianceWithStatusAllOfDetails) SetState(v string) {
	o.State = &v
}

// GetVolumeNumber returns the VolumeNumber field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOfDetails) GetVolumeNumber() int32 {
	if o == nil || o.VolumeNumber == nil {
		var ret int32
		return ret
	}
	return *o.VolumeNumber
}

// GetVolumeNumberOk returns a tuple with the VolumeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOfDetails) GetVolumeNumberOk() (*int32, bool) {
	if o == nil || o.VolumeNumber == nil {
		return nil, false
	}
	return o.VolumeNumber, true
}

// HasVolumeNumber returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOfDetails) HasVolumeNumber() bool {
	if o != nil && o.VolumeNumber != nil {
		return true
	}

	return false
}

// SetVolumeNumber gets a reference to the given int32 and assigns it to the VolumeNumber field.
func (o *ApplianceWithStatusAllOfDetails) SetVolumeNumber(v int32) {
	o.VolumeNumber = &v
}

// GetUpgrade returns the Upgrade field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOfDetails) GetUpgrade() ApplianceWithStatusAllOfDetailsUpgrade {
	if o == nil || o.Upgrade == nil {
		var ret ApplianceWithStatusAllOfDetailsUpgrade
		return ret
	}
	return *o.Upgrade
}

// GetUpgradeOk returns a tuple with the Upgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOfDetails) GetUpgradeOk() (*ApplianceWithStatusAllOfDetailsUpgrade, bool) {
	if o == nil || o.Upgrade == nil {
		return nil, false
	}
	return o.Upgrade, true
}

// HasUpgrade returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOfDetails) HasUpgrade() bool {
	if o != nil && o.Upgrade != nil {
		return true
	}

	return false
}

// SetUpgrade gets a reference to the given ApplianceWithStatusAllOfDetailsUpgrade and assigns it to the Upgrade field.
func (o *ApplianceWithStatusAllOfDetails) SetUpgrade(v ApplianceWithStatusAllOfDetailsUpgrade) {
	o.Upgrade = &v
}

func (o ApplianceWithStatusAllOfDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Disk != nil {
		toSerialize["disk"] = o.Disk
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.VolumeNumber != nil {
		toSerialize["volumeNumber"] = o.VolumeNumber
	}
	if o.Upgrade != nil {
		toSerialize["upgrade"] = o.Upgrade
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceWithStatusAllOfDetails struct {
	value *ApplianceWithStatusAllOfDetails
	isSet bool
}

func (v NullableApplianceWithStatusAllOfDetails) Get() *ApplianceWithStatusAllOfDetails {
	return v.value
}

func (v *NullableApplianceWithStatusAllOfDetails) Set(val *ApplianceWithStatusAllOfDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceWithStatusAllOfDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceWithStatusAllOfDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceWithStatusAllOfDetails(val *ApplianceWithStatusAllOfDetails) *NullableApplianceWithStatusAllOfDetails {
	return &NullableApplianceWithStatusAllOfDetails{value: val, isSet: true}
}

func (v NullableApplianceWithStatusAllOfDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceWithStatusAllOfDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
