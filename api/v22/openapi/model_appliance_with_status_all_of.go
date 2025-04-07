/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceWithStatusAllOf struct for ApplianceWithStatusAllOf
type ApplianceWithStatusAllOf struct {
	// State of the Appliance. For internal use.
	State *string `json:"state,omitempty"`
	// Functions of the Appliance.
	Functions []ApplianceFunction `json:"functions,omitempty"`
	// Appliance Status.
	Status *string `json:"status,omitempty"`
	// The name of the customization applied to the Appliance.
	CustomizationName *string `json:"customizationName,omitempty"`
	// CPU utilization in percentage.
	Cpu *float32 `json:"cpu,omitempty"`
	// Memory utilization in percentage.
	Memory *float32 `json:"memory,omitempty"`
	// Disk utilization in percentage.
	Disk *float32 `json:"disk,omitempty"`
	// Number of sessions.
	NumberOfSessions *int32 `json:"numberOfSessions,omitempty"`
	// The Appliance build version.
	ApplianceVersion *string                          `json:"applianceVersion,omitempty"`
	Details          *ApplianceWithStatusAllOfDetails `json:"details,omitempty"`
}

// NewApplianceWithStatusAllOf instantiates a new ApplianceWithStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceWithStatusAllOf() *ApplianceWithStatusAllOf {
	this := ApplianceWithStatusAllOf{}
	return &this
}

// NewApplianceWithStatusAllOfWithDefaults instantiates a new ApplianceWithStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceWithStatusAllOfWithDefaults() *ApplianceWithStatusAllOf {
	this := ApplianceWithStatusAllOf{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ApplianceWithStatusAllOf) SetState(v string) {
	o.State = &v
}

// GetFunctions returns the Functions field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOf) GetFunctions() []ApplianceFunction {
	if o == nil || o.Functions == nil {
		var ret []ApplianceFunction
		return ret
	}
	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOf) GetFunctionsOk() ([]ApplianceFunction, bool) {
	if o == nil || o.Functions == nil {
		return nil, false
	}
	return o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOf) HasFunctions() bool {
	if o != nil && o.Functions != nil {
		return true
	}

	return false
}

// SetFunctions gets a reference to the given []ApplianceFunction and assigns it to the Functions field.
func (o *ApplianceWithStatusAllOf) SetFunctions(v []ApplianceFunction) {
	o.Functions = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceWithStatusAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetCustomizationName returns the CustomizationName field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOf) GetCustomizationName() string {
	if o == nil || o.CustomizationName == nil {
		var ret string
		return ret
	}
	return *o.CustomizationName
}

// GetCustomizationNameOk returns a tuple with the CustomizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOf) GetCustomizationNameOk() (*string, bool) {
	if o == nil || o.CustomizationName == nil {
		return nil, false
	}
	return o.CustomizationName, true
}

// HasCustomizationName returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOf) HasCustomizationName() bool {
	if o != nil && o.CustomizationName != nil {
		return true
	}

	return false
}

// SetCustomizationName gets a reference to the given string and assigns it to the CustomizationName field.
func (o *ApplianceWithStatusAllOf) SetCustomizationName(v string) {
	o.CustomizationName = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOf) GetCpu() float32 {
	if o == nil || o.Cpu == nil {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOf) GetCpuOk() (*float32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOf) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *ApplianceWithStatusAllOf) SetCpu(v float32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOf) GetMemory() float32 {
	if o == nil || o.Memory == nil {
		var ret float32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOf) GetMemoryOk() (*float32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOf) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given float32 and assigns it to the Memory field.
func (o *ApplianceWithStatusAllOf) SetMemory(v float32) {
	o.Memory = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOf) GetDisk() float32 {
	if o == nil || o.Disk == nil {
		var ret float32
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOf) GetDiskOk() (*float32, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOf) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given float32 and assigns it to the Disk field.
func (o *ApplianceWithStatusAllOf) SetDisk(v float32) {
	o.Disk = &v
}

// GetNumberOfSessions returns the NumberOfSessions field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOf) GetNumberOfSessions() int32 {
	if o == nil || o.NumberOfSessions == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfSessions
}

// GetNumberOfSessionsOk returns a tuple with the NumberOfSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOf) GetNumberOfSessionsOk() (*int32, bool) {
	if o == nil || o.NumberOfSessions == nil {
		return nil, false
	}
	return o.NumberOfSessions, true
}

// HasNumberOfSessions returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOf) HasNumberOfSessions() bool {
	if o != nil && o.NumberOfSessions != nil {
		return true
	}

	return false
}

// SetNumberOfSessions gets a reference to the given int32 and assigns it to the NumberOfSessions field.
func (o *ApplianceWithStatusAllOf) SetNumberOfSessions(v int32) {
	o.NumberOfSessions = &v
}

// GetApplianceVersion returns the ApplianceVersion field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOf) GetApplianceVersion() string {
	if o == nil || o.ApplianceVersion == nil {
		var ret string
		return ret
	}
	return *o.ApplianceVersion
}

// GetApplianceVersionOk returns a tuple with the ApplianceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOf) GetApplianceVersionOk() (*string, bool) {
	if o == nil || o.ApplianceVersion == nil {
		return nil, false
	}
	return o.ApplianceVersion, true
}

// HasApplianceVersion returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOf) HasApplianceVersion() bool {
	if o != nil && o.ApplianceVersion != nil {
		return true
	}

	return false
}

// SetApplianceVersion gets a reference to the given string and assigns it to the ApplianceVersion field.
func (o *ApplianceWithStatusAllOf) SetApplianceVersion(v string) {
	o.ApplianceVersion = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ApplianceWithStatusAllOf) GetDetails() ApplianceWithStatusAllOfDetails {
	if o == nil || o.Details == nil {
		var ret ApplianceWithStatusAllOfDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceWithStatusAllOf) GetDetailsOk() (*ApplianceWithStatusAllOfDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ApplianceWithStatusAllOf) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given ApplianceWithStatusAllOfDetails and assigns it to the Details field.
func (o *ApplianceWithStatusAllOf) SetDetails(v ApplianceWithStatusAllOfDetails) {
	o.Details = &v
}

func (o ApplianceWithStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Functions != nil {
		toSerialize["functions"] = o.Functions
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CustomizationName != nil {
		toSerialize["customizationName"] = o.CustomizationName
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
	if o.NumberOfSessions != nil {
		toSerialize["numberOfSessions"] = o.NumberOfSessions
	}
	if o.ApplianceVersion != nil {
		toSerialize["applianceVersion"] = o.ApplianceVersion
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceWithStatusAllOf struct {
	value *ApplianceWithStatusAllOf
	isSet bool
}

func (v NullableApplianceWithStatusAllOf) Get() *ApplianceWithStatusAllOf {
	return v.value
}

func (v *NullableApplianceWithStatusAllOf) Set(val *ApplianceWithStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceWithStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceWithStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceWithStatusAllOf(val *ApplianceWithStatusAllOf) *NullableApplianceWithStatusAllOf {
	return &NullableApplianceWithStatusAllOf{value: val, isSet: true}
}

func (v NullableApplianceWithStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceWithStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
