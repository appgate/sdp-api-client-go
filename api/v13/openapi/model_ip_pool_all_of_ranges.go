/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v13+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 13
 * Contact: appgatesdp.support@appgate.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IpPoolAllOfRanges Represents an IP range. Either \"cidr\" or \"first-last\" format can be used.
type IpPoolAllOfRanges struct {
	// IP subnet in CIDR format.
	Cidr *string `json:"cidr,omitempty"`
	// The beginning of the IP range.
	First *string `json:"first,omitempty"`
	// The end of the IP range.
	Last *string `json:"last,omitempty"`
}

// NewIpPoolAllOfRanges instantiates a new IpPoolAllOfRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpPoolAllOfRanges() *IpPoolAllOfRanges {
	this := IpPoolAllOfRanges{}
	return &this
}

// NewIpPoolAllOfRangesWithDefaults instantiates a new IpPoolAllOfRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpPoolAllOfRangesWithDefaults() *IpPoolAllOfRanges {
	this := IpPoolAllOfRanges{}
	return &this
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *IpPoolAllOfRanges) GetCidr() string {
	if o == nil || o.Cidr == nil {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolAllOfRanges) GetCidrOk() (*string, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *IpPoolAllOfRanges) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *IpPoolAllOfRanges) SetCidr(v string) {
	o.Cidr = &v
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *IpPoolAllOfRanges) GetFirst() string {
	if o == nil || o.First == nil {
		var ret string
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolAllOfRanges) GetFirstOk() (*string, bool) {
	if o == nil || o.First == nil {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *IpPoolAllOfRanges) HasFirst() bool {
	if o != nil && o.First != nil {
		return true
	}

	return false
}

// SetFirst gets a reference to the given string and assigns it to the First field.
func (o *IpPoolAllOfRanges) SetFirst(v string) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *IpPoolAllOfRanges) GetLast() string {
	if o == nil || o.Last == nil {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolAllOfRanges) GetLastOk() (*string, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *IpPoolAllOfRanges) HasLast() bool {
	if o != nil && o.Last != nil {
		return true
	}

	return false
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *IpPoolAllOfRanges) SetLast(v string) {
	o.Last = &v
}

func (o IpPoolAllOfRanges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	if o.First != nil {
		toSerialize["first"] = o.First
	}
	if o.Last != nil {
		toSerialize["last"] = o.Last
	}
	return json.Marshal(toSerialize)
}

type NullableIpPoolAllOfRanges struct {
	value *IpPoolAllOfRanges
	isSet bool
}

func (v NullableIpPoolAllOfRanges) Get() *IpPoolAllOfRanges {
	return v.value
}

func (v *NullableIpPoolAllOfRanges) Set(val *IpPoolAllOfRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableIpPoolAllOfRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableIpPoolAllOfRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpPoolAllOfRanges(val *IpPoolAllOfRanges) *NullableIpPoolAllOfRanges {
	return &NullableIpPoolAllOfRanges{value: val, isSet: true}
}

func (v NullableIpPoolAllOfRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpPoolAllOfRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
