/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfPeerInterface The details of peer connection interface. Used by other appliances and administrative UI. This interface is deprecated as of 5.4. All connections will be handled by clientInterface and adminInterface in the future. The hostname field is used as identifier and will take over the hostname field in the root of Appliance when this interface is removed.
type ApplianceAllOfPeerInterface struct {
	// Hostname to connect by the peers. It will be used to validate the appliance certificate.
	Hostname string `json:"hostname"`
	// Port to connect for peer specific services.
	HttpsPort *int32 `json:"httpsPort,omitempty"`
	// Source configuration to allow via iptables.
	AllowSources []AllowSourcesInner `json:"allowSources,omitempty"`
}

// NewApplianceAllOfPeerInterface instantiates a new ApplianceAllOfPeerInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfPeerInterface(hostname string) *ApplianceAllOfPeerInterface {
	this := ApplianceAllOfPeerInterface{}
	this.Hostname = hostname
	var httpsPort int32 = 444
	this.HttpsPort = &httpsPort
	return &this
}

// NewApplianceAllOfPeerInterfaceWithDefaults instantiates a new ApplianceAllOfPeerInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfPeerInterfaceWithDefaults() *ApplianceAllOfPeerInterface {
	this := ApplianceAllOfPeerInterface{}
	var httpsPort int32 = 444
	this.HttpsPort = &httpsPort
	return &this
}

// GetHostname returns the Hostname field value
func (o *ApplianceAllOfPeerInterface) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfPeerInterface) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *ApplianceAllOfPeerInterface) SetHostname(v string) {
	o.Hostname = v
}

// GetHttpsPort returns the HttpsPort field value if set, zero value otherwise.
func (o *ApplianceAllOfPeerInterface) GetHttpsPort() int32 {
	if o == nil || o.HttpsPort == nil {
		var ret int32
		return ret
	}
	return *o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfPeerInterface) GetHttpsPortOk() (*int32, bool) {
	if o == nil || o.HttpsPort == nil {
		return nil, false
	}
	return o.HttpsPort, true
}

// HasHttpsPort returns a boolean if a field has been set.
func (o *ApplianceAllOfPeerInterface) HasHttpsPort() bool {
	if o != nil && o.HttpsPort != nil {
		return true
	}

	return false
}

// SetHttpsPort gets a reference to the given int32 and assigns it to the HttpsPort field.
func (o *ApplianceAllOfPeerInterface) SetHttpsPort(v int32) {
	o.HttpsPort = &v
}

// GetAllowSources returns the AllowSources field value if set, zero value otherwise.
func (o *ApplianceAllOfPeerInterface) GetAllowSources() []AllowSourcesInner {
	if o == nil || o.AllowSources == nil {
		var ret []AllowSourcesInner
		return ret
	}
	return o.AllowSources
}

// GetAllowSourcesOk returns a tuple with the AllowSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfPeerInterface) GetAllowSourcesOk() ([]AllowSourcesInner, bool) {
	if o == nil || o.AllowSources == nil {
		return nil, false
	}
	return o.AllowSources, true
}

// HasAllowSources returns a boolean if a field has been set.
func (o *ApplianceAllOfPeerInterface) HasAllowSources() bool {
	if o != nil && o.AllowSources != nil {
		return true
	}

	return false
}

// SetAllowSources gets a reference to the given []AllowSourcesInner and assigns it to the AllowSources field.
func (o *ApplianceAllOfPeerInterface) SetAllowSources(v []AllowSourcesInner) {
	o.AllowSources = v
}

func (o ApplianceAllOfPeerInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if o.HttpsPort != nil {
		toSerialize["httpsPort"] = o.HttpsPort
	}
	if o.AllowSources != nil {
		toSerialize["allowSources"] = o.AllowSources
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfPeerInterface struct {
	value *ApplianceAllOfPeerInterface
	isSet bool
}

func (v NullableApplianceAllOfPeerInterface) Get() *ApplianceAllOfPeerInterface {
	return v.value
}

func (v *NullableApplianceAllOfPeerInterface) Set(val *ApplianceAllOfPeerInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfPeerInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfPeerInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfPeerInterface(val *ApplianceAllOfPeerInterface) *NullableApplianceAllOfPeerInterface {
	return &NullableApplianceAllOfPeerInterface{value: val, isSet: true}
}

func (v NullableApplianceAllOfPeerInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfPeerInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
