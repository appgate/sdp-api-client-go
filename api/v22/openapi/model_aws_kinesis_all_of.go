/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AwsKinesisAllOf struct for AwsKinesisAllOf
type AwsKinesisAllOf struct {
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

// NewAwsKinesisAllOf instantiates a new AwsKinesisAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsKinesisAllOf(type_ string, streamName string) *AwsKinesisAllOf {
	this := AwsKinesisAllOf{}
	this.Type = type_
	this.StreamName = streamName
	var batchSize int32 = 400
	this.BatchSize = &batchSize
	var numberOfPartitionKeys int32 = 10
	this.NumberOfPartitionKeys = &numberOfPartitionKeys
	return &this
}

// NewAwsKinesisAllOfWithDefaults instantiates a new AwsKinesisAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsKinesisAllOfWithDefaults() *AwsKinesisAllOf {
	this := AwsKinesisAllOf{}
	var batchSize int32 = 400
	this.BatchSize = &batchSize
	var numberOfPartitionKeys int32 = 10
	this.NumberOfPartitionKeys = &numberOfPartitionKeys
	return &this
}

// GetType returns the Type field value
func (o *AwsKinesisAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AwsKinesisAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AwsKinesisAllOf) SetType(v string) {
	o.Type = v
}

// GetStreamName returns the StreamName field value
func (o *AwsKinesisAllOf) GetStreamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value
// and a boolean to check if the value has been set.
func (o *AwsKinesisAllOf) GetStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamName, true
}

// SetStreamName sets field value
func (o *AwsKinesisAllOf) SetStreamName(v string) {
	o.StreamName = v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *AwsKinesisAllOf) GetBatchSize() int32 {
	if o == nil || o.BatchSize == nil {
		var ret int32
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsKinesisAllOf) GetBatchSizeOk() (*int32, bool) {
	if o == nil || o.BatchSize == nil {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *AwsKinesisAllOf) HasBatchSize() bool {
	if o != nil && o.BatchSize != nil {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given int32 and assigns it to the BatchSize field.
func (o *AwsKinesisAllOf) SetBatchSize(v int32) {
	o.BatchSize = &v
}

// GetNumberOfPartitionKeys returns the NumberOfPartitionKeys field value if set, zero value otherwise.
func (o *AwsKinesisAllOf) GetNumberOfPartitionKeys() int32 {
	if o == nil || o.NumberOfPartitionKeys == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfPartitionKeys
}

// GetNumberOfPartitionKeysOk returns a tuple with the NumberOfPartitionKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsKinesisAllOf) GetNumberOfPartitionKeysOk() (*int32, bool) {
	if o == nil || o.NumberOfPartitionKeys == nil {
		return nil, false
	}
	return o.NumberOfPartitionKeys, true
}

// HasNumberOfPartitionKeys returns a boolean if a field has been set.
func (o *AwsKinesisAllOf) HasNumberOfPartitionKeys() bool {
	if o != nil && o.NumberOfPartitionKeys != nil {
		return true
	}

	return false
}

// SetNumberOfPartitionKeys gets a reference to the given int32 and assigns it to the NumberOfPartitionKeys field.
func (o *AwsKinesisAllOf) SetNumberOfPartitionKeys(v int32) {
	o.NumberOfPartitionKeys = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AwsKinesisAllOf) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsKinesisAllOf) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AwsKinesisAllOf) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *AwsKinesisAllOf) SetFilter(v string) {
	o.Filter = &v
}

func (o AwsKinesisAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableAwsKinesisAllOf struct {
	value *AwsKinesisAllOf
	isSet bool
}

func (v NullableAwsKinesisAllOf) Get() *AwsKinesisAllOf {
	return v.value
}

func (v *NullableAwsKinesisAllOf) Set(val *AwsKinesisAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsKinesisAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsKinesisAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsKinesisAllOf(val *AwsKinesisAllOf) *NullableAwsKinesisAllOf {
	return &NullableAwsKinesisAllOf{value: val, isSet: true}
}

func (v NullableAwsKinesisAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsKinesisAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
