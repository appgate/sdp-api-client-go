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

// IdentityProvidersNamesGet200Response struct for IdentityProvidersNamesGet200Response
type IdentityProvidersNamesGet200Response struct {
	// The list of the identity providers.
	Data []IdentityProvidersNamesGet200ResponseDataInner `json:"data,omitempty"`
	// Configured banner message to display before login.
	BannerMessage *string `json:"bannerMessage,omitempty"`
}

// NewIdentityProvidersNamesGet200Response instantiates a new IdentityProvidersNamesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProvidersNamesGet200Response() *IdentityProvidersNamesGet200Response {
	this := IdentityProvidersNamesGet200Response{}
	return &this
}

// NewIdentityProvidersNamesGet200ResponseWithDefaults instantiates a new IdentityProvidersNamesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProvidersNamesGet200ResponseWithDefaults() *IdentityProvidersNamesGet200Response {
	this := IdentityProvidersNamesGet200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IdentityProvidersNamesGet200Response) GetData() []IdentityProvidersNamesGet200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []IdentityProvidersNamesGet200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersNamesGet200Response) GetDataOk() ([]IdentityProvidersNamesGet200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IdentityProvidersNamesGet200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []IdentityProvidersNamesGet200ResponseDataInner and assigns it to the Data field.
func (o *IdentityProvidersNamesGet200Response) SetData(v []IdentityProvidersNamesGet200ResponseDataInner) {
	o.Data = v
}

// GetBannerMessage returns the BannerMessage field value if set, zero value otherwise.
func (o *IdentityProvidersNamesGet200Response) GetBannerMessage() string {
	if o == nil || o.BannerMessage == nil {
		var ret string
		return ret
	}
	return *o.BannerMessage
}

// GetBannerMessageOk returns a tuple with the BannerMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersNamesGet200Response) GetBannerMessageOk() (*string, bool) {
	if o == nil || o.BannerMessage == nil {
		return nil, false
	}
	return o.BannerMessage, true
}

// HasBannerMessage returns a boolean if a field has been set.
func (o *IdentityProvidersNamesGet200Response) HasBannerMessage() bool {
	if o != nil && o.BannerMessage != nil {
		return true
	}

	return false
}

// SetBannerMessage gets a reference to the given string and assigns it to the BannerMessage field.
func (o *IdentityProvidersNamesGet200Response) SetBannerMessage(v string) {
	o.BannerMessage = &v
}

func (o IdentityProvidersNamesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.BannerMessage != nil {
		toSerialize["bannerMessage"] = o.BannerMessage
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityProvidersNamesGet200Response struct {
	value *IdentityProvidersNamesGet200Response
	isSet bool
}

func (v NullableIdentityProvidersNamesGet200Response) Get() *IdentityProvidersNamesGet200Response {
	return v.value
}

func (v *NullableIdentityProvidersNamesGet200Response) Set(val *IdentityProvidersNamesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProvidersNamesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProvidersNamesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProvidersNamesGet200Response(val *IdentityProvidersNamesGet200Response) *NullableIdentityProvidersNamesGet200Response {
	return &NullableIdentityProvidersNamesGet200Response{value: val, isSet: true}
}

func (v NullableIdentityProvidersNamesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProvidersNamesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
