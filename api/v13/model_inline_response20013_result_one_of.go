/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v13+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 13
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"fmt"
)

// InlineResponse20013ResultOneOf struct for InlineResponse20013ResultOneOf
type InlineResponse20013ResultOneOf struct {
	InlineResponse20013ResultOneOfInterface interface{}
}

func (s InlineResponse20013ResultOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.InlineResponse20013ResultOneOfInterface)
}

func (s *InlineResponse20013ResultOneOf) UnmarshalJSON(src []byte) error {
	var err error
	var unmarshaledAppShortcut *AppShortcut = &AppShortcut{}
	err = json.Unmarshal(src, unmarshaledAppShortcut)
	if err == nil {
		s.InlineResponse20013ResultOneOfInterface = unmarshaledAppShortcut
		return nil
	}
	return fmt.Errorf("No oneOf model could be deserialized from payload: %s", string(src))
}

type NullableInlineResponse20013ResultOneOf struct {
	value *InlineResponse20013ResultOneOf
	isSet bool
}

func (v NullableInlineResponse20013ResultOneOf) Get() *InlineResponse20013ResultOneOf {
	return v.value
}

func (v *NullableInlineResponse20013ResultOneOf) Set(val *InlineResponse20013ResultOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20013ResultOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20013ResultOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20013ResultOneOf(val *InlineResponse20013ResultOneOf) *NullableInlineResponse20013ResultOneOf {
	return &NullableInlineResponse20013ResultOneOf{value: val, isSet: true}
}

func (v NullableInlineResponse20013ResultOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20013ResultOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
