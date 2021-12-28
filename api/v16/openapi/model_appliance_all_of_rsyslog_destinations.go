/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOfRsyslogDestinations struct for ApplianceAllOfRsyslogDestinations
type ApplianceAllOfRsyslogDestinations struct {
	// Rsyslog selector.
	Selector *string `json:"selector,omitempty"`
	// Rsyslog template to forward logs with.
	Template *string `json:"template,omitempty"`
	// Rsyslog server destination.
	Destination string `json:"destination"`
}

// NewApplianceAllOfRsyslogDestinations instantiates a new ApplianceAllOfRsyslogDestinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfRsyslogDestinations(destination string) *ApplianceAllOfRsyslogDestinations {
	this := ApplianceAllOfRsyslogDestinations{}
	var selector string = "*.*"
	this.Selector = &selector
	var template string = "%HOSTNAME% %msg%"
	this.Template = &template
	this.Destination = destination
	return &this
}

// NewApplianceAllOfRsyslogDestinationsWithDefaults instantiates a new ApplianceAllOfRsyslogDestinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfRsyslogDestinationsWithDefaults() *ApplianceAllOfRsyslogDestinations {
	this := ApplianceAllOfRsyslogDestinations{}
	var selector string = "*.*"
	this.Selector = &selector
	var template string = "%HOSTNAME% %msg%"
	this.Template = &template
	return &this
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *ApplianceAllOfRsyslogDestinations) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfRsyslogDestinations) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *ApplianceAllOfRsyslogDestinations) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *ApplianceAllOfRsyslogDestinations) SetSelector(v string) {
	o.Selector = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ApplianceAllOfRsyslogDestinations) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfRsyslogDestinations) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ApplianceAllOfRsyslogDestinations) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *ApplianceAllOfRsyslogDestinations) SetTemplate(v string) {
	o.Template = &v
}

// GetDestination returns the Destination field value
func (o *ApplianceAllOfRsyslogDestinations) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfRsyslogDestinations) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *ApplianceAllOfRsyslogDestinations) SetDestination(v string) {
	o.Destination = v
}

func (o ApplianceAllOfRsyslogDestinations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfRsyslogDestinations struct {
	value *ApplianceAllOfRsyslogDestinations
	isSet bool
}

func (v NullableApplianceAllOfRsyslogDestinations) Get() *ApplianceAllOfRsyslogDestinations {
	return v.value
}

func (v *NullableApplianceAllOfRsyslogDestinations) Set(val *ApplianceAllOfRsyslogDestinations) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfRsyslogDestinations) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfRsyslogDestinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfRsyslogDestinations(val *ApplianceAllOfRsyslogDestinations) *NullableApplianceAllOfRsyslogDestinations {
	return &NullableApplianceAllOfRsyslogDestinations{value: val, isSet: true}
}

func (v NullableApplianceAllOfRsyslogDestinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfRsyslogDestinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
