/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AppliancesIdBackupBackupIdStatusGet200Response Backups status.
type AppliancesIdBackupBackupIdStatusGet200Response struct {
	// Encoding used for the output field
	Encoding *string `json:"encoding,omitempty"`
	// The output of the command. Must be checked if result is failure
	Output *string `json:"output,omitempty"`
	// Only set when status is done.
	Result *string `json:"result,omitempty"`
	// A free text field that shows what internally the backup is doing.
	Message *string `json:"message,omitempty"`
	// Current status of the Appliance Backup.
	Status *string `json:"status,omitempty"`
}

// NewAppliancesIdBackupBackupIdStatusGet200Response instantiates a new AppliancesIdBackupBackupIdStatusGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliancesIdBackupBackupIdStatusGet200Response() *AppliancesIdBackupBackupIdStatusGet200Response {
	this := AppliancesIdBackupBackupIdStatusGet200Response{}
	return &this
}

// NewAppliancesIdBackupBackupIdStatusGet200ResponseWithDefaults instantiates a new AppliancesIdBackupBackupIdStatusGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppliancesIdBackupBackupIdStatusGet200ResponseWithDefaults() *AppliancesIdBackupBackupIdStatusGet200Response {
	this := AppliancesIdBackupBackupIdStatusGet200Response{}
	return &this
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) GetEncoding() string {
	if o == nil || o.Encoding == nil {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) GetEncodingOk() (*string, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) HasEncoding() bool {
	if o != nil && o.Encoding != nil {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) SetEncoding(v string) {
	o.Encoding = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) GetOutput() string {
	if o == nil || o.Output == nil {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) GetOutputOk() (*string, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) SetOutput(v string) {
	o.Output = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) SetResult(v string) {
	o.Result = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AppliancesIdBackupBackupIdStatusGet200Response) SetStatus(v string) {
	o.Status = &v
}

func (o AppliancesIdBackupBackupIdStatusGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableAppliancesIdBackupBackupIdStatusGet200Response struct {
	value *AppliancesIdBackupBackupIdStatusGet200Response
	isSet bool
}

func (v NullableAppliancesIdBackupBackupIdStatusGet200Response) Get() *AppliancesIdBackupBackupIdStatusGet200Response {
	return v.value
}

func (v *NullableAppliancesIdBackupBackupIdStatusGet200Response) Set(val *AppliancesIdBackupBackupIdStatusGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliancesIdBackupBackupIdStatusGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliancesIdBackupBackupIdStatusGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliancesIdBackupBackupIdStatusGet200Response(val *AppliancesIdBackupBackupIdStatusGet200Response) *NullableAppliancesIdBackupBackupIdStatusGet200Response {
	return &NullableAppliancesIdBackupBackupIdStatusGet200Response{value: val, isSet: true}
}

func (v NullableAppliancesIdBackupBackupIdStatusGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliancesIdBackupBackupIdStatusGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
