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

// SiteAllOfNameResolutionGcpResolvers struct for SiteAllOfNameResolutionGcpResolvers
type SiteAllOfNameResolutionGcpResolvers struct {
	// Identifier name. Has no functional effect.
	Name string `json:"name"`
	// How often will the resolver poll the server. In seconds.
	UpdateInterval *int32 `json:"updateInterval,omitempty"`
	// GCP project filter.
	ProjectFilter *string `json:"projectFilter,omitempty"`
	// GCP instance filter.
	InstanceFilter *string `json:"instanceFilter,omitempty"`
	// GCP forwarding rules filter.
	ForwardingRulesFilter *string `json:"forwardingRulesFilter,omitempty"`
}

// NewSiteAllOfNameResolutionGcpResolvers instantiates a new SiteAllOfNameResolutionGcpResolvers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAllOfNameResolutionGcpResolvers(name string) *SiteAllOfNameResolutionGcpResolvers {
	this := SiteAllOfNameResolutionGcpResolvers{}
	this.Name = name
	var updateInterval int32 = 60
	this.UpdateInterval = &updateInterval
	return &this
}

// NewSiteAllOfNameResolutionGcpResolversWithDefaults instantiates a new SiteAllOfNameResolutionGcpResolvers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAllOfNameResolutionGcpResolversWithDefaults() *SiteAllOfNameResolutionGcpResolvers {
	this := SiteAllOfNameResolutionGcpResolvers{}
	var updateInterval int32 = 60
	this.UpdateInterval = &updateInterval
	return &this
}

// GetName returns the Name field value
func (o *SiteAllOfNameResolutionGcpResolvers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionGcpResolvers) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteAllOfNameResolutionGcpResolvers) SetName(v string) {
	o.Name = v
}

// GetUpdateInterval returns the UpdateInterval field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionGcpResolvers) GetUpdateInterval() int32 {
	if o == nil || o.UpdateInterval == nil {
		var ret int32
		return ret
	}
	return *o.UpdateInterval
}

// GetUpdateIntervalOk returns a tuple with the UpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionGcpResolvers) GetUpdateIntervalOk() (*int32, bool) {
	if o == nil || o.UpdateInterval == nil {
		return nil, false
	}
	return o.UpdateInterval, true
}

// HasUpdateInterval returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionGcpResolvers) HasUpdateInterval() bool {
	if o != nil && o.UpdateInterval != nil {
		return true
	}

	return false
}

// SetUpdateInterval gets a reference to the given int32 and assigns it to the UpdateInterval field.
func (o *SiteAllOfNameResolutionGcpResolvers) SetUpdateInterval(v int32) {
	o.UpdateInterval = &v
}

// GetProjectFilter returns the ProjectFilter field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionGcpResolvers) GetProjectFilter() string {
	if o == nil || o.ProjectFilter == nil {
		var ret string
		return ret
	}
	return *o.ProjectFilter
}

// GetProjectFilterOk returns a tuple with the ProjectFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionGcpResolvers) GetProjectFilterOk() (*string, bool) {
	if o == nil || o.ProjectFilter == nil {
		return nil, false
	}
	return o.ProjectFilter, true
}

// HasProjectFilter returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionGcpResolvers) HasProjectFilter() bool {
	if o != nil && o.ProjectFilter != nil {
		return true
	}

	return false
}

// SetProjectFilter gets a reference to the given string and assigns it to the ProjectFilter field.
func (o *SiteAllOfNameResolutionGcpResolvers) SetProjectFilter(v string) {
	o.ProjectFilter = &v
}

// GetInstanceFilter returns the InstanceFilter field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionGcpResolvers) GetInstanceFilter() string {
	if o == nil || o.InstanceFilter == nil {
		var ret string
		return ret
	}
	return *o.InstanceFilter
}

// GetInstanceFilterOk returns a tuple with the InstanceFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionGcpResolvers) GetInstanceFilterOk() (*string, bool) {
	if o == nil || o.InstanceFilter == nil {
		return nil, false
	}
	return o.InstanceFilter, true
}

// HasInstanceFilter returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionGcpResolvers) HasInstanceFilter() bool {
	if o != nil && o.InstanceFilter != nil {
		return true
	}

	return false
}

// SetInstanceFilter gets a reference to the given string and assigns it to the InstanceFilter field.
func (o *SiteAllOfNameResolutionGcpResolvers) SetInstanceFilter(v string) {
	o.InstanceFilter = &v
}

// GetForwardingRulesFilter returns the ForwardingRulesFilter field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionGcpResolvers) GetForwardingRulesFilter() string {
	if o == nil || o.ForwardingRulesFilter == nil {
		var ret string
		return ret
	}
	return *o.ForwardingRulesFilter
}

// GetForwardingRulesFilterOk returns a tuple with the ForwardingRulesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionGcpResolvers) GetForwardingRulesFilterOk() (*string, bool) {
	if o == nil || o.ForwardingRulesFilter == nil {
		return nil, false
	}
	return o.ForwardingRulesFilter, true
}

// HasForwardingRulesFilter returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionGcpResolvers) HasForwardingRulesFilter() bool {
	if o != nil && o.ForwardingRulesFilter != nil {
		return true
	}

	return false
}

// SetForwardingRulesFilter gets a reference to the given string and assigns it to the ForwardingRulesFilter field.
func (o *SiteAllOfNameResolutionGcpResolvers) SetForwardingRulesFilter(v string) {
	o.ForwardingRulesFilter = &v
}

func (o SiteAllOfNameResolutionGcpResolvers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.UpdateInterval != nil {
		toSerialize["updateInterval"] = o.UpdateInterval
	}
	if o.ProjectFilter != nil {
		toSerialize["projectFilter"] = o.ProjectFilter
	}
	if o.InstanceFilter != nil {
		toSerialize["instanceFilter"] = o.InstanceFilter
	}
	if o.ForwardingRulesFilter != nil {
		toSerialize["forwardingRulesFilter"] = o.ForwardingRulesFilter
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAllOfNameResolutionGcpResolvers struct {
	value *SiteAllOfNameResolutionGcpResolvers
	isSet bool
}

func (v NullableSiteAllOfNameResolutionGcpResolvers) Get() *SiteAllOfNameResolutionGcpResolvers {
	return v.value
}

func (v *NullableSiteAllOfNameResolutionGcpResolvers) Set(val *SiteAllOfNameResolutionGcpResolvers) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAllOfNameResolutionGcpResolvers) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAllOfNameResolutionGcpResolvers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAllOfNameResolutionGcpResolvers(val *SiteAllOfNameResolutionGcpResolvers) *NullableSiteAllOfNameResolutionGcpResolvers {
	return &NullableSiteAllOfNameResolutionGcpResolvers{value: val, isSet: true}
}

func (v NullableSiteAllOfNameResolutionGcpResolvers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAllOfNameResolutionGcpResolvers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
