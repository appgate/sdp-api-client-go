/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v21+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 21.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ClaimNamesInner Details of a Claim.
type ClaimNamesInner struct {
	// Name of the Claim that's available to the JavaScript engine.
	ClaimName *string `json:"claimName,omitempty"`
	// Data type of the Claim.
	Type *string `json:"type,omitempty"`
	// Whether the claim is available to Policy and Criteria Scripts or not. All Claims are available to Conditions and Entitlement Scripts.
	AvailableForPolicy *bool `json:"availableForPolicy,omitempty"`
}

// NewClaimNamesInner instantiates a new ClaimNamesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimNamesInner() *ClaimNamesInner {
	this := ClaimNamesInner{}
	return &this
}

// NewClaimNamesInnerWithDefaults instantiates a new ClaimNamesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimNamesInnerWithDefaults() *ClaimNamesInner {
	this := ClaimNamesInner{}
	return &this
}

// GetClaimName returns the ClaimName field value if set, zero value otherwise.
func (o *ClaimNamesInner) GetClaimName() string {
	if o == nil || o.ClaimName == nil {
		var ret string
		return ret
	}
	return *o.ClaimName
}

// GetClaimNameOk returns a tuple with the ClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimNamesInner) GetClaimNameOk() (*string, bool) {
	if o == nil || o.ClaimName == nil {
		return nil, false
	}
	return o.ClaimName, true
}

// HasClaimName returns a boolean if a field has been set.
func (o *ClaimNamesInner) HasClaimName() bool {
	if o != nil && o.ClaimName != nil {
		return true
	}

	return false
}

// SetClaimName gets a reference to the given string and assigns it to the ClaimName field.
func (o *ClaimNamesInner) SetClaimName(v string) {
	o.ClaimName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClaimNamesInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimNamesInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClaimNamesInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClaimNamesInner) SetType(v string) {
	o.Type = &v
}

// GetAvailableForPolicy returns the AvailableForPolicy field value if set, zero value otherwise.
func (o *ClaimNamesInner) GetAvailableForPolicy() bool {
	if o == nil || o.AvailableForPolicy == nil {
		var ret bool
		return ret
	}
	return *o.AvailableForPolicy
}

// GetAvailableForPolicyOk returns a tuple with the AvailableForPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimNamesInner) GetAvailableForPolicyOk() (*bool, bool) {
	if o == nil || o.AvailableForPolicy == nil {
		return nil, false
	}
	return o.AvailableForPolicy, true
}

// HasAvailableForPolicy returns a boolean if a field has been set.
func (o *ClaimNamesInner) HasAvailableForPolicy() bool {
	if o != nil && o.AvailableForPolicy != nil {
		return true
	}

	return false
}

// SetAvailableForPolicy gets a reference to the given bool and assigns it to the AvailableForPolicy field.
func (o *ClaimNamesInner) SetAvailableForPolicy(v bool) {
	o.AvailableForPolicy = &v
}

func (o ClaimNamesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClaimName != nil {
		toSerialize["claimName"] = o.ClaimName
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.AvailableForPolicy != nil {
		toSerialize["availableForPolicy"] = o.AvailableForPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableClaimNamesInner struct {
	value *ClaimNamesInner
	isSet bool
}

func (v NullableClaimNamesInner) Get() *ClaimNamesInner {
	return v.value
}

func (v *NullableClaimNamesInner) Set(val *ClaimNamesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimNamesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimNamesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimNamesInner(val *ClaimNamesInner) *NullableClaimNamesInner {
	return &NullableClaimNamesInner{value: val, isSet: true}
}

func (v NullableClaimNamesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimNamesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
