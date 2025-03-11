/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v21+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 21.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ResolverResourcesAllOf Resource query result.
type ResolverResourcesAllOf struct {
	// List of queries resources.
	Data     []string      `json:"data,omitempty"`
	Resolver *ResolverType `json:"resolver,omitempty"`
	Type     *ResourceType `json:"type,omitempty"`
	// The name of the Gateway which returned the result.
	GatewayName *string `json:"gatewayName,omitempty"`
}

// NewResolverResourcesAllOf instantiates a new ResolverResourcesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolverResourcesAllOf() *ResolverResourcesAllOf {
	this := ResolverResourcesAllOf{}
	return &this
}

// NewResolverResourcesAllOfWithDefaults instantiates a new ResolverResourcesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolverResourcesAllOfWithDefaults() *ResolverResourcesAllOf {
	this := ResolverResourcesAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResolverResourcesAllOf) GetData() []string {
	if o == nil || o.Data == nil {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolverResourcesAllOf) GetDataOk() ([]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResolverResourcesAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *ResolverResourcesAllOf) SetData(v []string) {
	o.Data = v
}

// GetResolver returns the Resolver field value if set, zero value otherwise.
func (o *ResolverResourcesAllOf) GetResolver() ResolverType {
	if o == nil || o.Resolver == nil {
		var ret ResolverType
		return ret
	}
	return *o.Resolver
}

// GetResolverOk returns a tuple with the Resolver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolverResourcesAllOf) GetResolverOk() (*ResolverType, bool) {
	if o == nil || o.Resolver == nil {
		return nil, false
	}
	return o.Resolver, true
}

// HasResolver returns a boolean if a field has been set.
func (o *ResolverResourcesAllOf) HasResolver() bool {
	if o != nil && o.Resolver != nil {
		return true
	}

	return false
}

// SetResolver gets a reference to the given ResolverType and assigns it to the Resolver field.
func (o *ResolverResourcesAllOf) SetResolver(v ResolverType) {
	o.Resolver = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResolverResourcesAllOf) GetType() ResourceType {
	if o == nil || o.Type == nil {
		var ret ResourceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolverResourcesAllOf) GetTypeOk() (*ResourceType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResolverResourcesAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ResourceType and assigns it to the Type field.
func (o *ResolverResourcesAllOf) SetType(v ResourceType) {
	o.Type = &v
}

// GetGatewayName returns the GatewayName field value if set, zero value otherwise.
func (o *ResolverResourcesAllOf) GetGatewayName() string {
	if o == nil || o.GatewayName == nil {
		var ret string
		return ret
	}
	return *o.GatewayName
}

// GetGatewayNameOk returns a tuple with the GatewayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolverResourcesAllOf) GetGatewayNameOk() (*string, bool) {
	if o == nil || o.GatewayName == nil {
		return nil, false
	}
	return o.GatewayName, true
}

// HasGatewayName returns a boolean if a field has been set.
func (o *ResolverResourcesAllOf) HasGatewayName() bool {
	if o != nil && o.GatewayName != nil {
		return true
	}

	return false
}

// SetGatewayName gets a reference to the given string and assigns it to the GatewayName field.
func (o *ResolverResourcesAllOf) SetGatewayName(v string) {
	o.GatewayName = &v
}

func (o ResolverResourcesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Resolver != nil {
		toSerialize["resolver"] = o.Resolver
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.GatewayName != nil {
		toSerialize["gatewayName"] = o.GatewayName
	}
	return json.Marshal(toSerialize)
}

type NullableResolverResourcesAllOf struct {
	value *ResolverResourcesAllOf
	isSet bool
}

func (v NullableResolverResourcesAllOf) Get() *ResolverResourcesAllOf {
	return v.value
}

func (v *NullableResolverResourcesAllOf) Set(val *ResolverResourcesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResolverResourcesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResolverResourcesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolverResourcesAllOf(val *ResolverResourcesAllOf) *NullableResolverResourcesAllOf {
	return &NullableResolverResourcesAllOf{value: val, isSet: true}
}

func (v NullableResolverResourcesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolverResourcesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
