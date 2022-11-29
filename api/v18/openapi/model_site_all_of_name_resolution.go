/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SiteAllOfNameResolution Settings for asset name resolution.
type SiteAllOfNameResolution struct {
	// Name resolution to use Appliance's /etc/hosts file.
	UseHostsFile *bool `json:"useHostsFile,omitempty"`
	// Resolver to resolve hostnames using DNS servers.
	DnsResolvers []SiteAllOfNameResolutionDnsResolvers `json:"dnsResolvers,omitempty"`
	// Resolvers to resolve Amazon machines by querying Amazon Web Services.
	AwsResolvers []SiteAllOfNameResolutionAwsResolvers `json:"awsResolvers,omitempty"`
	// Resolvers to resolve Azure machines by querying Azure App Service.
	AzureResolvers []SiteAllOfNameResolutionAzureResolvers `json:"azureResolvers,omitempty"`
	// Resolvers to resolve VMware vSphere machines by querying the vCenter.
	EsxResolvers []SiteAllOfNameResolutionEsxResolvers `json:"esxResolvers,omitempty"`
	// Resolvers to resolve GCP machine by querying Google web services.
	GcpResolvers []SiteAllOfNameResolutionGcpResolvers `json:"gcpResolvers,omitempty"`
	// Resolvers to resolve names by querying Appgate Illumio Resolver.
	IllumioResolvers []SiteAllOfNameResolutionIllumioResolvers `json:"illumioResolvers,omitempty"`
	DnsForwarding    *SiteAllOfNameResolutionDnsForwarding     `json:"dnsForwarding,omitempty"`
}

// NewSiteAllOfNameResolution instantiates a new SiteAllOfNameResolution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAllOfNameResolution() *SiteAllOfNameResolution {
	this := SiteAllOfNameResolution{}
	var useHostsFile bool = false
	this.UseHostsFile = &useHostsFile
	return &this
}

// NewSiteAllOfNameResolutionWithDefaults instantiates a new SiteAllOfNameResolution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAllOfNameResolutionWithDefaults() *SiteAllOfNameResolution {
	this := SiteAllOfNameResolution{}
	var useHostsFile bool = false
	this.UseHostsFile = &useHostsFile
	return &this
}

// GetUseHostsFile returns the UseHostsFile field value if set, zero value otherwise.
func (o *SiteAllOfNameResolution) GetUseHostsFile() bool {
	if o == nil || o.UseHostsFile == nil {
		var ret bool
		return ret
	}
	return *o.UseHostsFile
}

// GetUseHostsFileOk returns a tuple with the UseHostsFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolution) GetUseHostsFileOk() (*bool, bool) {
	if o == nil || o.UseHostsFile == nil {
		return nil, false
	}
	return o.UseHostsFile, true
}

// HasUseHostsFile returns a boolean if a field has been set.
func (o *SiteAllOfNameResolution) HasUseHostsFile() bool {
	if o != nil && o.UseHostsFile != nil {
		return true
	}

	return false
}

// SetUseHostsFile gets a reference to the given bool and assigns it to the UseHostsFile field.
func (o *SiteAllOfNameResolution) SetUseHostsFile(v bool) {
	o.UseHostsFile = &v
}

// GetDnsResolvers returns the DnsResolvers field value if set, zero value otherwise.
func (o *SiteAllOfNameResolution) GetDnsResolvers() []SiteAllOfNameResolutionDnsResolvers {
	if o == nil || o.DnsResolvers == nil {
		var ret []SiteAllOfNameResolutionDnsResolvers
		return ret
	}
	return o.DnsResolvers
}

// GetDnsResolversOk returns a tuple with the DnsResolvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolution) GetDnsResolversOk() ([]SiteAllOfNameResolutionDnsResolvers, bool) {
	if o == nil || o.DnsResolvers == nil {
		return nil, false
	}
	return o.DnsResolvers, true
}

// HasDnsResolvers returns a boolean if a field has been set.
func (o *SiteAllOfNameResolution) HasDnsResolvers() bool {
	if o != nil && o.DnsResolvers != nil {
		return true
	}

	return false
}

// SetDnsResolvers gets a reference to the given []SiteAllOfNameResolutionDnsResolvers and assigns it to the DnsResolvers field.
func (o *SiteAllOfNameResolution) SetDnsResolvers(v []SiteAllOfNameResolutionDnsResolvers) {
	o.DnsResolvers = v
}

// GetAwsResolvers returns the AwsResolvers field value if set, zero value otherwise.
func (o *SiteAllOfNameResolution) GetAwsResolvers() []SiteAllOfNameResolutionAwsResolvers {
	if o == nil || o.AwsResolvers == nil {
		var ret []SiteAllOfNameResolutionAwsResolvers
		return ret
	}
	return o.AwsResolvers
}

// GetAwsResolversOk returns a tuple with the AwsResolvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolution) GetAwsResolversOk() ([]SiteAllOfNameResolutionAwsResolvers, bool) {
	if o == nil || o.AwsResolvers == nil {
		return nil, false
	}
	return o.AwsResolvers, true
}

// HasAwsResolvers returns a boolean if a field has been set.
func (o *SiteAllOfNameResolution) HasAwsResolvers() bool {
	if o != nil && o.AwsResolvers != nil {
		return true
	}

	return false
}

// SetAwsResolvers gets a reference to the given []SiteAllOfNameResolutionAwsResolvers and assigns it to the AwsResolvers field.
func (o *SiteAllOfNameResolution) SetAwsResolvers(v []SiteAllOfNameResolutionAwsResolvers) {
	o.AwsResolvers = v
}

// GetAzureResolvers returns the AzureResolvers field value if set, zero value otherwise.
func (o *SiteAllOfNameResolution) GetAzureResolvers() []SiteAllOfNameResolutionAzureResolvers {
	if o == nil || o.AzureResolvers == nil {
		var ret []SiteAllOfNameResolutionAzureResolvers
		return ret
	}
	return o.AzureResolvers
}

// GetAzureResolversOk returns a tuple with the AzureResolvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolution) GetAzureResolversOk() ([]SiteAllOfNameResolutionAzureResolvers, bool) {
	if o == nil || o.AzureResolvers == nil {
		return nil, false
	}
	return o.AzureResolvers, true
}

// HasAzureResolvers returns a boolean if a field has been set.
func (o *SiteAllOfNameResolution) HasAzureResolvers() bool {
	if o != nil && o.AzureResolvers != nil {
		return true
	}

	return false
}

// SetAzureResolvers gets a reference to the given []SiteAllOfNameResolutionAzureResolvers and assigns it to the AzureResolvers field.
func (o *SiteAllOfNameResolution) SetAzureResolvers(v []SiteAllOfNameResolutionAzureResolvers) {
	o.AzureResolvers = v
}

// GetEsxResolvers returns the EsxResolvers field value if set, zero value otherwise.
func (o *SiteAllOfNameResolution) GetEsxResolvers() []SiteAllOfNameResolutionEsxResolvers {
	if o == nil || o.EsxResolvers == nil {
		var ret []SiteAllOfNameResolutionEsxResolvers
		return ret
	}
	return o.EsxResolvers
}

// GetEsxResolversOk returns a tuple with the EsxResolvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolution) GetEsxResolversOk() ([]SiteAllOfNameResolutionEsxResolvers, bool) {
	if o == nil || o.EsxResolvers == nil {
		return nil, false
	}
	return o.EsxResolvers, true
}

// HasEsxResolvers returns a boolean if a field has been set.
func (o *SiteAllOfNameResolution) HasEsxResolvers() bool {
	if o != nil && o.EsxResolvers != nil {
		return true
	}

	return false
}

// SetEsxResolvers gets a reference to the given []SiteAllOfNameResolutionEsxResolvers and assigns it to the EsxResolvers field.
func (o *SiteAllOfNameResolution) SetEsxResolvers(v []SiteAllOfNameResolutionEsxResolvers) {
	o.EsxResolvers = v
}

// GetGcpResolvers returns the GcpResolvers field value if set, zero value otherwise.
func (o *SiteAllOfNameResolution) GetGcpResolvers() []SiteAllOfNameResolutionGcpResolvers {
	if o == nil || o.GcpResolvers == nil {
		var ret []SiteAllOfNameResolutionGcpResolvers
		return ret
	}
	return o.GcpResolvers
}

// GetGcpResolversOk returns a tuple with the GcpResolvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolution) GetGcpResolversOk() ([]SiteAllOfNameResolutionGcpResolvers, bool) {
	if o == nil || o.GcpResolvers == nil {
		return nil, false
	}
	return o.GcpResolvers, true
}

// HasGcpResolvers returns a boolean if a field has been set.
func (o *SiteAllOfNameResolution) HasGcpResolvers() bool {
	if o != nil && o.GcpResolvers != nil {
		return true
	}

	return false
}

// SetGcpResolvers gets a reference to the given []SiteAllOfNameResolutionGcpResolvers and assigns it to the GcpResolvers field.
func (o *SiteAllOfNameResolution) SetGcpResolvers(v []SiteAllOfNameResolutionGcpResolvers) {
	o.GcpResolvers = v
}

// GetIllumioResolvers returns the IllumioResolvers field value if set, zero value otherwise.
func (o *SiteAllOfNameResolution) GetIllumioResolvers() []SiteAllOfNameResolutionIllumioResolvers {
	if o == nil || o.IllumioResolvers == nil {
		var ret []SiteAllOfNameResolutionIllumioResolvers
		return ret
	}
	return o.IllumioResolvers
}

// GetIllumioResolversOk returns a tuple with the IllumioResolvers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolution) GetIllumioResolversOk() ([]SiteAllOfNameResolutionIllumioResolvers, bool) {
	if o == nil || o.IllumioResolvers == nil {
		return nil, false
	}
	return o.IllumioResolvers, true
}

// HasIllumioResolvers returns a boolean if a field has been set.
func (o *SiteAllOfNameResolution) HasIllumioResolvers() bool {
	if o != nil && o.IllumioResolvers != nil {
		return true
	}

	return false
}

// SetIllumioResolvers gets a reference to the given []SiteAllOfNameResolutionIllumioResolvers and assigns it to the IllumioResolvers field.
func (o *SiteAllOfNameResolution) SetIllumioResolvers(v []SiteAllOfNameResolutionIllumioResolvers) {
	o.IllumioResolvers = v
}

// GetDnsForwarding returns the DnsForwarding field value if set, zero value otherwise.
func (o *SiteAllOfNameResolution) GetDnsForwarding() SiteAllOfNameResolutionDnsForwarding {
	if o == nil || o.DnsForwarding == nil {
		var ret SiteAllOfNameResolutionDnsForwarding
		return ret
	}
	return *o.DnsForwarding
}

// GetDnsForwardingOk returns a tuple with the DnsForwarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolution) GetDnsForwardingOk() (*SiteAllOfNameResolutionDnsForwarding, bool) {
	if o == nil || o.DnsForwarding == nil {
		return nil, false
	}
	return o.DnsForwarding, true
}

// HasDnsForwarding returns a boolean if a field has been set.
func (o *SiteAllOfNameResolution) HasDnsForwarding() bool {
	if o != nil && o.DnsForwarding != nil {
		return true
	}

	return false
}

// SetDnsForwarding gets a reference to the given SiteAllOfNameResolutionDnsForwarding and assigns it to the DnsForwarding field.
func (o *SiteAllOfNameResolution) SetDnsForwarding(v SiteAllOfNameResolutionDnsForwarding) {
	o.DnsForwarding = &v
}

func (o SiteAllOfNameResolution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UseHostsFile != nil {
		toSerialize["useHostsFile"] = o.UseHostsFile
	}
	if o.DnsResolvers != nil {
		toSerialize["dnsResolvers"] = o.DnsResolvers
	}
	if o.AwsResolvers != nil {
		toSerialize["awsResolvers"] = o.AwsResolvers
	}
	if o.AzureResolvers != nil {
		toSerialize["azureResolvers"] = o.AzureResolvers
	}
	if o.EsxResolvers != nil {
		toSerialize["esxResolvers"] = o.EsxResolvers
	}
	if o.GcpResolvers != nil {
		toSerialize["gcpResolvers"] = o.GcpResolvers
	}
	if o.IllumioResolvers != nil {
		toSerialize["illumioResolvers"] = o.IllumioResolvers
	}
	if o.DnsForwarding != nil {
		toSerialize["dnsForwarding"] = o.DnsForwarding
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAllOfNameResolution struct {
	value *SiteAllOfNameResolution
	isSet bool
}

func (v NullableSiteAllOfNameResolution) Get() *SiteAllOfNameResolution {
	return v.value
}

func (v *NullableSiteAllOfNameResolution) Set(val *SiteAllOfNameResolution) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAllOfNameResolution) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAllOfNameResolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAllOfNameResolution(val *SiteAllOfNameResolution) *NullableSiteAllOfNameResolution {
	return &NullableSiteAllOfNameResolution{value: val, isSet: true}
}

func (v NullableSiteAllOfNameResolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAllOfNameResolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
