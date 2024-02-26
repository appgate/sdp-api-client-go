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
)

// Portal12 struct for Portal12
type Portal12 struct {
	// Identifier to track the object on update since all the other fields are write-only. A random one will be assigned if left empty.
	Id *string `json:"id,omitempty"`
	// Contents of the P12 file in Base64 format.
	Content *string `json:"content,omitempty"`
	// Password for the P12 file.
	Password *string `json:"password,omitempty"`
	// Subject name of the certificate in the file.
	SubjectName *string `json:"subjectName,omitempty"`
	// Portal will verify upstream certificate of the endpoints.
	VerifyUpstream *bool `json:"verifyUpstream,omitempty"`
}

// NewPortal12 instantiates a new Portal12 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortal12() *Portal12 {
	this := Portal12{}
	var verifyUpstream bool = true
	this.VerifyUpstream = &verifyUpstream
	return &this
}

// NewPortal12WithDefaults instantiates a new Portal12 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortal12WithDefaults() *Portal12 {
	this := Portal12{}
	var verifyUpstream bool = true
	this.VerifyUpstream = &verifyUpstream
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Portal12) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal12) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Portal12) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Portal12) SetId(v string) {
	o.Id = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Portal12) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal12) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Portal12) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *Portal12) SetContent(v string) {
	o.Content = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Portal12) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal12) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Portal12) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *Portal12) SetPassword(v string) {
	o.Password = &v
}

// GetSubjectName returns the SubjectName field value if set, zero value otherwise.
func (o *Portal12) GetSubjectName() string {
	if o == nil || o.SubjectName == nil {
		var ret string
		return ret
	}
	return *o.SubjectName
}

// GetSubjectNameOk returns a tuple with the SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal12) GetSubjectNameOk() (*string, bool) {
	if o == nil || o.SubjectName == nil {
		return nil, false
	}
	return o.SubjectName, true
}

// HasSubjectName returns a boolean if a field has been set.
func (o *Portal12) HasSubjectName() bool {
	if o != nil && o.SubjectName != nil {
		return true
	}

	return false
}

// SetSubjectName gets a reference to the given string and assigns it to the SubjectName field.
func (o *Portal12) SetSubjectName(v string) {
	o.SubjectName = &v
}

// GetVerifyUpstream returns the VerifyUpstream field value if set, zero value otherwise.
func (o *Portal12) GetVerifyUpstream() bool {
	if o == nil || o.VerifyUpstream == nil {
		var ret bool
		return ret
	}
	return *o.VerifyUpstream
}

// GetVerifyUpstreamOk returns a tuple with the VerifyUpstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Portal12) GetVerifyUpstreamOk() (*bool, bool) {
	if o == nil || o.VerifyUpstream == nil {
		return nil, false
	}
	return o.VerifyUpstream, true
}

// HasVerifyUpstream returns a boolean if a field has been set.
func (o *Portal12) HasVerifyUpstream() bool {
	if o != nil && o.VerifyUpstream != nil {
		return true
	}

	return false
}

// SetVerifyUpstream gets a reference to the given bool and assigns it to the VerifyUpstream field.
func (o *Portal12) SetVerifyUpstream(v bool) {
	o.VerifyUpstream = &v
}

func (o Portal12) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.SubjectName != nil {
		toSerialize["subjectName"] = o.SubjectName
	}
	if o.VerifyUpstream != nil {
		toSerialize["verifyUpstream"] = o.VerifyUpstream
	}
	return json.Marshal(toSerialize)
}

type NullablePortal12 struct {
	value *Portal12
	isSet bool
}

func (v NullablePortal12) Get() *Portal12 {
	return v.value
}

func (v *NullablePortal12) Set(val *Portal12) {
	v.value = val
	v.isSet = true
}

func (v NullablePortal12) IsSet() bool {
	return v.isSet
}

func (v *NullablePortal12) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortal12(val *Portal12) *NullablePortal12 {
	return &NullablePortal12{value: val, isSet: true}
}

func (v NullablePortal12) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortal12) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
