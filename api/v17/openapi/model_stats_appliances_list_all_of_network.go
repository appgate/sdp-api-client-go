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

// StatsAppliancesListAllOfNetwork Current network utilization on the Appliance.
type StatsAppliancesListAllOfNetwork struct {
	// The name of the NIC that's most used.
	BusiestNic *string `json:"busiestNic,omitempty"`
	// Number of inbound packets dropped for the whole uptime.
	Dropin *float32 `json:"dropin,omitempty"`
	// Number of outbound packets dropped for the whole uptime.
	Dropout *float32 `json:"dropout,omitempty"`
	// The speed of the inbound network activity on the busiest NIC. The average is taken for the last 10 seconds.
	RxSpeed *string `json:"rxSpeed,omitempty"`
	// The speed of the outbound network activity on the busiest NIC. The average is taken for the last 10 seconds.
	TxSpeed *string `json:"txSpeed,omitempty"`
	// IPs of the Appliance per NIC.
	Ips *map[string][]string `json:"ips,omitempty"`
}

// NewStatsAppliancesListAllOfNetwork instantiates a new StatsAppliancesListAllOfNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsAppliancesListAllOfNetwork() *StatsAppliancesListAllOfNetwork {
	this := StatsAppliancesListAllOfNetwork{}
	return &this
}

// NewStatsAppliancesListAllOfNetworkWithDefaults instantiates a new StatsAppliancesListAllOfNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsAppliancesListAllOfNetworkWithDefaults() *StatsAppliancesListAllOfNetwork {
	this := StatsAppliancesListAllOfNetwork{}
	return &this
}

// GetBusiestNic returns the BusiestNic field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfNetwork) GetBusiestNic() string {
	if o == nil || o.BusiestNic == nil {
		var ret string
		return ret
	}
	return *o.BusiestNic
}

// GetBusiestNicOk returns a tuple with the BusiestNic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfNetwork) GetBusiestNicOk() (*string, bool) {
	if o == nil || o.BusiestNic == nil {
		return nil, false
	}
	return o.BusiestNic, true
}

// HasBusiestNic returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfNetwork) HasBusiestNic() bool {
	if o != nil && o.BusiestNic != nil {
		return true
	}

	return false
}

// SetBusiestNic gets a reference to the given string and assigns it to the BusiestNic field.
func (o *StatsAppliancesListAllOfNetwork) SetBusiestNic(v string) {
	o.BusiestNic = &v
}

// GetDropin returns the Dropin field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfNetwork) GetDropin() float32 {
	if o == nil || o.Dropin == nil {
		var ret float32
		return ret
	}
	return *o.Dropin
}

// GetDropinOk returns a tuple with the Dropin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfNetwork) GetDropinOk() (*float32, bool) {
	if o == nil || o.Dropin == nil {
		return nil, false
	}
	return o.Dropin, true
}

// HasDropin returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfNetwork) HasDropin() bool {
	if o != nil && o.Dropin != nil {
		return true
	}

	return false
}

// SetDropin gets a reference to the given float32 and assigns it to the Dropin field.
func (o *StatsAppliancesListAllOfNetwork) SetDropin(v float32) {
	o.Dropin = &v
}

// GetDropout returns the Dropout field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfNetwork) GetDropout() float32 {
	if o == nil || o.Dropout == nil {
		var ret float32
		return ret
	}
	return *o.Dropout
}

// GetDropoutOk returns a tuple with the Dropout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfNetwork) GetDropoutOk() (*float32, bool) {
	if o == nil || o.Dropout == nil {
		return nil, false
	}
	return o.Dropout, true
}

// HasDropout returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfNetwork) HasDropout() bool {
	if o != nil && o.Dropout != nil {
		return true
	}

	return false
}

// SetDropout gets a reference to the given float32 and assigns it to the Dropout field.
func (o *StatsAppliancesListAllOfNetwork) SetDropout(v float32) {
	o.Dropout = &v
}

// GetRxSpeed returns the RxSpeed field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfNetwork) GetRxSpeed() string {
	if o == nil || o.RxSpeed == nil {
		var ret string
		return ret
	}
	return *o.RxSpeed
}

// GetRxSpeedOk returns a tuple with the RxSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfNetwork) GetRxSpeedOk() (*string, bool) {
	if o == nil || o.RxSpeed == nil {
		return nil, false
	}
	return o.RxSpeed, true
}

// HasRxSpeed returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfNetwork) HasRxSpeed() bool {
	if o != nil && o.RxSpeed != nil {
		return true
	}

	return false
}

// SetRxSpeed gets a reference to the given string and assigns it to the RxSpeed field.
func (o *StatsAppliancesListAllOfNetwork) SetRxSpeed(v string) {
	o.RxSpeed = &v
}

// GetTxSpeed returns the TxSpeed field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfNetwork) GetTxSpeed() string {
	if o == nil || o.TxSpeed == nil {
		var ret string
		return ret
	}
	return *o.TxSpeed
}

// GetTxSpeedOk returns a tuple with the TxSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfNetwork) GetTxSpeedOk() (*string, bool) {
	if o == nil || o.TxSpeed == nil {
		return nil, false
	}
	return o.TxSpeed, true
}

// HasTxSpeed returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfNetwork) HasTxSpeed() bool {
	if o != nil && o.TxSpeed != nil {
		return true
	}

	return false
}

// SetTxSpeed gets a reference to the given string and assigns it to the TxSpeed field.
func (o *StatsAppliancesListAllOfNetwork) SetTxSpeed(v string) {
	o.TxSpeed = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfNetwork) GetIps() map[string][]string {
	if o == nil || o.Ips == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfNetwork) GetIpsOk() (*map[string][]string, bool) {
	if o == nil || o.Ips == nil {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfNetwork) HasIps() bool {
	if o != nil && o.Ips != nil {
		return true
	}

	return false
}

// SetIps gets a reference to the given map[string][]string and assigns it to the Ips field.
func (o *StatsAppliancesListAllOfNetwork) SetIps(v map[string][]string) {
	o.Ips = &v
}

func (o StatsAppliancesListAllOfNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BusiestNic != nil {
		toSerialize["busiestNic"] = o.BusiestNic
	}
	if o.Dropin != nil {
		toSerialize["dropin"] = o.Dropin
	}
	if o.Dropout != nil {
		toSerialize["dropout"] = o.Dropout
	}
	if o.RxSpeed != nil {
		toSerialize["rxSpeed"] = o.RxSpeed
	}
	if o.TxSpeed != nil {
		toSerialize["txSpeed"] = o.TxSpeed
	}
	if o.Ips != nil {
		toSerialize["ips"] = o.Ips
	}
	return json.Marshal(toSerialize)
}

type NullableStatsAppliancesListAllOfNetwork struct {
	value *StatsAppliancesListAllOfNetwork
	isSet bool
}

func (v NullableStatsAppliancesListAllOfNetwork) Get() *StatsAppliancesListAllOfNetwork {
	return v.value
}

func (v *NullableStatsAppliancesListAllOfNetwork) Set(val *StatsAppliancesListAllOfNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsAppliancesListAllOfNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsAppliancesListAllOfNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsAppliancesListAllOfNetwork(val *StatsAppliancesListAllOfNetwork) *NullableStatsAppliancesListAllOfNetwork {
	return &NullableStatsAppliancesListAllOfNetwork{value: val, isSet: true}
}

func (v NullableStatsAppliancesListAllOfNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsAppliancesListAllOfNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
