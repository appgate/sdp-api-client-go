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

// LicenseDetails License details and current usage.
type LicenseDetails struct {
	Licenses []License           `json:"licenses,omitempty"`
	Entitled *LicenseEntitlement `json:"entitled,omitempty"`
	// Request code for the license. Use this code to get a license. It's based on the CA certificate.
	RequestCode *string              `json:"requestCode,omitempty"`
	Usage       *LicenseDetailsUsage `json:"usage,omitempty"`
}

// NewLicenseDetails instantiates a new LicenseDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseDetails() *LicenseDetails {
	this := LicenseDetails{}
	return &this
}

// NewLicenseDetailsWithDefaults instantiates a new LicenseDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseDetailsWithDefaults() *LicenseDetails {
	this := LicenseDetails{}
	return &this
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *LicenseDetails) GetLicenses() []License {
	if o == nil || o.Licenses == nil {
		var ret []License
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetails) GetLicensesOk() ([]License, bool) {
	if o == nil || o.Licenses == nil {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *LicenseDetails) HasLicenses() bool {
	if o != nil && o.Licenses != nil {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []License and assigns it to the Licenses field.
func (o *LicenseDetails) SetLicenses(v []License) {
	o.Licenses = v
}

// GetEntitled returns the Entitled field value if set, zero value otherwise.
func (o *LicenseDetails) GetEntitled() LicenseEntitlement {
	if o == nil || o.Entitled == nil {
		var ret LicenseEntitlement
		return ret
	}
	return *o.Entitled
}

// GetEntitledOk returns a tuple with the Entitled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetails) GetEntitledOk() (*LicenseEntitlement, bool) {
	if o == nil || o.Entitled == nil {
		return nil, false
	}
	return o.Entitled, true
}

// HasEntitled returns a boolean if a field has been set.
func (o *LicenseDetails) HasEntitled() bool {
	if o != nil && o.Entitled != nil {
		return true
	}

	return false
}

// SetEntitled gets a reference to the given LicenseEntitlement and assigns it to the Entitled field.
func (o *LicenseDetails) SetEntitled(v LicenseEntitlement) {
	o.Entitled = &v
}

// GetRequestCode returns the RequestCode field value if set, zero value otherwise.
func (o *LicenseDetails) GetRequestCode() string {
	if o == nil || o.RequestCode == nil {
		var ret string
		return ret
	}
	return *o.RequestCode
}

// GetRequestCodeOk returns a tuple with the RequestCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetails) GetRequestCodeOk() (*string, bool) {
	if o == nil || o.RequestCode == nil {
		return nil, false
	}
	return o.RequestCode, true
}

// HasRequestCode returns a boolean if a field has been set.
func (o *LicenseDetails) HasRequestCode() bool {
	if o != nil && o.RequestCode != nil {
		return true
	}

	return false
}

// SetRequestCode gets a reference to the given string and assigns it to the RequestCode field.
func (o *LicenseDetails) SetRequestCode(v string) {
	o.RequestCode = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *LicenseDetails) GetUsage() LicenseDetailsUsage {
	if o == nil || o.Usage == nil {
		var ret LicenseDetailsUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetails) GetUsageOk() (*LicenseDetailsUsage, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *LicenseDetails) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given LicenseDetailsUsage and assigns it to the Usage field.
func (o *LicenseDetails) SetUsage(v LicenseDetailsUsage) {
	o.Usage = &v
}

func (o LicenseDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Licenses != nil {
		toSerialize["licenses"] = o.Licenses
	}
	if o.Entitled != nil {
		toSerialize["entitled"] = o.Entitled
	}
	if o.RequestCode != nil {
		toSerialize["requestCode"] = o.RequestCode
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseDetails struct {
	value *LicenseDetails
	isSet bool
}

func (v NullableLicenseDetails) Get() *LicenseDetails {
	return v.value
}

func (v *NullableLicenseDetails) Set(val *LicenseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseDetails(val *LicenseDetails) *NullableLicenseDetails {
	return &NullableLicenseDetails{value: val, isSet: true}
}

func (v NullableLicenseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}