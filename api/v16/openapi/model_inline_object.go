/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.5
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InlineObject struct for InlineObject
type InlineObject struct {
	// Whether the Appliance Backup should include syslog or not.
	Logs *bool `json:"logs,omitempty"`
	// Whether the Appliance Backup should include the audit logs or not.
	Audit *bool `json:"audit,omitempty"`
	// URL to send a POST request to notify about the completion of the backup, that it can now be downloaded and deleted.
	NotifyUrl *string `json:"notifyUrl,omitempty"`
}

// NewInlineObject instantiates a new InlineObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject() *InlineObject {
	this := InlineObject{}
	return &this
}

// NewInlineObjectWithDefaults instantiates a new InlineObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObjectWithDefaults() *InlineObject {
	this := InlineObject{}
	return &this
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *InlineObject) GetLogs() bool {
	if o == nil || o.Logs == nil {
		var ret bool
		return ret
	}
	return *o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetLogsOk() (*bool, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *InlineObject) HasLogs() bool {
	if o != nil && o.Logs != nil {
		return true
	}

	return false
}

// SetLogs gets a reference to the given bool and assigns it to the Logs field.
func (o *InlineObject) SetLogs(v bool) {
	o.Logs = &v
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *InlineObject) GetAudit() bool {
	if o == nil || o.Audit == nil {
		var ret bool
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetAuditOk() (*bool, bool) {
	if o == nil || o.Audit == nil {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *InlineObject) HasAudit() bool {
	if o != nil && o.Audit != nil {
		return true
	}

	return false
}

// SetAudit gets a reference to the given bool and assigns it to the Audit field.
func (o *InlineObject) SetAudit(v bool) {
	o.Audit = &v
}

// GetNotifyUrl returns the NotifyUrl field value if set, zero value otherwise.
func (o *InlineObject) GetNotifyUrl() string {
	if o == nil || o.NotifyUrl == nil {
		var ret string
		return ret
	}
	return *o.NotifyUrl
}

// GetNotifyUrlOk returns a tuple with the NotifyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetNotifyUrlOk() (*string, bool) {
	if o == nil || o.NotifyUrl == nil {
		return nil, false
	}
	return o.NotifyUrl, true
}

// HasNotifyUrl returns a boolean if a field has been set.
func (o *InlineObject) HasNotifyUrl() bool {
	if o != nil && o.NotifyUrl != nil {
		return true
	}

	return false
}

// SetNotifyUrl gets a reference to the given string and assigns it to the NotifyUrl field.
func (o *InlineObject) SetNotifyUrl(v string) {
	o.NotifyUrl = &v
}

func (o InlineObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Logs != nil {
		toSerialize["logs"] = o.Logs
	}
	if o.Audit != nil {
		toSerialize["audit"] = o.Audit
	}
	if o.NotifyUrl != nil {
		toSerialize["notifyUrl"] = o.NotifyUrl
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject struct {
	value *InlineObject
	isSet bool
}

func (v NullableInlineObject) Get() *InlineObject {
	return v.value
}

func (v *NullableInlineObject) Set(val *InlineObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject(val *InlineObject) *NullableInlineObject {
	return &NullableInlineObject{value: val, isSet: true}
}

func (v NullableInlineObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
