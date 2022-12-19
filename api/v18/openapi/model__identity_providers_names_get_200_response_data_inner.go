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

// IdentityProvidersNamesGet200ResponseDataInner Details of an identity provider required for logging in.
type IdentityProvidersNamesGet200ResponseDataInner struct {
	// Name of the identity provider.
	Name *string `json:"name,omitempty"`
	// The type of the identity provider.
	Type *string `json:"type,omitempty"`
	// The SAML login URL.
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// The prioritization of which client certificate to use.
	CertificatePriorities []map[string]interface{} `json:"certificatePriorities,omitempty"`
	// Authorization endpoint URL for OIDC login.
	AuthUrl *string `json:"authUrl,omitempty"`
	// Token endpoint URL for OIDC login.
	TokenUrl *string `json:"tokenUrl,omitempty"`
	// Client ID for OIDC login.
	ClientId *string `json:"clientId,omitempty"`
	// OIDC scope for getting user information.
	Scope *string `json:"scope,omitempty"`
}

// NewIdentityProvidersNamesGet200ResponseDataInner instantiates a new IdentityProvidersNamesGet200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProvidersNamesGet200ResponseDataInner() *IdentityProvidersNamesGet200ResponseDataInner {
	this := IdentityProvidersNamesGet200ResponseDataInner{}
	return &this
}

// NewIdentityProvidersNamesGet200ResponseDataInnerWithDefaults instantiates a new IdentityProvidersNamesGet200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProvidersNamesGet200ResponseDataInnerWithDefaults() *IdentityProvidersNamesGet200ResponseDataInner {
	this := IdentityProvidersNamesGet200ResponseDataInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityProvidersNamesGet200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentityProvidersNamesGet200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl != nil {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *IdentityProvidersNamesGet200ResponseDataInner) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetCertificatePriorities returns the CertificatePriorities field value if set, zero value otherwise.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetCertificatePriorities() []map[string]interface{} {
	if o == nil || o.CertificatePriorities == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.CertificatePriorities
}

// GetCertificatePrioritiesOk returns a tuple with the CertificatePriorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetCertificatePrioritiesOk() ([]map[string]interface{}, bool) {
	if o == nil || o.CertificatePriorities == nil {
		return nil, false
	}
	return o.CertificatePriorities, true
}

// HasCertificatePriorities returns a boolean if a field has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) HasCertificatePriorities() bool {
	if o != nil && o.CertificatePriorities != nil {
		return true
	}

	return false
}

// SetCertificatePriorities gets a reference to the given []map[string]interface{} and assigns it to the CertificatePriorities field.
func (o *IdentityProvidersNamesGet200ResponseDataInner) SetCertificatePriorities(v []map[string]interface{}) {
	o.CertificatePriorities = v
}

// GetAuthUrl returns the AuthUrl field value if set, zero value otherwise.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetAuthUrl() string {
	if o == nil || o.AuthUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthUrl
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetAuthUrlOk() (*string, bool) {
	if o == nil || o.AuthUrl == nil {
		return nil, false
	}
	return o.AuthUrl, true
}

// HasAuthUrl returns a boolean if a field has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) HasAuthUrl() bool {
	if o != nil && o.AuthUrl != nil {
		return true
	}

	return false
}

// SetAuthUrl gets a reference to the given string and assigns it to the AuthUrl field.
func (o *IdentityProvidersNamesGet200ResponseDataInner) SetAuthUrl(v string) {
	o.AuthUrl = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetTokenUrl() string {
	if o == nil || o.TokenUrl == nil {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetTokenUrlOk() (*string, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *IdentityProvidersNamesGet200ResponseDataInner) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IdentityProvidersNamesGet200ResponseDataInner) SetClientId(v string) {
	o.ClientId = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *IdentityProvidersNamesGet200ResponseDataInner) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *IdentityProvidersNamesGet200ResponseDataInner) SetScope(v string) {
	o.Scope = &v
}

func (o IdentityProvidersNamesGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.RedirectUrl != nil {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if o.CertificatePriorities != nil {
		toSerialize["certificatePriorities"] = o.CertificatePriorities
	}
	if o.AuthUrl != nil {
		toSerialize["authUrl"] = o.AuthUrl
	}
	if o.TokenUrl != nil {
		toSerialize["tokenUrl"] = o.TokenUrl
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityProvidersNamesGet200ResponseDataInner struct {
	value *IdentityProvidersNamesGet200ResponseDataInner
	isSet bool
}

func (v NullableIdentityProvidersNamesGet200ResponseDataInner) Get() *IdentityProvidersNamesGet200ResponseDataInner {
	return v.value
}

func (v *NullableIdentityProvidersNamesGet200ResponseDataInner) Set(val *IdentityProvidersNamesGet200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProvidersNamesGet200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProvidersNamesGet200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProvidersNamesGet200ResponseDataInner(val *IdentityProvidersNamesGet200ResponseDataInner) *NullableIdentityProvidersNamesGet200ResponseDataInner {
	return &NullableIdentityProvidersNamesGet200ResponseDataInner{value: val, isSet: true}
}

func (v NullableIdentityProvidersNamesGet200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProvidersNamesGet200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
