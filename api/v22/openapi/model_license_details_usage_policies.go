/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LicenseDetailsUsagePolicies The amount of Policies in the system at present broken out by type.
type LicenseDetailsUsagePolicies struct {
	Access *float32 `json:"access,omitempty"`
	Stop   *float32 `json:"stop,omitempty"`
	Device *float32 `json:"device,omitempty"`
	Dns    *float32 `json:"dns,omitempty"`
	Admin  *float32 `json:"admin,omitempty"`
	Mixed  *float32 `json:"mixed,omitempty"`
}

// NewLicenseDetailsUsagePolicies instantiates a new LicenseDetailsUsagePolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseDetailsUsagePolicies() *LicenseDetailsUsagePolicies {
	this := LicenseDetailsUsagePolicies{}
	return &this
}

// NewLicenseDetailsUsagePoliciesWithDefaults instantiates a new LicenseDetailsUsagePolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseDetailsUsagePoliciesWithDefaults() *LicenseDetailsUsagePolicies {
	this := LicenseDetailsUsagePolicies{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *LicenseDetailsUsagePolicies) GetAccess() float32 {
	if o == nil || o.Access == nil {
		var ret float32
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsagePolicies) GetAccessOk() (*float32, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *LicenseDetailsUsagePolicies) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given float32 and assigns it to the Access field.
func (o *LicenseDetailsUsagePolicies) SetAccess(v float32) {
	o.Access = &v
}

// GetStop returns the Stop field value if set, zero value otherwise.
func (o *LicenseDetailsUsagePolicies) GetStop() float32 {
	if o == nil || o.Stop == nil {
		var ret float32
		return ret
	}
	return *o.Stop
}

// GetStopOk returns a tuple with the Stop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsagePolicies) GetStopOk() (*float32, bool) {
	if o == nil || o.Stop == nil {
		return nil, false
	}
	return o.Stop, true
}

// HasStop returns a boolean if a field has been set.
func (o *LicenseDetailsUsagePolicies) HasStop() bool {
	if o != nil && o.Stop != nil {
		return true
	}

	return false
}

// SetStop gets a reference to the given float32 and assigns it to the Stop field.
func (o *LicenseDetailsUsagePolicies) SetStop(v float32) {
	o.Stop = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *LicenseDetailsUsagePolicies) GetDevice() float32 {
	if o == nil || o.Device == nil {
		var ret float32
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsagePolicies) GetDeviceOk() (*float32, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *LicenseDetailsUsagePolicies) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given float32 and assigns it to the Device field.
func (o *LicenseDetailsUsagePolicies) SetDevice(v float32) {
	o.Device = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *LicenseDetailsUsagePolicies) GetDns() float32 {
	if o == nil || o.Dns == nil {
		var ret float32
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsagePolicies) GetDnsOk() (*float32, bool) {
	if o == nil || o.Dns == nil {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *LicenseDetailsUsagePolicies) HasDns() bool {
	if o != nil && o.Dns != nil {
		return true
	}

	return false
}

// SetDns gets a reference to the given float32 and assigns it to the Dns field.
func (o *LicenseDetailsUsagePolicies) SetDns(v float32) {
	o.Dns = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *LicenseDetailsUsagePolicies) GetAdmin() float32 {
	if o == nil || o.Admin == nil {
		var ret float32
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsagePolicies) GetAdminOk() (*float32, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *LicenseDetailsUsagePolicies) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given float32 and assigns it to the Admin field.
func (o *LicenseDetailsUsagePolicies) SetAdmin(v float32) {
	o.Admin = &v
}

// GetMixed returns the Mixed field value if set, zero value otherwise.
func (o *LicenseDetailsUsagePolicies) GetMixed() float32 {
	if o == nil || o.Mixed == nil {
		var ret float32
		return ret
	}
	return *o.Mixed
}

// GetMixedOk returns a tuple with the Mixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsagePolicies) GetMixedOk() (*float32, bool) {
	if o == nil || o.Mixed == nil {
		return nil, false
	}
	return o.Mixed, true
}

// HasMixed returns a boolean if a field has been set.
func (o *LicenseDetailsUsagePolicies) HasMixed() bool {
	if o != nil && o.Mixed != nil {
		return true
	}

	return false
}

// SetMixed gets a reference to the given float32 and assigns it to the Mixed field.
func (o *LicenseDetailsUsagePolicies) SetMixed(v float32) {
	o.Mixed = &v
}

func (o LicenseDetailsUsagePolicies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Stop != nil {
		toSerialize["stop"] = o.Stop
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Dns != nil {
		toSerialize["dns"] = o.Dns
	}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.Mixed != nil {
		toSerialize["mixed"] = o.Mixed
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseDetailsUsagePolicies struct {
	value *LicenseDetailsUsagePolicies
	isSet bool
}

func (v NullableLicenseDetailsUsagePolicies) Get() *LicenseDetailsUsagePolicies {
	return v.value
}

func (v *NullableLicenseDetailsUsagePolicies) Set(val *LicenseDetailsUsagePolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseDetailsUsagePolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseDetailsUsagePolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseDetailsUsagePolicies(val *LicenseDetailsUsagePolicies) *NullableLicenseDetailsUsagePolicies {
	return &NullableLicenseDetailsUsagePolicies{value: val, isSet: true}
}

func (v NullableLicenseDetailsUsagePolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseDetailsUsagePolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
