/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v19+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 19.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SiteAllOfNameResolutionAzureResolvers struct for SiteAllOfNameResolutionAzureResolvers
type SiteAllOfNameResolutionAzureResolvers struct {
	// Identifier name. Has no functional effect.
	Name string `json:"name"`
	// How often will the resolver poll the server. In seconds.
	UpdateInterval *int32 `json:"updateInterval,omitempty"`
	// Uses the built-in Managed Identities in Azure instances to authenticate against the API.
	UseManagedIdentities *bool `json:"useManagedIdentities,omitempty"`
	// Azure subscription id, visible with the azure cli command `azure account show`.
	// Deprecated
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// Azure tenant id, visible with the azure cli command `azure account show`.
	TenantId *string `json:"tenantId,omitempty"`
	// Azure client id, also called app id. Visible for a given application using the azure cli command `azure ad app show`.
	ClientId *string `json:"clientId,omitempty"`
	// Azure client secret. For Azure AD Apps this is done by creating a key for the app.
	Secret *string `json:"secret,omitempty"`
}

// NewSiteAllOfNameResolutionAzureResolvers instantiates a new SiteAllOfNameResolutionAzureResolvers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAllOfNameResolutionAzureResolvers(name string) *SiteAllOfNameResolutionAzureResolvers {
	this := SiteAllOfNameResolutionAzureResolvers{}
	this.Name = name
	var updateInterval int32 = 60
	this.UpdateInterval = &updateInterval
	var useManagedIdentities bool = false
	this.UseManagedIdentities = &useManagedIdentities
	return &this
}

// NewSiteAllOfNameResolutionAzureResolversWithDefaults instantiates a new SiteAllOfNameResolutionAzureResolvers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAllOfNameResolutionAzureResolversWithDefaults() *SiteAllOfNameResolutionAzureResolvers {
	this := SiteAllOfNameResolutionAzureResolvers{}
	var updateInterval int32 = 60
	this.UpdateInterval = &updateInterval
	var useManagedIdentities bool = false
	this.UseManagedIdentities = &useManagedIdentities
	return &this
}

// GetName returns the Name field value
func (o *SiteAllOfNameResolutionAzureResolvers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAzureResolvers) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteAllOfNameResolutionAzureResolvers) SetName(v string) {
	o.Name = v
}

// GetUpdateInterval returns the UpdateInterval field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAzureResolvers) GetUpdateInterval() int32 {
	if o == nil || o.UpdateInterval == nil {
		var ret int32
		return ret
	}
	return *o.UpdateInterval
}

// GetUpdateIntervalOk returns a tuple with the UpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAzureResolvers) GetUpdateIntervalOk() (*int32, bool) {
	if o == nil || o.UpdateInterval == nil {
		return nil, false
	}
	return o.UpdateInterval, true
}

// HasUpdateInterval returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAzureResolvers) HasUpdateInterval() bool {
	if o != nil && o.UpdateInterval != nil {
		return true
	}

	return false
}

// SetUpdateInterval gets a reference to the given int32 and assigns it to the UpdateInterval field.
func (o *SiteAllOfNameResolutionAzureResolvers) SetUpdateInterval(v int32) {
	o.UpdateInterval = &v
}

// GetUseManagedIdentities returns the UseManagedIdentities field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAzureResolvers) GetUseManagedIdentities() bool {
	if o == nil || o.UseManagedIdentities == nil {
		var ret bool
		return ret
	}
	return *o.UseManagedIdentities
}

// GetUseManagedIdentitiesOk returns a tuple with the UseManagedIdentities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAzureResolvers) GetUseManagedIdentitiesOk() (*bool, bool) {
	if o == nil || o.UseManagedIdentities == nil {
		return nil, false
	}
	return o.UseManagedIdentities, true
}

// HasUseManagedIdentities returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAzureResolvers) HasUseManagedIdentities() bool {
	if o != nil && o.UseManagedIdentities != nil {
		return true
	}

	return false
}

// SetUseManagedIdentities gets a reference to the given bool and assigns it to the UseManagedIdentities field.
func (o *SiteAllOfNameResolutionAzureResolvers) SetUseManagedIdentities(v bool) {
	o.UseManagedIdentities = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
// Deprecated
func (o *SiteAllOfNameResolutionAzureResolvers) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SiteAllOfNameResolutionAzureResolvers) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAzureResolvers) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
// Deprecated
func (o *SiteAllOfNameResolutionAzureResolvers) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAzureResolvers) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAzureResolvers) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAzureResolvers) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *SiteAllOfNameResolutionAzureResolvers) SetTenantId(v string) {
	o.TenantId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAzureResolvers) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAzureResolvers) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAzureResolvers) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SiteAllOfNameResolutionAzureResolvers) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAzureResolvers) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAzureResolvers) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAzureResolvers) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SiteAllOfNameResolutionAzureResolvers) SetSecret(v string) {
	o.Secret = &v
}

func (o SiteAllOfNameResolutionAzureResolvers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.UpdateInterval != nil {
		toSerialize["updateInterval"] = o.UpdateInterval
	}
	if o.UseManagedIdentities != nil {
		toSerialize["useManagedIdentities"] = o.UseManagedIdentities
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.TenantId != nil {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAllOfNameResolutionAzureResolvers struct {
	value *SiteAllOfNameResolutionAzureResolvers
	isSet bool
}

func (v NullableSiteAllOfNameResolutionAzureResolvers) Get() *SiteAllOfNameResolutionAzureResolvers {
	return v.value
}

func (v *NullableSiteAllOfNameResolutionAzureResolvers) Set(val *SiteAllOfNameResolutionAzureResolvers) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAllOfNameResolutionAzureResolvers) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAllOfNameResolutionAzureResolvers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAllOfNameResolutionAzureResolvers(val *SiteAllOfNameResolutionAzureResolvers) *NullableSiteAllOfNameResolutionAzureResolvers {
	return &NullableSiteAllOfNameResolutionAzureResolvers{value: val, isSet: true}
}

func (v NullableSiteAllOfNameResolutionAzureResolvers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAllOfNameResolutionAzureResolvers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
