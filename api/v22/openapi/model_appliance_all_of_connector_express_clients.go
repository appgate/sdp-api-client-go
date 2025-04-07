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

// ApplianceAllOfConnectorExpressClients struct for ApplianceAllOfConnectorExpressClients
type ApplianceAllOfConnectorExpressClients struct {
	// Name for the Client. It will be mapped to the user claim 'clientName'.
	Name string `json:"name"`
	// The device ID to assign to this Client. It will be used to generate device distinguished name.
	DeviceId *string `json:"deviceId,omitempty"`
	// A list of subnets to allow access.
	AllowResources []AllowResourcesInner `json:"allowResources,omitempty"`
	// Use SNAT for outgoing traffic from the Express Connector, endpoints will see traffic as coming from the Connector itself
	SnatToResources *bool `json:"snatToResources,omitempty"`
	// Apply destination NAT to traffic from tunnel into a resource
	DnatToResource *bool `json:"dnatToResource,omitempty"`
}

// NewApplianceAllOfConnectorExpressClients instantiates a new ApplianceAllOfConnectorExpressClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfConnectorExpressClients(name string) *ApplianceAllOfConnectorExpressClients {
	this := ApplianceAllOfConnectorExpressClients{}
	this.Name = name
	var snatToResources bool = true
	this.SnatToResources = &snatToResources
	var dnatToResource bool = false
	this.DnatToResource = &dnatToResource
	return &this
}

// NewApplianceAllOfConnectorExpressClientsWithDefaults instantiates a new ApplianceAllOfConnectorExpressClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfConnectorExpressClientsWithDefaults() *ApplianceAllOfConnectorExpressClients {
	this := ApplianceAllOfConnectorExpressClients{}
	var snatToResources bool = true
	this.SnatToResources = &snatToResources
	var dnatToResource bool = false
	this.DnatToResource = &dnatToResource
	return &this
}

// GetName returns the Name field value
func (o *ApplianceAllOfConnectorExpressClients) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorExpressClients) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplianceAllOfConnectorExpressClients) SetName(v string) {
	o.Name = v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ApplianceAllOfConnectorExpressClients) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorExpressClients) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ApplianceAllOfConnectorExpressClients) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ApplianceAllOfConnectorExpressClients) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetAllowResources returns the AllowResources field value if set, zero value otherwise.
func (o *ApplianceAllOfConnectorExpressClients) GetAllowResources() []AllowResourcesInner {
	if o == nil || o.AllowResources == nil {
		var ret []AllowResourcesInner
		return ret
	}
	return o.AllowResources
}

// GetAllowResourcesOk returns a tuple with the AllowResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorExpressClients) GetAllowResourcesOk() ([]AllowResourcesInner, bool) {
	if o == nil || o.AllowResources == nil {
		return nil, false
	}
	return o.AllowResources, true
}

// HasAllowResources returns a boolean if a field has been set.
func (o *ApplianceAllOfConnectorExpressClients) HasAllowResources() bool {
	if o != nil && o.AllowResources != nil {
		return true
	}

	return false
}

// SetAllowResources gets a reference to the given []AllowResourcesInner and assigns it to the AllowResources field.
func (o *ApplianceAllOfConnectorExpressClients) SetAllowResources(v []AllowResourcesInner) {
	o.AllowResources = v
}

// GetSnatToResources returns the SnatToResources field value if set, zero value otherwise.
func (o *ApplianceAllOfConnectorExpressClients) GetSnatToResources() bool {
	if o == nil || o.SnatToResources == nil {
		var ret bool
		return ret
	}
	return *o.SnatToResources
}

// GetSnatToResourcesOk returns a tuple with the SnatToResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorExpressClients) GetSnatToResourcesOk() (*bool, bool) {
	if o == nil || o.SnatToResources == nil {
		return nil, false
	}
	return o.SnatToResources, true
}

// HasSnatToResources returns a boolean if a field has been set.
func (o *ApplianceAllOfConnectorExpressClients) HasSnatToResources() bool {
	if o != nil && o.SnatToResources != nil {
		return true
	}

	return false
}

// SetSnatToResources gets a reference to the given bool and assigns it to the SnatToResources field.
func (o *ApplianceAllOfConnectorExpressClients) SetSnatToResources(v bool) {
	o.SnatToResources = &v
}

// GetDnatToResource returns the DnatToResource field value if set, zero value otherwise.
func (o *ApplianceAllOfConnectorExpressClients) GetDnatToResource() bool {
	if o == nil || o.DnatToResource == nil {
		var ret bool
		return ret
	}
	return *o.DnatToResource
}

// GetDnatToResourceOk returns a tuple with the DnatToResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfConnectorExpressClients) GetDnatToResourceOk() (*bool, bool) {
	if o == nil || o.DnatToResource == nil {
		return nil, false
	}
	return o.DnatToResource, true
}

// HasDnatToResource returns a boolean if a field has been set.
func (o *ApplianceAllOfConnectorExpressClients) HasDnatToResource() bool {
	if o != nil && o.DnatToResource != nil {
		return true
	}

	return false
}

// SetDnatToResource gets a reference to the given bool and assigns it to the DnatToResource field.
func (o *ApplianceAllOfConnectorExpressClients) SetDnatToResource(v bool) {
	o.DnatToResource = &v
}

func (o ApplianceAllOfConnectorExpressClients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.AllowResources != nil {
		toSerialize["allowResources"] = o.AllowResources
	}
	if o.SnatToResources != nil {
		toSerialize["snatToResources"] = o.SnatToResources
	}
	if o.DnatToResource != nil {
		toSerialize["dnatToResource"] = o.DnatToResource
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfConnectorExpressClients struct {
	value *ApplianceAllOfConnectorExpressClients
	isSet bool
}

func (v NullableApplianceAllOfConnectorExpressClients) Get() *ApplianceAllOfConnectorExpressClients {
	return v.value
}

func (v *NullableApplianceAllOfConnectorExpressClients) Set(val *ApplianceAllOfConnectorExpressClients) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfConnectorExpressClients) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfConnectorExpressClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfConnectorExpressClients(val *ApplianceAllOfConnectorExpressClients) *NullableApplianceAllOfConnectorExpressClients {
	return &NullableApplianceAllOfConnectorExpressClients{value: val, isSet: true}
}

func (v NullableApplianceAllOfConnectorExpressClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfConnectorExpressClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
