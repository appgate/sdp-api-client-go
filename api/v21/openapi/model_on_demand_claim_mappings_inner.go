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

// OnDemandClaimMappingsInner struct for OnDemandClaimMappingsInner
type OnDemandClaimMappingsInner struct {
	// The name of the command.
	Command string `json:"command"`
	// The name of the device claim to be used in Appgate SDP.
	ClaimName  string                                `json:"claimName"`
	Parameters *OnDemandClaimMappingsInnerParameters `json:"parameters,omitempty"`
	// The platform(s) to run the on-demand claim.
	Platform string `json:"platform"`
}

// NewOnDemandClaimMappingsInner instantiates a new OnDemandClaimMappingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandClaimMappingsInner(command string, claimName string, platform string) *OnDemandClaimMappingsInner {
	this := OnDemandClaimMappingsInner{}
	this.Command = command
	this.ClaimName = claimName
	this.Platform = platform
	return &this
}

// NewOnDemandClaimMappingsInnerWithDefaults instantiates a new OnDemandClaimMappingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandClaimMappingsInnerWithDefaults() *OnDemandClaimMappingsInner {
	this := OnDemandClaimMappingsInner{}
	return &this
}

// GetCommand returns the Command field value
func (o *OnDemandClaimMappingsInner) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *OnDemandClaimMappingsInner) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *OnDemandClaimMappingsInner) SetCommand(v string) {
	o.Command = v
}

// GetClaimName returns the ClaimName field value
func (o *OnDemandClaimMappingsInner) GetClaimName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClaimName
}

// GetClaimNameOk returns a tuple with the ClaimName field value
// and a boolean to check if the value has been set.
func (o *OnDemandClaimMappingsInner) GetClaimNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimName, true
}

// SetClaimName sets field value
func (o *OnDemandClaimMappingsInner) SetClaimName(v string) {
	o.ClaimName = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *OnDemandClaimMappingsInner) GetParameters() OnDemandClaimMappingsInnerParameters {
	if o == nil || o.Parameters == nil {
		var ret OnDemandClaimMappingsInnerParameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandClaimMappingsInner) GetParametersOk() (*OnDemandClaimMappingsInnerParameters, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *OnDemandClaimMappingsInner) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given OnDemandClaimMappingsInnerParameters and assigns it to the Parameters field.
func (o *OnDemandClaimMappingsInner) SetParameters(v OnDemandClaimMappingsInnerParameters) {
	o.Parameters = &v
}

// GetPlatform returns the Platform field value
func (o *OnDemandClaimMappingsInner) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *OnDemandClaimMappingsInner) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *OnDemandClaimMappingsInner) SetPlatform(v string) {
	o.Platform = v
}

func (o OnDemandClaimMappingsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["command"] = o.Command
	}
	if true {
		toSerialize["claimName"] = o.ClaimName
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if true {
		toSerialize["platform"] = o.Platform
	}
	return json.Marshal(toSerialize)
}

type NullableOnDemandClaimMappingsInner struct {
	value *OnDemandClaimMappingsInner
	isSet bool
}

func (v NullableOnDemandClaimMappingsInner) Get() *OnDemandClaimMappingsInner {
	return v.value
}

func (v *NullableOnDemandClaimMappingsInner) Set(val *OnDemandClaimMappingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandClaimMappingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandClaimMappingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandClaimMappingsInner(val *OnDemandClaimMappingsInner) *NullableOnDemandClaimMappingsInner {
	return &NullableOnDemandClaimMappingsInner{value: val, isSet: true}
}

func (v NullableOnDemandClaimMappingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandClaimMappingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
