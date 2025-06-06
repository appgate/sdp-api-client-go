/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LicenseDetailsUsageClientTypes Counts of the current client types in the system at present.
type LicenseDetailsUsageClientTypes struct {
	// The amount of full client types in the system at present.
	Full *float32 `json:"full,omitempty"`
	// The amount of full multi client types in the system at present.
	FullMulti *float32 `json:"full_multi,omitempty"`
	// The amount of full always on client types in the system at present.
	FullAlwaysOn *float32 `json:"full_always_on,omitempty"`
	// The amount of lite client types in the system at present.
	Lite *float32 `json:"lite,omitempty"`
	// The amount of headless client types in the system at present.
	Headless *float32 `json:"headless,omitempty"`
	// The amount of full always on client types in the system at present.
	HeadlessAlwaysOn *float32 `json:"headless_always_on,omitempty"`
	// The amount of connector client types in the system at present.
	Connector *float32 `json:"connector,omitempty"`
	// The amount of sso client types in the system at present.
	Sso *float32 `json:"sso,omitempty"`
	// The amount of web client types in the system at present.
	Web *float32 `json:"web,omitempty"`
	// The amount of unknown client types in the system at present.
	Unknown *float32 `json:"unknown,omitempty"`
}

// NewLicenseDetailsUsageClientTypes instantiates a new LicenseDetailsUsageClientTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseDetailsUsageClientTypes() *LicenseDetailsUsageClientTypes {
	this := LicenseDetailsUsageClientTypes{}
	return &this
}

// NewLicenseDetailsUsageClientTypesWithDefaults instantiates a new LicenseDetailsUsageClientTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseDetailsUsageClientTypesWithDefaults() *LicenseDetailsUsageClientTypes {
	this := LicenseDetailsUsageClientTypes{}
	return &this
}

// GetFull returns the Full field value if set, zero value otherwise.
func (o *LicenseDetailsUsageClientTypes) GetFull() float32 {
	if o == nil || o.Full == nil {
		var ret float32
		return ret
	}
	return *o.Full
}

// GetFullOk returns a tuple with the Full field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageClientTypes) GetFullOk() (*float32, bool) {
	if o == nil || o.Full == nil {
		return nil, false
	}
	return o.Full, true
}

// HasFull returns a boolean if a field has been set.
func (o *LicenseDetailsUsageClientTypes) HasFull() bool {
	if o != nil && o.Full != nil {
		return true
	}

	return false
}

// SetFull gets a reference to the given float32 and assigns it to the Full field.
func (o *LicenseDetailsUsageClientTypes) SetFull(v float32) {
	o.Full = &v
}

// GetFullMulti returns the FullMulti field value if set, zero value otherwise.
func (o *LicenseDetailsUsageClientTypes) GetFullMulti() float32 {
	if o == nil || o.FullMulti == nil {
		var ret float32
		return ret
	}
	return *o.FullMulti
}

// GetFullMultiOk returns a tuple with the FullMulti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageClientTypes) GetFullMultiOk() (*float32, bool) {
	if o == nil || o.FullMulti == nil {
		return nil, false
	}
	return o.FullMulti, true
}

// HasFullMulti returns a boolean if a field has been set.
func (o *LicenseDetailsUsageClientTypes) HasFullMulti() bool {
	if o != nil && o.FullMulti != nil {
		return true
	}

	return false
}

// SetFullMulti gets a reference to the given float32 and assigns it to the FullMulti field.
func (o *LicenseDetailsUsageClientTypes) SetFullMulti(v float32) {
	o.FullMulti = &v
}

// GetFullAlwaysOn returns the FullAlwaysOn field value if set, zero value otherwise.
func (o *LicenseDetailsUsageClientTypes) GetFullAlwaysOn() float32 {
	if o == nil || o.FullAlwaysOn == nil {
		var ret float32
		return ret
	}
	return *o.FullAlwaysOn
}

// GetFullAlwaysOnOk returns a tuple with the FullAlwaysOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageClientTypes) GetFullAlwaysOnOk() (*float32, bool) {
	if o == nil || o.FullAlwaysOn == nil {
		return nil, false
	}
	return o.FullAlwaysOn, true
}

// HasFullAlwaysOn returns a boolean if a field has been set.
func (o *LicenseDetailsUsageClientTypes) HasFullAlwaysOn() bool {
	if o != nil && o.FullAlwaysOn != nil {
		return true
	}

	return false
}

// SetFullAlwaysOn gets a reference to the given float32 and assigns it to the FullAlwaysOn field.
func (o *LicenseDetailsUsageClientTypes) SetFullAlwaysOn(v float32) {
	o.FullAlwaysOn = &v
}

// GetLite returns the Lite field value if set, zero value otherwise.
func (o *LicenseDetailsUsageClientTypes) GetLite() float32 {
	if o == nil || o.Lite == nil {
		var ret float32
		return ret
	}
	return *o.Lite
}

// GetLiteOk returns a tuple with the Lite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageClientTypes) GetLiteOk() (*float32, bool) {
	if o == nil || o.Lite == nil {
		return nil, false
	}
	return o.Lite, true
}

// HasLite returns a boolean if a field has been set.
func (o *LicenseDetailsUsageClientTypes) HasLite() bool {
	if o != nil && o.Lite != nil {
		return true
	}

	return false
}

// SetLite gets a reference to the given float32 and assigns it to the Lite field.
func (o *LicenseDetailsUsageClientTypes) SetLite(v float32) {
	o.Lite = &v
}

// GetHeadless returns the Headless field value if set, zero value otherwise.
func (o *LicenseDetailsUsageClientTypes) GetHeadless() float32 {
	if o == nil || o.Headless == nil {
		var ret float32
		return ret
	}
	return *o.Headless
}

// GetHeadlessOk returns a tuple with the Headless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageClientTypes) GetHeadlessOk() (*float32, bool) {
	if o == nil || o.Headless == nil {
		return nil, false
	}
	return o.Headless, true
}

// HasHeadless returns a boolean if a field has been set.
func (o *LicenseDetailsUsageClientTypes) HasHeadless() bool {
	if o != nil && o.Headless != nil {
		return true
	}

	return false
}

// SetHeadless gets a reference to the given float32 and assigns it to the Headless field.
func (o *LicenseDetailsUsageClientTypes) SetHeadless(v float32) {
	o.Headless = &v
}

// GetHeadlessAlwaysOn returns the HeadlessAlwaysOn field value if set, zero value otherwise.
func (o *LicenseDetailsUsageClientTypes) GetHeadlessAlwaysOn() float32 {
	if o == nil || o.HeadlessAlwaysOn == nil {
		var ret float32
		return ret
	}
	return *o.HeadlessAlwaysOn
}

// GetHeadlessAlwaysOnOk returns a tuple with the HeadlessAlwaysOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageClientTypes) GetHeadlessAlwaysOnOk() (*float32, bool) {
	if o == nil || o.HeadlessAlwaysOn == nil {
		return nil, false
	}
	return o.HeadlessAlwaysOn, true
}

// HasHeadlessAlwaysOn returns a boolean if a field has been set.
func (o *LicenseDetailsUsageClientTypes) HasHeadlessAlwaysOn() bool {
	if o != nil && o.HeadlessAlwaysOn != nil {
		return true
	}

	return false
}

// SetHeadlessAlwaysOn gets a reference to the given float32 and assigns it to the HeadlessAlwaysOn field.
func (o *LicenseDetailsUsageClientTypes) SetHeadlessAlwaysOn(v float32) {
	o.HeadlessAlwaysOn = &v
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *LicenseDetailsUsageClientTypes) GetConnector() float32 {
	if o == nil || o.Connector == nil {
		var ret float32
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageClientTypes) GetConnectorOk() (*float32, bool) {
	if o == nil || o.Connector == nil {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *LicenseDetailsUsageClientTypes) HasConnector() bool {
	if o != nil && o.Connector != nil {
		return true
	}

	return false
}

// SetConnector gets a reference to the given float32 and assigns it to the Connector field.
func (o *LicenseDetailsUsageClientTypes) SetConnector(v float32) {
	o.Connector = &v
}

// GetSso returns the Sso field value if set, zero value otherwise.
func (o *LicenseDetailsUsageClientTypes) GetSso() float32 {
	if o == nil || o.Sso == nil {
		var ret float32
		return ret
	}
	return *o.Sso
}

// GetSsoOk returns a tuple with the Sso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageClientTypes) GetSsoOk() (*float32, bool) {
	if o == nil || o.Sso == nil {
		return nil, false
	}
	return o.Sso, true
}

// HasSso returns a boolean if a field has been set.
func (o *LicenseDetailsUsageClientTypes) HasSso() bool {
	if o != nil && o.Sso != nil {
		return true
	}

	return false
}

// SetSso gets a reference to the given float32 and assigns it to the Sso field.
func (o *LicenseDetailsUsageClientTypes) SetSso(v float32) {
	o.Sso = &v
}

// GetWeb returns the Web field value if set, zero value otherwise.
func (o *LicenseDetailsUsageClientTypes) GetWeb() float32 {
	if o == nil || o.Web == nil {
		var ret float32
		return ret
	}
	return *o.Web
}

// GetWebOk returns a tuple with the Web field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageClientTypes) GetWebOk() (*float32, bool) {
	if o == nil || o.Web == nil {
		return nil, false
	}
	return o.Web, true
}

// HasWeb returns a boolean if a field has been set.
func (o *LicenseDetailsUsageClientTypes) HasWeb() bool {
	if o != nil && o.Web != nil {
		return true
	}

	return false
}

// SetWeb gets a reference to the given float32 and assigns it to the Web field.
func (o *LicenseDetailsUsageClientTypes) SetWeb(v float32) {
	o.Web = &v
}

// GetUnknown returns the Unknown field value if set, zero value otherwise.
func (o *LicenseDetailsUsageClientTypes) GetUnknown() float32 {
	if o == nil || o.Unknown == nil {
		var ret float32
		return ret
	}
	return *o.Unknown
}

// GetUnknownOk returns a tuple with the Unknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetailsUsageClientTypes) GetUnknownOk() (*float32, bool) {
	if o == nil || o.Unknown == nil {
		return nil, false
	}
	return o.Unknown, true
}

// HasUnknown returns a boolean if a field has been set.
func (o *LicenseDetailsUsageClientTypes) HasUnknown() bool {
	if o != nil && o.Unknown != nil {
		return true
	}

	return false
}

// SetUnknown gets a reference to the given float32 and assigns it to the Unknown field.
func (o *LicenseDetailsUsageClientTypes) SetUnknown(v float32) {
	o.Unknown = &v
}

func (o LicenseDetailsUsageClientTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Full != nil {
		toSerialize["full"] = o.Full
	}
	if o.FullMulti != nil {
		toSerialize["full_multi"] = o.FullMulti
	}
	if o.FullAlwaysOn != nil {
		toSerialize["full_always_on"] = o.FullAlwaysOn
	}
	if o.Lite != nil {
		toSerialize["lite"] = o.Lite
	}
	if o.Headless != nil {
		toSerialize["headless"] = o.Headless
	}
	if o.HeadlessAlwaysOn != nil {
		toSerialize["headless_always_on"] = o.HeadlessAlwaysOn
	}
	if o.Connector != nil {
		toSerialize["connector"] = o.Connector
	}
	if o.Sso != nil {
		toSerialize["sso"] = o.Sso
	}
	if o.Web != nil {
		toSerialize["web"] = o.Web
	}
	if o.Unknown != nil {
		toSerialize["unknown"] = o.Unknown
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseDetailsUsageClientTypes struct {
	value *LicenseDetailsUsageClientTypes
	isSet bool
}

func (v NullableLicenseDetailsUsageClientTypes) Get() *LicenseDetailsUsageClientTypes {
	return v.value
}

func (v *NullableLicenseDetailsUsageClientTypes) Set(val *LicenseDetailsUsageClientTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseDetailsUsageClientTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseDetailsUsageClientTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseDetailsUsageClientTypes(val *LicenseDetailsUsageClientTypes) *NullableLicenseDetailsUsageClientTypes {
	return &NullableLicenseDetailsUsageClientTypes{value: val, isSet: true}
}

func (v NullableLicenseDetailsUsageClientTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseDetailsUsageClientTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
