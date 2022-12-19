/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ClaimNamesList struct for ClaimNamesList
type ClaimNamesList struct {
	User     []ClaimNamesInner `json:"user,omitempty"`
	Device   []ClaimNamesInner `json:"device,omitempty"`
	System   []ClaimNamesInner `json:"system,omitempty"`
	OnDemand []ClaimNamesInner `json:"onDemand,omitempty"`
	// Risk Rules available via ZTP integration.
	ZtpRiskRules []ClaimNamesListZtpRiskRulesInner `json:"ztpRiskRules,omitempty"`
}

// NewClaimNamesList instantiates a new ClaimNamesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimNamesList() *ClaimNamesList {
	this := ClaimNamesList{}
	return &this
}

// NewClaimNamesListWithDefaults instantiates a new ClaimNamesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimNamesListWithDefaults() *ClaimNamesList {
	this := ClaimNamesList{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ClaimNamesList) GetUser() []ClaimNamesInner {
	if o == nil || o.User == nil {
		var ret []ClaimNamesInner
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimNamesList) GetUserOk() ([]ClaimNamesInner, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ClaimNamesList) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given []ClaimNamesInner and assigns it to the User field.
func (o *ClaimNamesList) SetUser(v []ClaimNamesInner) {
	o.User = v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *ClaimNamesList) GetDevice() []ClaimNamesInner {
	if o == nil || o.Device == nil {
		var ret []ClaimNamesInner
		return ret
	}
	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimNamesList) GetDeviceOk() ([]ClaimNamesInner, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *ClaimNamesList) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given []ClaimNamesInner and assigns it to the Device field.
func (o *ClaimNamesList) SetDevice(v []ClaimNamesInner) {
	o.Device = v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *ClaimNamesList) GetSystem() []ClaimNamesInner {
	if o == nil || o.System == nil {
		var ret []ClaimNamesInner
		return ret
	}
	return o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimNamesList) GetSystemOk() ([]ClaimNamesInner, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *ClaimNamesList) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given []ClaimNamesInner and assigns it to the System field.
func (o *ClaimNamesList) SetSystem(v []ClaimNamesInner) {
	o.System = v
}

// GetOnDemand returns the OnDemand field value if set, zero value otherwise.
func (o *ClaimNamesList) GetOnDemand() []ClaimNamesInner {
	if o == nil || o.OnDemand == nil {
		var ret []ClaimNamesInner
		return ret
	}
	return o.OnDemand
}

// GetOnDemandOk returns a tuple with the OnDemand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimNamesList) GetOnDemandOk() ([]ClaimNamesInner, bool) {
	if o == nil || o.OnDemand == nil {
		return nil, false
	}
	return o.OnDemand, true
}

// HasOnDemand returns a boolean if a field has been set.
func (o *ClaimNamesList) HasOnDemand() bool {
	if o != nil && o.OnDemand != nil {
		return true
	}

	return false
}

// SetOnDemand gets a reference to the given []ClaimNamesInner and assigns it to the OnDemand field.
func (o *ClaimNamesList) SetOnDemand(v []ClaimNamesInner) {
	o.OnDemand = v
}

// GetZtpRiskRules returns the ZtpRiskRules field value if set, zero value otherwise.
func (o *ClaimNamesList) GetZtpRiskRules() []ClaimNamesListZtpRiskRulesInner {
	if o == nil || o.ZtpRiskRules == nil {
		var ret []ClaimNamesListZtpRiskRulesInner
		return ret
	}
	return o.ZtpRiskRules
}

// GetZtpRiskRulesOk returns a tuple with the ZtpRiskRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimNamesList) GetZtpRiskRulesOk() ([]ClaimNamesListZtpRiskRulesInner, bool) {
	if o == nil || o.ZtpRiskRules == nil {
		return nil, false
	}
	return o.ZtpRiskRules, true
}

// HasZtpRiskRules returns a boolean if a field has been set.
func (o *ClaimNamesList) HasZtpRiskRules() bool {
	if o != nil && o.ZtpRiskRules != nil {
		return true
	}

	return false
}

// SetZtpRiskRules gets a reference to the given []ClaimNamesListZtpRiskRulesInner and assigns it to the ZtpRiskRules field.
func (o *ClaimNamesList) SetZtpRiskRules(v []ClaimNamesListZtpRiskRulesInner) {
	o.ZtpRiskRules = v
}

func (o ClaimNamesList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.System != nil {
		toSerialize["system"] = o.System
	}
	if o.OnDemand != nil {
		toSerialize["onDemand"] = o.OnDemand
	}
	if o.ZtpRiskRules != nil {
		toSerialize["ztpRiskRules"] = o.ZtpRiskRules
	}
	return json.Marshal(toSerialize)
}

type NullableClaimNamesList struct {
	value *ClaimNamesList
	isSet bool
}

func (v NullableClaimNamesList) Get() *ClaimNamesList {
	return v.value
}

func (v *NullableClaimNamesList) Set(val *ClaimNamesList) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimNamesList) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimNamesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimNamesList(val *ClaimNamesList) *NullableClaimNamesList {
	return &NullableClaimNamesList{value: val, isSet: true}
}

func (v NullableClaimNamesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimNamesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
