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

// AppDiscoveryAllOfData A Discovered App.
type AppDiscoveryAllOfData struct {
	// The app name, with the domain field it makes up the discovered hostname.
	App *string `json:"app,omitempty"`
	// The domain name as configured.
	Domain *string `json:"domain,omitempty"`
	// Number of unique users have attempted to access this app.
	AccessCount *float32 `json:"accessCount,omitempty"`
	// The resolved IP addresses for this app.
	Ips []string `json:"ips,omitempty"`
	// A list of distinguished names that tried to access this app.
	DistinguishedNames []AppDiscoveryAllOfDistinguishedNames `json:"distinguishedNames,omitempty"`
}

// NewAppDiscoveryAllOfData instantiates a new AppDiscoveryAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDiscoveryAllOfData() *AppDiscoveryAllOfData {
	this := AppDiscoveryAllOfData{}
	return &this
}

// NewAppDiscoveryAllOfDataWithDefaults instantiates a new AppDiscoveryAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDiscoveryAllOfDataWithDefaults() *AppDiscoveryAllOfData {
	this := AppDiscoveryAllOfData{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *AppDiscoveryAllOfData) GetApp() string {
	if o == nil || o.App == nil {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDiscoveryAllOfData) GetAppOk() (*string, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *AppDiscoveryAllOfData) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *AppDiscoveryAllOfData) SetApp(v string) {
	o.App = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *AppDiscoveryAllOfData) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDiscoveryAllOfData) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *AppDiscoveryAllOfData) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *AppDiscoveryAllOfData) SetDomain(v string) {
	o.Domain = &v
}

// GetAccessCount returns the AccessCount field value if set, zero value otherwise.
func (o *AppDiscoveryAllOfData) GetAccessCount() float32 {
	if o == nil || o.AccessCount == nil {
		var ret float32
		return ret
	}
	return *o.AccessCount
}

// GetAccessCountOk returns a tuple with the AccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDiscoveryAllOfData) GetAccessCountOk() (*float32, bool) {
	if o == nil || o.AccessCount == nil {
		return nil, false
	}
	return o.AccessCount, true
}

// HasAccessCount returns a boolean if a field has been set.
func (o *AppDiscoveryAllOfData) HasAccessCount() bool {
	if o != nil && o.AccessCount != nil {
		return true
	}

	return false
}

// SetAccessCount gets a reference to the given float32 and assigns it to the AccessCount field.
func (o *AppDiscoveryAllOfData) SetAccessCount(v float32) {
	o.AccessCount = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *AppDiscoveryAllOfData) GetIps() []string {
	if o == nil || o.Ips == nil {
		var ret []string
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDiscoveryAllOfData) GetIpsOk() ([]string, bool) {
	if o == nil || o.Ips == nil {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *AppDiscoveryAllOfData) HasIps() bool {
	if o != nil && o.Ips != nil {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *AppDiscoveryAllOfData) SetIps(v []string) {
	o.Ips = v
}

// GetDistinguishedNames returns the DistinguishedNames field value if set, zero value otherwise.
func (o *AppDiscoveryAllOfData) GetDistinguishedNames() []AppDiscoveryAllOfDistinguishedNames {
	if o == nil || o.DistinguishedNames == nil {
		var ret []AppDiscoveryAllOfDistinguishedNames
		return ret
	}
	return o.DistinguishedNames
}

// GetDistinguishedNamesOk returns a tuple with the DistinguishedNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDiscoveryAllOfData) GetDistinguishedNamesOk() ([]AppDiscoveryAllOfDistinguishedNames, bool) {
	if o == nil || o.DistinguishedNames == nil {
		return nil, false
	}
	return o.DistinguishedNames, true
}

// HasDistinguishedNames returns a boolean if a field has been set.
func (o *AppDiscoveryAllOfData) HasDistinguishedNames() bool {
	if o != nil && o.DistinguishedNames != nil {
		return true
	}

	return false
}

// SetDistinguishedNames gets a reference to the given []AppDiscoveryAllOfDistinguishedNames and assigns it to the DistinguishedNames field.
func (o *AppDiscoveryAllOfData) SetDistinguishedNames(v []AppDiscoveryAllOfDistinguishedNames) {
	o.DistinguishedNames = v
}

func (o AppDiscoveryAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.AccessCount != nil {
		toSerialize["accessCount"] = o.AccessCount
	}
	if o.Ips != nil {
		toSerialize["ips"] = o.Ips
	}
	if o.DistinguishedNames != nil {
		toSerialize["distinguishedNames"] = o.DistinguishedNames
	}
	return json.Marshal(toSerialize)
}

type NullableAppDiscoveryAllOfData struct {
	value *AppDiscoveryAllOfData
	isSet bool
}

func (v NullableAppDiscoveryAllOfData) Get() *AppDiscoveryAllOfData {
	return v.value
}

func (v *NullableAppDiscoveryAllOfData) Set(val *AppDiscoveryAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDiscoveryAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDiscoveryAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDiscoveryAllOfData(val *AppDiscoveryAllOfData) *NullableAppDiscoveryAllOfData {
	return &NullableAppDiscoveryAllOfData{value: val, isSet: true}
}

func (v NullableAppDiscoveryAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDiscoveryAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
