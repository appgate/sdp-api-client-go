/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AwsKinesis struct for AwsKinesis
type AwsKinesis struct {
	// AWS ID to login. Only required if AWS Access Keys are being used to authenticate.
	AwsId *string `json:"awsId,omitempty"`
	// AWS secret to login. Only required if AWS Access Keys are being used to authenticate.
	AwsSecret *string `json:"awsSecret,omitempty"`
	// AWS region. Only required if AWS Access Keys are being used to authenticate.
	AwsRegion *string `json:"awsRegion,omitempty"`
	// Whether to use the credentials from the AWS instance or not.
	UseInstanceCredentials *bool `json:"useInstanceCredentials,omitempty"`
	// AWS Kinesis type
	Type string `json:"type"`
	// Name of the stream.
	StreamName string `json:"streamName"`
	// Batch size for the stream. Used only for \"Stream\" type.
	BatchSize *int32 `json:"batchSize,omitempty"`
	// Number of partition keys to use for the stream. Used only for \"Stream\" type.
	NumberOfPartitionKeys *int32 `json:"numberOfPartitionKeys,omitempty"`
	// JMESPath expression to filter audit logs to forward.
	Filter *string `json:"filter,omitempty"`
}

// NewAwsKinesis instantiates a new AwsKinesis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsKinesis(type_ string, streamName string) *AwsKinesis {
	this := AwsKinesis{}
	this.Type = type_
	this.StreamName = streamName
	var batchSize int32 = 400
	this.BatchSize = &batchSize
	var numberOfPartitionKeys int32 = 10
	this.NumberOfPartitionKeys = &numberOfPartitionKeys
	return &this
}

// NewAwsKinesisWithDefaults instantiates a new AwsKinesis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsKinesisWithDefaults() *AwsKinesis {
	this := AwsKinesis{}
	var batchSize int32 = 400
	this.BatchSize = &batchSize
	var numberOfPartitionKeys int32 = 10
	this.NumberOfPartitionKeys = &numberOfPartitionKeys
	return &this
}

// GetAwsId returns the AwsId field value if set, zero value otherwise.
func (o *AwsKinesis) GetAwsId() string {
	if o == nil || o.AwsId == nil {
		var ret string
		return ret
	}
	return *o.AwsId
}

// GetAwsIdOk returns a tuple with the AwsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsKinesis) GetAwsIdOk() (*string, bool) {
	if o == nil || o.AwsId == nil {
		return nil, false
	}
	return o.AwsId, true
}

// HasAwsId returns a boolean if a field has been set.
func (o *AwsKinesis) HasAwsId() bool {
	if o != nil && o.AwsId != nil {
		return true
	}

	return false
}

// SetAwsId gets a reference to the given string and assigns it to the AwsId field.
func (o *AwsKinesis) SetAwsId(v string) {
	o.AwsId = &v
}

// GetAwsSecret returns the AwsSecret field value if set, zero value otherwise.
func (o *AwsKinesis) GetAwsSecret() string {
	if o == nil || o.AwsSecret == nil {
		var ret string
		return ret
	}
	return *o.AwsSecret
}

// GetAwsSecretOk returns a tuple with the AwsSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsKinesis) GetAwsSecretOk() (*string, bool) {
	if o == nil || o.AwsSecret == nil {
		return nil, false
	}
	return o.AwsSecret, true
}

// HasAwsSecret returns a boolean if a field has been set.
func (o *AwsKinesis) HasAwsSecret() bool {
	if o != nil && o.AwsSecret != nil {
		return true
	}

	return false
}

// SetAwsSecret gets a reference to the given string and assigns it to the AwsSecret field.
func (o *AwsKinesis) SetAwsSecret(v string) {
	o.AwsSecret = &v
}

// GetAwsRegion returns the AwsRegion field value if set, zero value otherwise.
func (o *AwsKinesis) GetAwsRegion() string {
	if o == nil || o.AwsRegion == nil {
		var ret string
		return ret
	}
	return *o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsKinesis) GetAwsRegionOk() (*string, bool) {
	if o == nil || o.AwsRegion == nil {
		return nil, false
	}
	return o.AwsRegion, true
}

// HasAwsRegion returns a boolean if a field has been set.
func (o *AwsKinesis) HasAwsRegion() bool {
	if o != nil && o.AwsRegion != nil {
		return true
	}

	return false
}

// SetAwsRegion gets a reference to the given string and assigns it to the AwsRegion field.
func (o *AwsKinesis) SetAwsRegion(v string) {
	o.AwsRegion = &v
}

// GetUseInstanceCredentials returns the UseInstanceCredentials field value if set, zero value otherwise.
func (o *AwsKinesis) GetUseInstanceCredentials() bool {
	if o == nil || o.UseInstanceCredentials == nil {
		var ret bool
		return ret
	}
	return *o.UseInstanceCredentials
}

// GetUseInstanceCredentialsOk returns a tuple with the UseInstanceCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsKinesis) GetUseInstanceCredentialsOk() (*bool, bool) {
	if o == nil || o.UseInstanceCredentials == nil {
		return nil, false
	}
	return o.UseInstanceCredentials, true
}

// HasUseInstanceCredentials returns a boolean if a field has been set.
func (o *AwsKinesis) HasUseInstanceCredentials() bool {
	if o != nil && o.UseInstanceCredentials != nil {
		return true
	}

	return false
}

// SetUseInstanceCredentials gets a reference to the given bool and assigns it to the UseInstanceCredentials field.
func (o *AwsKinesis) SetUseInstanceCredentials(v bool) {
	o.UseInstanceCredentials = &v
}

// GetType returns the Type field value
func (o *AwsKinesis) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AwsKinesis) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AwsKinesis) SetType(v string) {
	o.Type = v
}

// GetStreamName returns the StreamName field value
func (o *AwsKinesis) GetStreamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value
// and a boolean to check if the value has been set.
func (o *AwsKinesis) GetStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamName, true
}

// SetStreamName sets field value
func (o *AwsKinesis) SetStreamName(v string) {
	o.StreamName = v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *AwsKinesis) GetBatchSize() int32 {
	if o == nil || o.BatchSize == nil {
		var ret int32
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsKinesis) GetBatchSizeOk() (*int32, bool) {
	if o == nil || o.BatchSize == nil {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *AwsKinesis) HasBatchSize() bool {
	if o != nil && o.BatchSize != nil {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given int32 and assigns it to the BatchSize field.
func (o *AwsKinesis) SetBatchSize(v int32) {
	o.BatchSize = &v
}

// GetNumberOfPartitionKeys returns the NumberOfPartitionKeys field value if set, zero value otherwise.
func (o *AwsKinesis) GetNumberOfPartitionKeys() int32 {
	if o == nil || o.NumberOfPartitionKeys == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfPartitionKeys
}

// GetNumberOfPartitionKeysOk returns a tuple with the NumberOfPartitionKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsKinesis) GetNumberOfPartitionKeysOk() (*int32, bool) {
	if o == nil || o.NumberOfPartitionKeys == nil {
		return nil, false
	}
	return o.NumberOfPartitionKeys, true
}

// HasNumberOfPartitionKeys returns a boolean if a field has been set.
func (o *AwsKinesis) HasNumberOfPartitionKeys() bool {
	if o != nil && o.NumberOfPartitionKeys != nil {
		return true
	}

	return false
}

// SetNumberOfPartitionKeys gets a reference to the given int32 and assigns it to the NumberOfPartitionKeys field.
func (o *AwsKinesis) SetNumberOfPartitionKeys(v int32) {
	o.NumberOfPartitionKeys = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AwsKinesis) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsKinesis) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AwsKinesis) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *AwsKinesis) SetFilter(v string) {
	o.Filter = &v
}

func (o AwsKinesis) MarshalJSON() ([]byte, error) {
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
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["streamName"] = o.StreamName
	}
	if o.BatchSize != nil {
		toSerialize["batchSize"] = o.BatchSize
	}
	if o.NumberOfPartitionKeys != nil {
		toSerialize["numberOfPartitionKeys"] = o.NumberOfPartitionKeys
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	return json.Marshal(toSerialize)
}

type NullableAwsKinesis struct {
	value *AwsKinesis
	isSet bool
}

func (v NullableAwsKinesis) Get() *AwsKinesis {
	return v.value
}

func (v *NullableAwsKinesis) Set(val *AwsKinesis) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsKinesis) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsKinesis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsKinesis(val *AwsKinesis) *NullableAwsKinesis {
	return &NullableAwsKinesis{value: val, isSet: true}
}

func (v NullableAwsKinesis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsKinesis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
