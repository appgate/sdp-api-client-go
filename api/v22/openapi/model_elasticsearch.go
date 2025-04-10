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

// Elasticsearch struct for Elasticsearch
type Elasticsearch struct {
	// AWS ID to login. Only required if AWS Access Keys are being used to authenticate.
	AwsId *string `json:"awsId,omitempty"`
	// AWS secret to login. Only required if AWS Access Keys are being used to authenticate.
	AwsSecret *string `json:"awsSecret,omitempty"`
	// AWS region. Only required if AWS Access Keys are being used to authenticate.
	AwsRegion *string `json:"awsRegion,omitempty"`
	// Whether to use the credentials from the AWS instance or not.
	UseInstanceCredentials *bool `json:"useInstanceCredentials,omitempty"`
	// The URL of the elasticsearch server.
	Url string `json:"url"`
	// Optional field to enable log retention on the configured AWS elasticsearch. Defines how many days the audit logs will be kept.
	RetentionDays *int32 `json:"retentionDays,omitempty"`
	// Which version of Elasticsearch that logs are forwarded to.
	CompatibilityMode *int32                            `json:"compatibilityMode,omitempty"`
	Authentication    *ElasticsearchAllOfAuthentication `json:"authentication,omitempty"`
}

// NewElasticsearch instantiates a new Elasticsearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewElasticsearch(url string) *Elasticsearch {
	this := Elasticsearch{}
	this.Url = url
	var compatibilityMode int32 = 6
	this.CompatibilityMode = &compatibilityMode
	return &this
}

// NewElasticsearchWithDefaults instantiates a new Elasticsearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewElasticsearchWithDefaults() *Elasticsearch {
	this := Elasticsearch{}
	var compatibilityMode int32 = 6
	this.CompatibilityMode = &compatibilityMode
	return &this
}

// GetAwsId returns the AwsId field value if set, zero value otherwise.
func (o *Elasticsearch) GetAwsId() string {
	if o == nil || o.AwsId == nil {
		var ret string
		return ret
	}
	return *o.AwsId
}

// GetAwsIdOk returns a tuple with the AwsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elasticsearch) GetAwsIdOk() (*string, bool) {
	if o == nil || o.AwsId == nil {
		return nil, false
	}
	return o.AwsId, true
}

// HasAwsId returns a boolean if a field has been set.
func (o *Elasticsearch) HasAwsId() bool {
	if o != nil && o.AwsId != nil {
		return true
	}

	return false
}

// SetAwsId gets a reference to the given string and assigns it to the AwsId field.
func (o *Elasticsearch) SetAwsId(v string) {
	o.AwsId = &v
}

// GetAwsSecret returns the AwsSecret field value if set, zero value otherwise.
func (o *Elasticsearch) GetAwsSecret() string {
	if o == nil || o.AwsSecret == nil {
		var ret string
		return ret
	}
	return *o.AwsSecret
}

// GetAwsSecretOk returns a tuple with the AwsSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elasticsearch) GetAwsSecretOk() (*string, bool) {
	if o == nil || o.AwsSecret == nil {
		return nil, false
	}
	return o.AwsSecret, true
}

// HasAwsSecret returns a boolean if a field has been set.
func (o *Elasticsearch) HasAwsSecret() bool {
	if o != nil && o.AwsSecret != nil {
		return true
	}

	return false
}

// SetAwsSecret gets a reference to the given string and assigns it to the AwsSecret field.
func (o *Elasticsearch) SetAwsSecret(v string) {
	o.AwsSecret = &v
}

// GetAwsRegion returns the AwsRegion field value if set, zero value otherwise.
func (o *Elasticsearch) GetAwsRegion() string {
	if o == nil || o.AwsRegion == nil {
		var ret string
		return ret
	}
	return *o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elasticsearch) GetAwsRegionOk() (*string, bool) {
	if o == nil || o.AwsRegion == nil {
		return nil, false
	}
	return o.AwsRegion, true
}

// HasAwsRegion returns a boolean if a field has been set.
func (o *Elasticsearch) HasAwsRegion() bool {
	if o != nil && o.AwsRegion != nil {
		return true
	}

	return false
}

// SetAwsRegion gets a reference to the given string and assigns it to the AwsRegion field.
func (o *Elasticsearch) SetAwsRegion(v string) {
	o.AwsRegion = &v
}

// GetUseInstanceCredentials returns the UseInstanceCredentials field value if set, zero value otherwise.
func (o *Elasticsearch) GetUseInstanceCredentials() bool {
	if o == nil || o.UseInstanceCredentials == nil {
		var ret bool
		return ret
	}
	return *o.UseInstanceCredentials
}

// GetUseInstanceCredentialsOk returns a tuple with the UseInstanceCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elasticsearch) GetUseInstanceCredentialsOk() (*bool, bool) {
	if o == nil || o.UseInstanceCredentials == nil {
		return nil, false
	}
	return o.UseInstanceCredentials, true
}

// HasUseInstanceCredentials returns a boolean if a field has been set.
func (o *Elasticsearch) HasUseInstanceCredentials() bool {
	if o != nil && o.UseInstanceCredentials != nil {
		return true
	}

	return false
}

// SetUseInstanceCredentials gets a reference to the given bool and assigns it to the UseInstanceCredentials field.
func (o *Elasticsearch) SetUseInstanceCredentials(v bool) {
	o.UseInstanceCredentials = &v
}

// GetUrl returns the Url field value
func (o *Elasticsearch) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Elasticsearch) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Elasticsearch) SetUrl(v string) {
	o.Url = v
}

// GetRetentionDays returns the RetentionDays field value if set, zero value otherwise.
func (o *Elasticsearch) GetRetentionDays() int32 {
	if o == nil || o.RetentionDays == nil {
		var ret int32
		return ret
	}
	return *o.RetentionDays
}

// GetRetentionDaysOk returns a tuple with the RetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elasticsearch) GetRetentionDaysOk() (*int32, bool) {
	if o == nil || o.RetentionDays == nil {
		return nil, false
	}
	return o.RetentionDays, true
}

// HasRetentionDays returns a boolean if a field has been set.
func (o *Elasticsearch) HasRetentionDays() bool {
	if o != nil && o.RetentionDays != nil {
		return true
	}

	return false
}

// SetRetentionDays gets a reference to the given int32 and assigns it to the RetentionDays field.
func (o *Elasticsearch) SetRetentionDays(v int32) {
	o.RetentionDays = &v
}

// GetCompatibilityMode returns the CompatibilityMode field value if set, zero value otherwise.
func (o *Elasticsearch) GetCompatibilityMode() int32 {
	if o == nil || o.CompatibilityMode == nil {
		var ret int32
		return ret
	}
	return *o.CompatibilityMode
}

// GetCompatibilityModeOk returns a tuple with the CompatibilityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elasticsearch) GetCompatibilityModeOk() (*int32, bool) {
	if o == nil || o.CompatibilityMode == nil {
		return nil, false
	}
	return o.CompatibilityMode, true
}

// HasCompatibilityMode returns a boolean if a field has been set.
func (o *Elasticsearch) HasCompatibilityMode() bool {
	if o != nil && o.CompatibilityMode != nil {
		return true
	}

	return false
}

// SetCompatibilityMode gets a reference to the given int32 and assigns it to the CompatibilityMode field.
func (o *Elasticsearch) SetCompatibilityMode(v int32) {
	o.CompatibilityMode = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *Elasticsearch) GetAuthentication() ElasticsearchAllOfAuthentication {
	if o == nil || o.Authentication == nil {
		var ret ElasticsearchAllOfAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Elasticsearch) GetAuthenticationOk() (*ElasticsearchAllOfAuthentication, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *Elasticsearch) HasAuthentication() bool {
	if o != nil && o.Authentication != nil {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given ElasticsearchAllOfAuthentication and assigns it to the Authentication field.
func (o *Elasticsearch) SetAuthentication(v ElasticsearchAllOfAuthentication) {
	o.Authentication = &v
}

func (o Elasticsearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsId != nil {
		toSerialize["awsId"] = o.AwsId
	}
	if o.AwsSecret != nil {
		toSerialize["awsSecret"] = o.AwsSecret
	}
	if o.AwsRegion != nil {
		toSerialize["awsRegion"] = o.AwsRegion
	}
	if o.UseInstanceCredentials != nil {
		toSerialize["useInstanceCredentials"] = o.UseInstanceCredentials
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.RetentionDays != nil {
		toSerialize["retentionDays"] = o.RetentionDays
	}
	if o.CompatibilityMode != nil {
		toSerialize["compatibilityMode"] = o.CompatibilityMode
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableElasticsearch struct {
	value *Elasticsearch
	isSet bool
}

func (v NullableElasticsearch) Get() *Elasticsearch {
	return v.value
}

func (v *NullableElasticsearch) Set(val *Elasticsearch) {
	v.value = val
	v.isSet = true
}

func (v NullableElasticsearch) IsSet() bool {
	return v.isSet
}

func (v *NullableElasticsearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElasticsearch(val *Elasticsearch) *NullableElasticsearch {
	return &NullableElasticsearch{value: val, isSet: true}
}

func (v NullableElasticsearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElasticsearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
