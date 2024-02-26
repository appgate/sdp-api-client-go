/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v19+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 19.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BaseActiveSession struct for BaseActiveSession
type BaseActiveSession struct {
	// Distinguished name of a user&device combination. Format: \"CN=,CN=,OU=\"
	DistinguishedName *string `json:"distinguishedName,omitempty"`
	// The device ID, same as the one in the Distinguished Name.
	DeviceId *string `json:"deviceId,omitempty"`
	// The username, same as the one in the Distinguished Name.
	Username *string `json:"username,omitempty"`
	// The provider name of the user, same as the one in the Distinguished Name.
	ProviderName *string `json:"providerName,omitempty"`
	// GeoIP Latitude of the client.
	GeoIpLatitude *float64 `json:"geoIpLatitude,omitempty"`
	// GeoIP Latitude of the client.
	GeoIpLongitude *float64 `json:"geoIpLongitude,omitempty"`
	// Hostname or Device name reported by the Client as Device Claim.
	Hostname *string `json:"hostname,omitempty"`
	// Family of the operating system of the Client.
	OsFamily *string `json:"osFamily,omitempty"`
	// Name of the operating system of the Client.
	OsName *string `json:"osName,omitempty"`
	// Parent of the operating system of the Client. Typically available for Linux.
	OsParent *string `json:"osParent,omitempty"`
	// Client version.
	ClientVersion *string `json:"clientVersion,omitempty"`
	// Client type.
	ClientType *string `json:"clientType,omitempty"`
	// The support status of the client.
	ClientSupport *string `json:"clientSupport,omitempty"`
}

// NewBaseActiveSession instantiates a new BaseActiveSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseActiveSession() *BaseActiveSession {
	this := BaseActiveSession{}
	return &this
}

// NewBaseActiveSessionWithDefaults instantiates a new BaseActiveSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseActiveSessionWithDefaults() *BaseActiveSession {
	this := BaseActiveSession{}
	return &this
}

// GetDistinguishedName returns the DistinguishedName field value if set, zero value otherwise.
func (o *BaseActiveSession) GetDistinguishedName() string {
	if o == nil || o.DistinguishedName == nil {
		var ret string
		return ret
	}
	return *o.DistinguishedName
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetDistinguishedNameOk() (*string, bool) {
	if o == nil || o.DistinguishedName == nil {
		return nil, false
	}
	return o.DistinguishedName, true
}

// HasDistinguishedName returns a boolean if a field has been set.
func (o *BaseActiveSession) HasDistinguishedName() bool {
	if o != nil && o.DistinguishedName != nil {
		return true
	}

	return false
}

// SetDistinguishedName gets a reference to the given string and assigns it to the DistinguishedName field.
func (o *BaseActiveSession) SetDistinguishedName(v string) {
	o.DistinguishedName = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *BaseActiveSession) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *BaseActiveSession) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *BaseActiveSession) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *BaseActiveSession) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *BaseActiveSession) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *BaseActiveSession) SetUsername(v string) {
	o.Username = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *BaseActiveSession) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *BaseActiveSession) HasProviderName() bool {
	if o != nil && o.ProviderName != nil {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *BaseActiveSession) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetGeoIpLatitude returns the GeoIpLatitude field value if set, zero value otherwise.
func (o *BaseActiveSession) GetGeoIpLatitude() float64 {
	if o == nil || o.GeoIpLatitude == nil {
		var ret float64
		return ret
	}
	return *o.GeoIpLatitude
}

// GetGeoIpLatitudeOk returns a tuple with the GeoIpLatitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetGeoIpLatitudeOk() (*float64, bool) {
	if o == nil || o.GeoIpLatitude == nil {
		return nil, false
	}
	return o.GeoIpLatitude, true
}

// HasGeoIpLatitude returns a boolean if a field has been set.
func (o *BaseActiveSession) HasGeoIpLatitude() bool {
	if o != nil && o.GeoIpLatitude != nil {
		return true
	}

	return false
}

// SetGeoIpLatitude gets a reference to the given float64 and assigns it to the GeoIpLatitude field.
func (o *BaseActiveSession) SetGeoIpLatitude(v float64) {
	o.GeoIpLatitude = &v
}

// GetGeoIpLongitude returns the GeoIpLongitude field value if set, zero value otherwise.
func (o *BaseActiveSession) GetGeoIpLongitude() float64 {
	if o == nil || o.GeoIpLongitude == nil {
		var ret float64
		return ret
	}
	return *o.GeoIpLongitude
}

// GetGeoIpLongitudeOk returns a tuple with the GeoIpLongitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetGeoIpLongitudeOk() (*float64, bool) {
	if o == nil || o.GeoIpLongitude == nil {
		return nil, false
	}
	return o.GeoIpLongitude, true
}

// HasGeoIpLongitude returns a boolean if a field has been set.
func (o *BaseActiveSession) HasGeoIpLongitude() bool {
	if o != nil && o.GeoIpLongitude != nil {
		return true
	}

	return false
}

// SetGeoIpLongitude gets a reference to the given float64 and assigns it to the GeoIpLongitude field.
func (o *BaseActiveSession) SetGeoIpLongitude(v float64) {
	o.GeoIpLongitude = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *BaseActiveSession) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *BaseActiveSession) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *BaseActiveSession) SetHostname(v string) {
	o.Hostname = &v
}

// GetOsFamily returns the OsFamily field value if set, zero value otherwise.
func (o *BaseActiveSession) GetOsFamily() string {
	if o == nil || o.OsFamily == nil {
		var ret string
		return ret
	}
	return *o.OsFamily
}

// GetOsFamilyOk returns a tuple with the OsFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetOsFamilyOk() (*string, bool) {
	if o == nil || o.OsFamily == nil {
		return nil, false
	}
	return o.OsFamily, true
}

// HasOsFamily returns a boolean if a field has been set.
func (o *BaseActiveSession) HasOsFamily() bool {
	if o != nil && o.OsFamily != nil {
		return true
	}

	return false
}

// SetOsFamily gets a reference to the given string and assigns it to the OsFamily field.
func (o *BaseActiveSession) SetOsFamily(v string) {
	o.OsFamily = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *BaseActiveSession) GetOsName() string {
	if o == nil || o.OsName == nil {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetOsNameOk() (*string, bool) {
	if o == nil || o.OsName == nil {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *BaseActiveSession) HasOsName() bool {
	if o != nil && o.OsName != nil {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *BaseActiveSession) SetOsName(v string) {
	o.OsName = &v
}

// GetOsParent returns the OsParent field value if set, zero value otherwise.
func (o *BaseActiveSession) GetOsParent() string {
	if o == nil || o.OsParent == nil {
		var ret string
		return ret
	}
	return *o.OsParent
}

// GetOsParentOk returns a tuple with the OsParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetOsParentOk() (*string, bool) {
	if o == nil || o.OsParent == nil {
		return nil, false
	}
	return o.OsParent, true
}

// HasOsParent returns a boolean if a field has been set.
func (o *BaseActiveSession) HasOsParent() bool {
	if o != nil && o.OsParent != nil {
		return true
	}

	return false
}

// SetOsParent gets a reference to the given string and assigns it to the OsParent field.
func (o *BaseActiveSession) SetOsParent(v string) {
	o.OsParent = &v
}

// GetClientVersion returns the ClientVersion field value if set, zero value otherwise.
func (o *BaseActiveSession) GetClientVersion() string {
	if o == nil || o.ClientVersion == nil {
		var ret string
		return ret
	}
	return *o.ClientVersion
}

// GetClientVersionOk returns a tuple with the ClientVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetClientVersionOk() (*string, bool) {
	if o == nil || o.ClientVersion == nil {
		return nil, false
	}
	return o.ClientVersion, true
}

// HasClientVersion returns a boolean if a field has been set.
func (o *BaseActiveSession) HasClientVersion() bool {
	if o != nil && o.ClientVersion != nil {
		return true
	}

	return false
}

// SetClientVersion gets a reference to the given string and assigns it to the ClientVersion field.
func (o *BaseActiveSession) SetClientVersion(v string) {
	o.ClientVersion = &v
}

// GetClientType returns the ClientType field value if set, zero value otherwise.
func (o *BaseActiveSession) GetClientType() string {
	if o == nil || o.ClientType == nil {
		var ret string
		return ret
	}
	return *o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetClientTypeOk() (*string, bool) {
	if o == nil || o.ClientType == nil {
		return nil, false
	}
	return o.ClientType, true
}

// HasClientType returns a boolean if a field has been set.
func (o *BaseActiveSession) HasClientType() bool {
	if o != nil && o.ClientType != nil {
		return true
	}

	return false
}

// SetClientType gets a reference to the given string and assigns it to the ClientType field.
func (o *BaseActiveSession) SetClientType(v string) {
	o.ClientType = &v
}

// GetClientSupport returns the ClientSupport field value if set, zero value otherwise.
func (o *BaseActiveSession) GetClientSupport() string {
	if o == nil || o.ClientSupport == nil {
		var ret string
		return ret
	}
	return *o.ClientSupport
}

// GetClientSupportOk returns a tuple with the ClientSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseActiveSession) GetClientSupportOk() (*string, bool) {
	if o == nil || o.ClientSupport == nil {
		return nil, false
	}
	return o.ClientSupport, true
}

// HasClientSupport returns a boolean if a field has been set.
func (o *BaseActiveSession) HasClientSupport() bool {
	if o != nil && o.ClientSupport != nil {
		return true
	}

	return false
}

// SetClientSupport gets a reference to the given string and assigns it to the ClientSupport field.
func (o *BaseActiveSession) SetClientSupport(v string) {
	o.ClientSupport = &v
}

func (o BaseActiveSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DistinguishedName != nil {
		toSerialize["distinguishedName"] = o.DistinguishedName
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
	if o.GeoIpLatitude != nil {
		toSerialize["geoIpLatitude"] = o.GeoIpLatitude
	}
	if o.GeoIpLongitude != nil {
		toSerialize["geoIpLongitude"] = o.GeoIpLongitude
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.OsFamily != nil {
		toSerialize["osFamily"] = o.OsFamily
	}
	if o.OsName != nil {
		toSerialize["osName"] = o.OsName
	}
	if o.OsParent != nil {
		toSerialize["osParent"] = o.OsParent
	}
	if o.ClientVersion != nil {
		toSerialize["clientVersion"] = o.ClientVersion
	}
	if o.ClientType != nil {
		toSerialize["clientType"] = o.ClientType
	}
	if o.ClientSupport != nil {
		toSerialize["clientSupport"] = o.ClientSupport
	}
	return json.Marshal(toSerialize)
}

type NullableBaseActiveSession struct {
	value *BaseActiveSession
	isSet bool
}

func (v NullableBaseActiveSession) Get() *BaseActiveSession {
	return v.value
}

func (v *NullableBaseActiveSession) Set(val *BaseActiveSession) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseActiveSession) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseActiveSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseActiveSession(val *BaseActiveSession) *NullableBaseActiveSession {
	return &NullableBaseActiveSession{value: val, isSet: true}
}

func (v NullableBaseActiveSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseActiveSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
