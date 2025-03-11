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

// TcpClient struct for TcpClient
type TcpClient struct {
	// Name of the endpoint.
	Name string `json:"name"`
	// Hostname or the IP address of the endpoint.
	Host string `json:"host"`
	// Port of the endpoint.
	Port int32 `json:"port"`
	// The format to send the audit logs.
	Format string `json:"format"`
	// Whether to use TLS to connect to endpoint or not. If enabled, make sure the LogForwarder appliance trusts the certificate of the endpoint.
	UseTLS *bool `json:"useTLS,omitempty"`
	// JMESPath expression to filter audit logs to forward.
	Filter *string `json:"filter,omitempty"`
}

// NewTcpClient instantiates a new TcpClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTcpClient(name string, host string, port int32, format string) *TcpClient {
	this := TcpClient{}
	this.Name = name
	this.Host = host
	this.Port = port
	this.Format = format
	return &this
}

// NewTcpClientWithDefaults instantiates a new TcpClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTcpClientWithDefaults() *TcpClient {
	this := TcpClient{}
	return &this
}

// GetName returns the Name field value
func (o *TcpClient) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TcpClient) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TcpClient) SetName(v string) {
	o.Name = v
}

// GetHost returns the Host field value
func (o *TcpClient) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *TcpClient) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *TcpClient) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *TcpClient) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *TcpClient) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *TcpClient) SetPort(v int32) {
	o.Port = v
}

// GetFormat returns the Format field value
func (o *TcpClient) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *TcpClient) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *TcpClient) SetFormat(v string) {
	o.Format = v
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *TcpClient) GetUseTLS() bool {
	if o == nil || o.UseTLS == nil {
		var ret bool
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpClient) GetUseTLSOk() (*bool, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *TcpClient) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given bool and assigns it to the UseTLS field.
func (o *TcpClient) SetUseTLS(v bool) {
	o.UseTLS = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TcpClient) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TcpClient) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TcpClient) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *TcpClient) SetFilter(v string) {
	o.Filter = &v
}

func (o TcpClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["format"] = o.Format
	}
	if o.UseTLS != nil {
		toSerialize["useTLS"] = o.UseTLS
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	return json.Marshal(toSerialize)
}

type NullableTcpClient struct {
	value *TcpClient
	isSet bool
}

func (v NullableTcpClient) Get() *TcpClient {
	return v.value
}

func (v *NullableTcpClient) Set(val *TcpClient) {
	v.value = val
	v.isSet = true
}

func (v NullableTcpClient) IsSet() bool {
	return v.isSet
}

func (v *NullableTcpClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTcpClient(val *TcpClient) *NullableTcpClient {
	return &NullableTcpClient{value: val, isSet: true}
}

func (v NullableTcpClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTcpClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
