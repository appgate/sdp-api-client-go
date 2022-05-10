/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// TokenRevocation Response for token revocation response.
type TokenRevocation struct {
	// The ID assigned to the token.
	TokenId *string `json:"tokenId,omitempty"`
	// The type of the token that was revoked.
	TokenType *string `json:"tokenType,omitempty"`
	// The Distinguished name of a user&device combination. Format: \"CN=,CN=,OU=\"
	DistinguishedName *string `json:"distinguishedName,omitempty"`
	// The issue time of the token.
	Issued *time.Time `json:"issued,omitempty"`
	// The expiration time of the token.
	Expires *time.Time `json:"expires,omitempty"`
	// Whether or not the token was revoked.
	Revoked *bool `json:"revoked,omitempty"`
	// The site ID of the token.
	SiteId *string `json:"siteId,omitempty"`
	// The name of the Site for this Entitlement. For convenience only.
	SiteName *string `json:"siteName,omitempty"`
	// The revocation time of the certificate.
	RevocationTime *time.Time `json:"revocationTime,omitempty"`
	// The device ID to assign to this Client. It will be used to generate device distinguished name.
	DeviceId *string `json:"deviceId,omitempty"`
	// The username, same as the one in the Distinguished Name.
	Username *string `json:"username,omitempty"`
	// The provider name of the user, same as the one in the user Distinguished Name.
	ProviderName *string `json:"providerName,omitempty"`
	// The hostname of the controller associated with this token.
	ControllerHostname *string `json:"controllerHostname,omitempty"`
}

// NewTokenRevocation instantiates a new TokenRevocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRevocation() *TokenRevocation {
	this := TokenRevocation{}
	return &this
}

// NewTokenRevocationWithDefaults instantiates a new TokenRevocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRevocationWithDefaults() *TokenRevocation {
	this := TokenRevocation{}
	return &this
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *TokenRevocation) GetTokenId() string {
	if o == nil || o.TokenId == nil {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetTokenIdOk() (*string, bool) {
	if o == nil || o.TokenId == nil {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *TokenRevocation) HasTokenId() bool {
	if o != nil && o.TokenId != nil {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *TokenRevocation) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *TokenRevocation) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *TokenRevocation) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *TokenRevocation) SetTokenType(v string) {
	o.TokenType = &v
}

// GetDistinguishedName returns the DistinguishedName field value if set, zero value otherwise.
func (o *TokenRevocation) GetDistinguishedName() string {
	if o == nil || o.DistinguishedName == nil {
		var ret string
		return ret
	}
	return *o.DistinguishedName
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetDistinguishedNameOk() (*string, bool) {
	if o == nil || o.DistinguishedName == nil {
		return nil, false
	}
	return o.DistinguishedName, true
}

// HasDistinguishedName returns a boolean if a field has been set.
func (o *TokenRevocation) HasDistinguishedName() bool {
	if o != nil && o.DistinguishedName != nil {
		return true
	}

	return false
}

// SetDistinguishedName gets a reference to the given string and assigns it to the DistinguishedName field.
func (o *TokenRevocation) SetDistinguishedName(v string) {
	o.DistinguishedName = &v
}

// GetIssued returns the Issued field value if set, zero value otherwise.
func (o *TokenRevocation) GetIssued() time.Time {
	if o == nil || o.Issued == nil {
		var ret time.Time
		return ret
	}
	return *o.Issued
}

// GetIssuedOk returns a tuple with the Issued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetIssuedOk() (*time.Time, bool) {
	if o == nil || o.Issued == nil {
		return nil, false
	}
	return o.Issued, true
}

// HasIssued returns a boolean if a field has been set.
func (o *TokenRevocation) HasIssued() bool {
	if o != nil && o.Issued != nil {
		return true
	}

	return false
}

// SetIssued gets a reference to the given time.Time and assigns it to the Issued field.
func (o *TokenRevocation) SetIssued(v time.Time) {
	o.Issued = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *TokenRevocation) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *TokenRevocation) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *TokenRevocation) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *TokenRevocation) GetRevoked() bool {
	if o == nil || o.Revoked == nil {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetRevokedOk() (*bool, bool) {
	if o == nil || o.Revoked == nil {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *TokenRevocation) HasRevoked() bool {
	if o != nil && o.Revoked != nil {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *TokenRevocation) SetRevoked(v bool) {
	o.Revoked = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *TokenRevocation) GetSiteId() string {
	if o == nil || o.SiteId == nil {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetSiteIdOk() (*string, bool) {
	if o == nil || o.SiteId == nil {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *TokenRevocation) HasSiteId() bool {
	if o != nil && o.SiteId != nil {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *TokenRevocation) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *TokenRevocation) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *TokenRevocation) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *TokenRevocation) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRevocationTime returns the RevocationTime field value if set, zero value otherwise.
func (o *TokenRevocation) GetRevocationTime() time.Time {
	if o == nil || o.RevocationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.RevocationTime
}

// GetRevocationTimeOk returns a tuple with the RevocationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetRevocationTimeOk() (*time.Time, bool) {
	if o == nil || o.RevocationTime == nil {
		return nil, false
	}
	return o.RevocationTime, true
}

// HasRevocationTime returns a boolean if a field has been set.
func (o *TokenRevocation) HasRevocationTime() bool {
	if o != nil && o.RevocationTime != nil {
		return true
	}

	return false
}

// SetRevocationTime gets a reference to the given time.Time and assigns it to the RevocationTime field.
func (o *TokenRevocation) SetRevocationTime(v time.Time) {
	o.RevocationTime = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *TokenRevocation) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *TokenRevocation) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *TokenRevocation) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *TokenRevocation) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *TokenRevocation) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *TokenRevocation) SetUsername(v string) {
	o.Username = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *TokenRevocation) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *TokenRevocation) HasProviderName() bool {
	if o != nil && o.ProviderName != nil {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *TokenRevocation) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetControllerHostname returns the ControllerHostname field value if set, zero value otherwise.
func (o *TokenRevocation) GetControllerHostname() string {
	if o == nil || o.ControllerHostname == nil {
		var ret string
		return ret
	}
	return *o.ControllerHostname
}

// GetControllerHostnameOk returns a tuple with the ControllerHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRevocation) GetControllerHostnameOk() (*string, bool) {
	if o == nil || o.ControllerHostname == nil {
		return nil, false
	}
	return o.ControllerHostname, true
}

// HasControllerHostname returns a boolean if a field has been set.
func (o *TokenRevocation) HasControllerHostname() bool {
	if o != nil && o.ControllerHostname != nil {
		return true
	}

	return false
}

// SetControllerHostname gets a reference to the given string and assigns it to the ControllerHostname field.
func (o *TokenRevocation) SetControllerHostname(v string) {
	o.ControllerHostname = &v
}

func (o TokenRevocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TokenId != nil {
		toSerialize["tokenId"] = o.TokenId
	}
	if o.TokenType != nil {
		toSerialize["tokenType"] = o.TokenType
	}
	if o.DistinguishedName != nil {
		toSerialize["distinguishedName"] = o.DistinguishedName
	}
	if o.Issued != nil {
		toSerialize["issued"] = o.Issued
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.Revoked != nil {
		toSerialize["revoked"] = o.Revoked
	}
	if o.SiteId != nil {
		toSerialize["siteId"] = o.SiteId
	}
	if o.SiteName != nil {
		toSerialize["siteName"] = o.SiteName
	}
	if o.RevocationTime != nil {
		toSerialize["revocationTime"] = o.RevocationTime
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.ProviderName != nil {
		toSerialize["providerName"] = o.ProviderName
	}
	if o.ControllerHostname != nil {
		toSerialize["controllerHostname"] = o.ControllerHostname
	}
	return json.Marshal(toSerialize)
}

type NullableTokenRevocation struct {
	value *TokenRevocation
	isSet bool
}

func (v NullableTokenRevocation) Get() *TokenRevocation {
	return v.value
}

func (v *NullableTokenRevocation) Set(val *TokenRevocation) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRevocation) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRevocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRevocation(val *TokenRevocation) *NullableTokenRevocation {
	return &NullableTokenRevocation{value: val, isSet: true}
}

func (v NullableTokenRevocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRevocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
