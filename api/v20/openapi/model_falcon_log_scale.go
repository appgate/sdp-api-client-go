/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.3
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FalconLogScale struct for FalconLogScale
type FalconLogScale struct {
	// URL of the Falcon LogScale collector.
	CollectorUrl string `json:"collectorUrl"`
	// Ingest token.
	Token *string `json:"token,omitempty"`
	// Optional name of the repository to ingest into.
	Index *string `json:"index,omitempty"`
	// Translated to \\#type inside Humio. If set, this is used to choose which Humio parser to use for extracting fields.
	SourceType *string `json:"sourceType,omitempty"`
	// Translated to the @source field in Humio.
	Source *string `json:"source,omitempty"`
}

// NewFalconLogScale instantiates a new FalconLogScale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFalconLogScale(collectorUrl string) *FalconLogScale {
	this := FalconLogScale{}
	this.CollectorUrl = collectorUrl
	return &this
}

// NewFalconLogScaleWithDefaults instantiates a new FalconLogScale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFalconLogScaleWithDefaults() *FalconLogScale {
	this := FalconLogScale{}
	return &this
}

// GetCollectorUrl returns the CollectorUrl field value
func (o *FalconLogScale) GetCollectorUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectorUrl
}

// GetCollectorUrlOk returns a tuple with the CollectorUrl field value
// and a boolean to check if the value has been set.
func (o *FalconLogScale) GetCollectorUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectorUrl, true
}

// SetCollectorUrl sets field value
func (o *FalconLogScale) SetCollectorUrl(v string) {
	o.CollectorUrl = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *FalconLogScale) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FalconLogScale) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *FalconLogScale) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *FalconLogScale) SetToken(v string) {
	o.Token = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *FalconLogScale) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FalconLogScale) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *FalconLogScale) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *FalconLogScale) SetIndex(v string) {
	o.Index = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *FalconLogScale) GetSourceType() string {
	if o == nil || o.SourceType == nil {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FalconLogScale) GetSourceTypeOk() (*string, bool) {
	if o == nil || o.SourceType == nil {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *FalconLogScale) HasSourceType() bool {
	if o != nil && o.SourceType != nil {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *FalconLogScale) SetSourceType(v string) {
	o.SourceType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *FalconLogScale) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FalconLogScale) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *FalconLogScale) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *FalconLogScale) SetSource(v string) {
	o.Source = &v
}

func (o FalconLogScale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["collectorUrl"] = o.CollectorUrl
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.SourceType != nil {
		toSerialize["sourceType"] = o.SourceType
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableFalconLogScale struct {
	value *FalconLogScale
	isSet bool
}

func (v NullableFalconLogScale) Get() *FalconLogScale {
	return v.value
}

func (v *NullableFalconLogScale) Set(val *FalconLogScale) {
	v.value = val
	v.isSet = true
}

func (v NullableFalconLogScale) IsSet() bool {
	return v.isSet
}

func (v *NullableFalconLogScale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFalconLogScale(val *FalconLogScale) *NullableFalconLogScale {
	return &NullableFalconLogScale{value: val, isSet: true}
}

func (v NullableFalconLogScale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFalconLogScale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
