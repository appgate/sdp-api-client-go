/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// LicenseAllOf License details.
type LicenseAllOf struct {
	// Unique ID for the license.
	Id *string `json:"id,omitempty"`
	// License version. There can be only one v1 License in the system, while v2 Licenses can be added as extra.
	Version *float32 `json:"version,omitempty"`
	// Type of the license. 1: production, 2: installation, 3: test, 4: built-in, 5: aws built-in, 6: metered
	Type *float32 `json:"type,omitempty"`
	// Request code for the license. If built-in license is in place, use this code to get a license. It's based on the CA certificate.
	Request *string `json:"request,omitempty"`
	// The expiration date of the license.
	Expiration *time.Time `json:"expiration,omitempty"`
	// Error message if the License is not invalid.
	Error *string `json:"error,omitempty"`
}

// NewLicenseAllOf instantiates a new LicenseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseAllOf() *LicenseAllOf {
	this := LicenseAllOf{}
	return &this
}

// NewLicenseAllOfWithDefaults instantiates a new LicenseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseAllOfWithDefaults() *LicenseAllOf {
	this := LicenseAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LicenseAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LicenseAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LicenseAllOf) SetId(v string) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LicenseAllOf) GetVersion() float32 {
	if o == nil || o.Version == nil {
		var ret float32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAllOf) GetVersionOk() (*float32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LicenseAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float32 and assigns it to the Version field.
func (o *LicenseAllOf) SetVersion(v float32) {
	o.Version = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LicenseAllOf) GetType() float32 {
	if o == nil || o.Type == nil {
		var ret float32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAllOf) GetTypeOk() (*float32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LicenseAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given float32 and assigns it to the Type field.
func (o *LicenseAllOf) SetType(v float32) {
	o.Type = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *LicenseAllOf) GetRequest() string {
	if o == nil || o.Request == nil {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAllOf) GetRequestOk() (*string, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *LicenseAllOf) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *LicenseAllOf) SetRequest(v string) {
	o.Request = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *LicenseAllOf) GetExpiration() time.Time {
	if o == nil || o.Expiration == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAllOf) GetExpirationOk() (*time.Time, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *LicenseAllOf) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *LicenseAllOf) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *LicenseAllOf) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAllOf) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *LicenseAllOf) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *LicenseAllOf) SetError(v string) {
	o.Error = &v
}

func (o LicenseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseAllOf struct {
	value *LicenseAllOf
	isSet bool
}

func (v NullableLicenseAllOf) Get() *LicenseAllOf {
	return v.value
}

func (v *NullableLicenseAllOf) Set(val *LicenseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseAllOf(val *LicenseAllOf) *NullableLicenseAllOf {
	return &NullableLicenseAllOf{value: val, isSet: true}
}

func (v NullableLicenseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
