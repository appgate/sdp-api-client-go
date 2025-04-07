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
	"time"
)

// LoginResponse struct for LoginResponse
type LoginResponse struct {
	User *LoginResponseUser `json:"user,omitempty"`
	// The AuthToken required for subsequent API calls.
	Token *string `json:"token,omitempty"`
	// Token expiration time.
	Expires *time.Time `json:"expires,omitempty"`
	// Message of the day configured by an admin.
	MessageOfTheDay *string `json:"messageOfTheDay,omitempty"`
	// ZTP type of the collective.
	ZtpCollectiveType *string `json:"ztpCollectiveType,omitempty"`
	// ZTP account type.
	ZtpAccountType *string `json:"ztpAccountType,omitempty"`
	// Whether X509 CRL is enabled for the system or not. Issued Certificates is disabled if it's not enabled.
	CrlEnabled *bool `json:"crlEnabled,omitempty"`
}

// NewLoginResponse instantiates a new LoginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginResponse() *LoginResponse {
	this := LoginResponse{}
	return &this
}

// NewLoginResponseWithDefaults instantiates a new LoginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginResponseWithDefaults() *LoginResponse {
	this := LoginResponse{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LoginResponse) GetUser() LoginResponseUser {
	if o == nil || o.User == nil {
		var ret LoginResponseUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetUserOk() (*LoginResponseUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LoginResponse) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given LoginResponseUser and assigns it to the User field.
func (o *LoginResponse) SetUser(v LoginResponseUser) {
	o.User = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LoginResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LoginResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LoginResponse) SetToken(v string) {
	o.Token = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *LoginResponse) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *LoginResponse) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *LoginResponse) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetMessageOfTheDay returns the MessageOfTheDay field value if set, zero value otherwise.
func (o *LoginResponse) GetMessageOfTheDay() string {
	if o == nil || o.MessageOfTheDay == nil {
		var ret string
		return ret
	}
	return *o.MessageOfTheDay
}

// GetMessageOfTheDayOk returns a tuple with the MessageOfTheDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetMessageOfTheDayOk() (*string, bool) {
	if o == nil || o.MessageOfTheDay == nil {
		return nil, false
	}
	return o.MessageOfTheDay, true
}

// HasMessageOfTheDay returns a boolean if a field has been set.
func (o *LoginResponse) HasMessageOfTheDay() bool {
	if o != nil && o.MessageOfTheDay != nil {
		return true
	}

	return false
}

// SetMessageOfTheDay gets a reference to the given string and assigns it to the MessageOfTheDay field.
func (o *LoginResponse) SetMessageOfTheDay(v string) {
	o.MessageOfTheDay = &v
}

// GetZtpCollectiveType returns the ZtpCollectiveType field value if set, zero value otherwise.
func (o *LoginResponse) GetZtpCollectiveType() string {
	if o == nil || o.ZtpCollectiveType == nil {
		var ret string
		return ret
	}
	return *o.ZtpCollectiveType
}

// GetZtpCollectiveTypeOk returns a tuple with the ZtpCollectiveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetZtpCollectiveTypeOk() (*string, bool) {
	if o == nil || o.ZtpCollectiveType == nil {
		return nil, false
	}
	return o.ZtpCollectiveType, true
}

// HasZtpCollectiveType returns a boolean if a field has been set.
func (o *LoginResponse) HasZtpCollectiveType() bool {
	if o != nil && o.ZtpCollectiveType != nil {
		return true
	}

	return false
}

// SetZtpCollectiveType gets a reference to the given string and assigns it to the ZtpCollectiveType field.
func (o *LoginResponse) SetZtpCollectiveType(v string) {
	o.ZtpCollectiveType = &v
}

// GetZtpAccountType returns the ZtpAccountType field value if set, zero value otherwise.
func (o *LoginResponse) GetZtpAccountType() string {
	if o == nil || o.ZtpAccountType == nil {
		var ret string
		return ret
	}
	return *o.ZtpAccountType
}

// GetZtpAccountTypeOk returns a tuple with the ZtpAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetZtpAccountTypeOk() (*string, bool) {
	if o == nil || o.ZtpAccountType == nil {
		return nil, false
	}
	return o.ZtpAccountType, true
}

// HasZtpAccountType returns a boolean if a field has been set.
func (o *LoginResponse) HasZtpAccountType() bool {
	if o != nil && o.ZtpAccountType != nil {
		return true
	}

	return false
}

// SetZtpAccountType gets a reference to the given string and assigns it to the ZtpAccountType field.
func (o *LoginResponse) SetZtpAccountType(v string) {
	o.ZtpAccountType = &v
}

// GetCrlEnabled returns the CrlEnabled field value if set, zero value otherwise.
func (o *LoginResponse) GetCrlEnabled() bool {
	if o == nil || o.CrlEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CrlEnabled
}

// GetCrlEnabledOk returns a tuple with the CrlEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetCrlEnabledOk() (*bool, bool) {
	if o == nil || o.CrlEnabled == nil {
		return nil, false
	}
	return o.CrlEnabled, true
}

// HasCrlEnabled returns a boolean if a field has been set.
func (o *LoginResponse) HasCrlEnabled() bool {
	if o != nil && o.CrlEnabled != nil {
		return true
	}

	return false
}

// SetCrlEnabled gets a reference to the given bool and assigns it to the CrlEnabled field.
func (o *LoginResponse) SetCrlEnabled(v bool) {
	o.CrlEnabled = &v
}

func (o LoginResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.MessageOfTheDay != nil {
		toSerialize["messageOfTheDay"] = o.MessageOfTheDay
	}
	if o.ZtpCollectiveType != nil {
		toSerialize["ztpCollectiveType"] = o.ZtpCollectiveType
	}
	if o.ZtpAccountType != nil {
		toSerialize["ztpAccountType"] = o.ZtpAccountType
	}
	if o.CrlEnabled != nil {
		toSerialize["crlEnabled"] = o.CrlEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableLoginResponse struct {
	value *LoginResponse
	isSet bool
}

func (v NullableLoginResponse) Get() *LoginResponse {
	return v.value
}

func (v *NullableLoginResponse) Set(val *LoginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginResponse(val *LoginResponse) *NullableLoginResponse {
	return &NullableLoginResponse{value: val, isSet: true}
}

func (v NullableLoginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
