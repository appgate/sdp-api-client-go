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
	"fmt"
)

// EntitlementScriptTestResultResultInner struct for EntitlementScriptTestResultResultInner
type EntitlementScriptTestResultResultInner struct {
	AppShortcut *AppShortcut
	string      *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EntitlementScriptTestResultResultInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AppShortcut
	err = json.Unmarshal(data, &dst.AppShortcut)
	if err == nil {
		jsonAppShortcut, _ := json.Marshal(dst.AppShortcut)
		if string(jsonAppShortcut) == "{}" { // empty struct
			dst.AppShortcut = nil
		} else {
			return nil // data stored in dst.AppShortcut, return on the first match
		}
	} else {
		dst.AppShortcut = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(EntitlementScriptTestResultResultInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EntitlementScriptTestResultResultInner) MarshalJSON() ([]byte, error) {
	if src.AppShortcut != nil {
		return json.Marshal(&src.AppShortcut)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEntitlementScriptTestResultResultInner struct {
	value *EntitlementScriptTestResultResultInner
	isSet bool
}

func (v NullableEntitlementScriptTestResultResultInner) Get() *EntitlementScriptTestResultResultInner {
	return v.value
}

func (v *NullableEntitlementScriptTestResultResultInner) Set(val *EntitlementScriptTestResultResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementScriptTestResultResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementScriptTestResultResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementScriptTestResultResultInner(val *EntitlementScriptTestResultResultInner) *NullableEntitlementScriptTestResultResultInner {
	return &NullableEntitlementScriptTestResultResultInner{value: val, isSet: true}
}

func (v NullableEntitlementScriptTestResultResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementScriptTestResultResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
