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

// AuthenticationOtpInitializePost200Response struct for AuthenticationOtpInitializePost200Response
type AuthenticationOtpInitializePost200Response struct {
	// The type of the Multi-Factor Authentication. * 'AlreadySeeded': The MFA provider is the built-in Time-based OTP provider and the user has already a seed in the system. OTP is required on the next step. * 'Secret': The MFA provider is the built-in Time-based OTP provider and this is the first time the user is doing an MFA. It includes details about the seed. OTP is required on the next step. * 'Challenge': The MFA provider is a RADIUS provider. It might include a challenge that needs to be sent back. OTP is required on the next step. * 'Push': The MFA provider is a RADIUS provider but the authentication is done externally, such as a  mobile app that prompts. Send a dummy OTP right away to trigger the external authentication.
	Type *string `json:"type,omitempty"`
	// The seed for the built-in Time-based OTP provider. Used when configuring TOTP apps manually. Only available in Secret type.
	Secret *string `json:"secret,omitempty"`
	// A URL for triggering TOTP apps directly and configuring an entry automatically. Only available in Secret type.
	OtpAuthUrl *string `json:"otpAuthUrl,omitempty"`
	// The barcode image in jpg format. Base64 encoded. Only available in Secret type.
	Barcode *string `json:"barcode,omitempty"`
	// A message from the RADIUS MFA provider. Only available in Challenge type.
	ResponseMessage *string `json:"responseMessage,omitempty"`
	// State send by the RADIUS MFA provider as challenge. It needs to be sent back during MFA authentication. Only available in Challenge type.
	State *string `json:"state,omitempty"`
	// How long the Controller wait for RADIUS response. Especially useful for external authentication mechanism. Clients waiting shorter than this timeout will fail with a wrong error. Only available in Challenge and Push type.
	Timeout *float32 `json:"timeout,omitempty"`
	// Whether the RADIUS MFA provider expecting the user password for authentication. If true, the user password needs to be sent as otp instead of a dummy value. Only available in Push type.
	SendPassword *bool `json:"sendPassword,omitempty"`
}

// NewAuthenticationOtpInitializePost200Response instantiates a new AuthenticationOtpInitializePost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationOtpInitializePost200Response() *AuthenticationOtpInitializePost200Response {
	this := AuthenticationOtpInitializePost200Response{}
	return &this
}

// NewAuthenticationOtpInitializePost200ResponseWithDefaults instantiates a new AuthenticationOtpInitializePost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationOtpInitializePost200ResponseWithDefaults() *AuthenticationOtpInitializePost200Response {
	this := AuthenticationOtpInitializePost200Response{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthenticationOtpInitializePost200Response) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationOtpInitializePost200Response) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthenticationOtpInitializePost200Response) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthenticationOtpInitializePost200Response) SetType(v string) {
	o.Type = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AuthenticationOtpInitializePost200Response) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationOtpInitializePost200Response) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AuthenticationOtpInitializePost200Response) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AuthenticationOtpInitializePost200Response) SetSecret(v string) {
	o.Secret = &v
}

// GetOtpAuthUrl returns the OtpAuthUrl field value if set, zero value otherwise.
func (o *AuthenticationOtpInitializePost200Response) GetOtpAuthUrl() string {
	if o == nil || o.OtpAuthUrl == nil {
		var ret string
		return ret
	}
	return *o.OtpAuthUrl
}

// GetOtpAuthUrlOk returns a tuple with the OtpAuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationOtpInitializePost200Response) GetOtpAuthUrlOk() (*string, bool) {
	if o == nil || o.OtpAuthUrl == nil {
		return nil, false
	}
	return o.OtpAuthUrl, true
}

// HasOtpAuthUrl returns a boolean if a field has been set.
func (o *AuthenticationOtpInitializePost200Response) HasOtpAuthUrl() bool {
	if o != nil && o.OtpAuthUrl != nil {
		return true
	}

	return false
}

// SetOtpAuthUrl gets a reference to the given string and assigns it to the OtpAuthUrl field.
func (o *AuthenticationOtpInitializePost200Response) SetOtpAuthUrl(v string) {
	o.OtpAuthUrl = &v
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *AuthenticationOtpInitializePost200Response) GetBarcode() string {
	if o == nil || o.Barcode == nil {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationOtpInitializePost200Response) GetBarcodeOk() (*string, bool) {
	if o == nil || o.Barcode == nil {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *AuthenticationOtpInitializePost200Response) HasBarcode() bool {
	if o != nil && o.Barcode != nil {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *AuthenticationOtpInitializePost200Response) SetBarcode(v string) {
	o.Barcode = &v
}

// GetResponseMessage returns the ResponseMessage field value if set, zero value otherwise.
func (o *AuthenticationOtpInitializePost200Response) GetResponseMessage() string {
	if o == nil || o.ResponseMessage == nil {
		var ret string
		return ret
	}
	return *o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationOtpInitializePost200Response) GetResponseMessageOk() (*string, bool) {
	if o == nil || o.ResponseMessage == nil {
		return nil, false
	}
	return o.ResponseMessage, true
}

// HasResponseMessage returns a boolean if a field has been set.
func (o *AuthenticationOtpInitializePost200Response) HasResponseMessage() bool {
	if o != nil && o.ResponseMessage != nil {
		return true
	}

	return false
}

// SetResponseMessage gets a reference to the given string and assigns it to the ResponseMessage field.
func (o *AuthenticationOtpInitializePost200Response) SetResponseMessage(v string) {
	o.ResponseMessage = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AuthenticationOtpInitializePost200Response) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationOtpInitializePost200Response) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AuthenticationOtpInitializePost200Response) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AuthenticationOtpInitializePost200Response) SetState(v string) {
	o.State = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *AuthenticationOtpInitializePost200Response) GetTimeout() float32 {
	if o == nil || o.Timeout == nil {
		var ret float32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationOtpInitializePost200Response) GetTimeoutOk() (*float32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *AuthenticationOtpInitializePost200Response) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given float32 and assigns it to the Timeout field.
func (o *AuthenticationOtpInitializePost200Response) SetTimeout(v float32) {
	o.Timeout = &v
}

// GetSendPassword returns the SendPassword field value if set, zero value otherwise.
func (o *AuthenticationOtpInitializePost200Response) GetSendPassword() bool {
	if o == nil || o.SendPassword == nil {
		var ret bool
		return ret
	}
	return *o.SendPassword
}

// GetSendPasswordOk returns a tuple with the SendPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationOtpInitializePost200Response) GetSendPasswordOk() (*bool, bool) {
	if o == nil || o.SendPassword == nil {
		return nil, false
	}
	return o.SendPassword, true
}

// HasSendPassword returns a boolean if a field has been set.
func (o *AuthenticationOtpInitializePost200Response) HasSendPassword() bool {
	if o != nil && o.SendPassword != nil {
		return true
	}

	return false
}

// SetSendPassword gets a reference to the given bool and assigns it to the SendPassword field.
func (o *AuthenticationOtpInitializePost200Response) SetSendPassword(v bool) {
	o.SendPassword = &v
}

func (o AuthenticationOtpInitializePost200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.OtpAuthUrl != nil {
		toSerialize["otpAuthUrl"] = o.OtpAuthUrl
	}
	if o.Barcode != nil {
		toSerialize["barcode"] = o.Barcode
	}
	if o.ResponseMessage != nil {
		toSerialize["responseMessage"] = o.ResponseMessage
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.SendPassword != nil {
		toSerialize["sendPassword"] = o.SendPassword
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticationOtpInitializePost200Response struct {
	value *AuthenticationOtpInitializePost200Response
	isSet bool
}

func (v NullableAuthenticationOtpInitializePost200Response) Get() *AuthenticationOtpInitializePost200Response {
	return v.value
}

func (v *NullableAuthenticationOtpInitializePost200Response) Set(val *AuthenticationOtpInitializePost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationOtpInitializePost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationOtpInitializePost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationOtpInitializePost200Response(val *AuthenticationOtpInitializePost200Response) *NullableAuthenticationOtpInitializePost200Response {
	return &NullableAuthenticationOtpInitializePost200Response{value: val, isSet: true}
}

func (v NullableAuthenticationOtpInitializePost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationOtpInitializePost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
