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

// SiteAllOfNameResolutionAwsResolvers struct for SiteAllOfNameResolutionAwsResolvers
type SiteAllOfNameResolutionAwsResolvers struct {
	// Identifier name. Has no functional effect.
	Name string `json:"name"`
	// How often will the resolver poll the server. In seconds.
	UpdateInterval *int32 `json:"updateInterval,omitempty"`
	// VPC IDs to resolve names.
	Vpcs []string `json:"vpcs,omitempty"`
	// Use VPC auto discovery.
	VpcAutoDiscovery *bool `json:"vpcAutoDiscovery,omitempty"`
	// Amazon regions.
	Regions []string `json:"regions,omitempty"`
	// Uses the built-in IAM role in AWS instances to authenticate against the API.
	UseIAMRole *bool `json:"useIAMRole,omitempty"`
	// ID of the access key.
	AccessKeyId *string `json:"accessKeyId,omitempty"`
	// Secret access key for accessKeyId.
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	// Proxy address to use while communicating with AWS. format: https://username:password@ip/hostname:port
	HttpsProxy *string `json:"httpsProxy,omitempty"`
	// Use master credentials to resolve names in addition to any assumed roles.
	ResolveWithMasterCredentials *bool `json:"resolveWithMasterCredentials,omitempty"`
	// What AWS partition to use such as 'aws-cn' or 'aws-us-gov'
	Partition *string `json:"partition,omitempty"`
	// Roles to be assumed to perform AWS name resolution.
	AssumedRoles []SiteAllOfNameResolutionAssumedRoles `json:"assumedRoles,omitempty"`
}

// NewSiteAllOfNameResolutionAwsResolvers instantiates a new SiteAllOfNameResolutionAwsResolvers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteAllOfNameResolutionAwsResolvers(name string) *SiteAllOfNameResolutionAwsResolvers {
	this := SiteAllOfNameResolutionAwsResolvers{}
	this.Name = name
	var updateInterval int32 = 60
	this.UpdateInterval = &updateInterval
	var partition string = "aws"
	this.Partition = &partition
	return &this
}

// NewSiteAllOfNameResolutionAwsResolversWithDefaults instantiates a new SiteAllOfNameResolutionAwsResolvers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteAllOfNameResolutionAwsResolversWithDefaults() *SiteAllOfNameResolutionAwsResolvers {
	this := SiteAllOfNameResolutionAwsResolvers{}
	var updateInterval int32 = 60
	this.UpdateInterval = &updateInterval
	var partition string = "aws"
	this.Partition = &partition
	return &this
}

// GetName returns the Name field value
func (o *SiteAllOfNameResolutionAwsResolvers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteAllOfNameResolutionAwsResolvers) SetName(v string) {
	o.Name = v
}

// GetUpdateInterval returns the UpdateInterval field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAwsResolvers) GetUpdateInterval() int32 {
	if o == nil || o.UpdateInterval == nil {
		var ret int32
		return ret
	}
	return *o.UpdateInterval
}

// GetUpdateIntervalOk returns a tuple with the UpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) GetUpdateIntervalOk() (*int32, bool) {
	if o == nil || o.UpdateInterval == nil {
		return nil, false
	}
	return o.UpdateInterval, true
}

// HasUpdateInterval returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) HasUpdateInterval() bool {
	if o != nil && o.UpdateInterval != nil {
		return true
	}

	return false
}

// SetUpdateInterval gets a reference to the given int32 and assigns it to the UpdateInterval field.
func (o *SiteAllOfNameResolutionAwsResolvers) SetUpdateInterval(v int32) {
	o.UpdateInterval = &v
}

// GetVpcs returns the Vpcs field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAwsResolvers) GetVpcs() []string {
	if o == nil || o.Vpcs == nil {
		var ret []string
		return ret
	}
	return o.Vpcs
}

// GetVpcsOk returns a tuple with the Vpcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) GetVpcsOk() ([]string, bool) {
	if o == nil || o.Vpcs == nil {
		return nil, false
	}
	return o.Vpcs, true
}

// HasVpcs returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) HasVpcs() bool {
	if o != nil && o.Vpcs != nil {
		return true
	}

	return false
}

// SetVpcs gets a reference to the given []string and assigns it to the Vpcs field.
func (o *SiteAllOfNameResolutionAwsResolvers) SetVpcs(v []string) {
	o.Vpcs = v
}

// GetVpcAutoDiscovery returns the VpcAutoDiscovery field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAwsResolvers) GetVpcAutoDiscovery() bool {
	if o == nil || o.VpcAutoDiscovery == nil {
		var ret bool
		return ret
	}
	return *o.VpcAutoDiscovery
}

// GetVpcAutoDiscoveryOk returns a tuple with the VpcAutoDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) GetVpcAutoDiscoveryOk() (*bool, bool) {
	if o == nil || o.VpcAutoDiscovery == nil {
		return nil, false
	}
	return o.VpcAutoDiscovery, true
}

// HasVpcAutoDiscovery returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) HasVpcAutoDiscovery() bool {
	if o != nil && o.VpcAutoDiscovery != nil {
		return true
	}

	return false
}

// SetVpcAutoDiscovery gets a reference to the given bool and assigns it to the VpcAutoDiscovery field.
func (o *SiteAllOfNameResolutionAwsResolvers) SetVpcAutoDiscovery(v bool) {
	o.VpcAutoDiscovery = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAwsResolvers) GetRegions() []string {
	if o == nil || o.Regions == nil {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) GetRegionsOk() ([]string, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *SiteAllOfNameResolutionAwsResolvers) SetRegions(v []string) {
	o.Regions = v
}

// GetUseIAMRole returns the UseIAMRole field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAwsResolvers) GetUseIAMRole() bool {
	if o == nil || o.UseIAMRole == nil {
		var ret bool
		return ret
	}
	return *o.UseIAMRole
}

// GetUseIAMRoleOk returns a tuple with the UseIAMRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) GetUseIAMRoleOk() (*bool, bool) {
	if o == nil || o.UseIAMRole == nil {
		return nil, false
	}
	return o.UseIAMRole, true
}

// HasUseIAMRole returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) HasUseIAMRole() bool {
	if o != nil && o.UseIAMRole != nil {
		return true
	}

	return false
}

// SetUseIAMRole gets a reference to the given bool and assigns it to the UseIAMRole field.
func (o *SiteAllOfNameResolutionAwsResolvers) SetUseIAMRole(v bool) {
	o.UseIAMRole = &v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAwsResolvers) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId != nil {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *SiteAllOfNameResolutionAwsResolvers) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAwsResolvers) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.SecretAccessKey == nil {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) HasSecretAccessKey() bool {
	if o != nil && o.SecretAccessKey != nil {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *SiteAllOfNameResolutionAwsResolvers) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

// GetHttpsProxy returns the HttpsProxy field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAwsResolvers) GetHttpsProxy() string {
	if o == nil || o.HttpsProxy == nil {
		var ret string
		return ret
	}
	return *o.HttpsProxy
}

// GetHttpsProxyOk returns a tuple with the HttpsProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) GetHttpsProxyOk() (*string, bool) {
	if o == nil || o.HttpsProxy == nil {
		return nil, false
	}
	return o.HttpsProxy, true
}

// HasHttpsProxy returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) HasHttpsProxy() bool {
	if o != nil && o.HttpsProxy != nil {
		return true
	}

	return false
}

// SetHttpsProxy gets a reference to the given string and assigns it to the HttpsProxy field.
func (o *SiteAllOfNameResolutionAwsResolvers) SetHttpsProxy(v string) {
	o.HttpsProxy = &v
}

// GetResolveWithMasterCredentials returns the ResolveWithMasterCredentials field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAwsResolvers) GetResolveWithMasterCredentials() bool {
	if o == nil || o.ResolveWithMasterCredentials == nil {
		var ret bool
		return ret
	}
	return *o.ResolveWithMasterCredentials
}

// GetResolveWithMasterCredentialsOk returns a tuple with the ResolveWithMasterCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) GetResolveWithMasterCredentialsOk() (*bool, bool) {
	if o == nil || o.ResolveWithMasterCredentials == nil {
		return nil, false
	}
	return o.ResolveWithMasterCredentials, true
}

// HasResolveWithMasterCredentials returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) HasResolveWithMasterCredentials() bool {
	if o != nil && o.ResolveWithMasterCredentials != nil {
		return true
	}

	return false
}

// SetResolveWithMasterCredentials gets a reference to the given bool and assigns it to the ResolveWithMasterCredentials field.
func (o *SiteAllOfNameResolutionAwsResolvers) SetResolveWithMasterCredentials(v bool) {
	o.ResolveWithMasterCredentials = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAwsResolvers) GetPartition() string {
	if o == nil || o.Partition == nil {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) GetPartitionOk() (*string, bool) {
	if o == nil || o.Partition == nil {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) HasPartition() bool {
	if o != nil && o.Partition != nil {
		return true
	}

	return false
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *SiteAllOfNameResolutionAwsResolvers) SetPartition(v string) {
	o.Partition = &v
}

// GetAssumedRoles returns the AssumedRoles field value if set, zero value otherwise.
func (o *SiteAllOfNameResolutionAwsResolvers) GetAssumedRoles() []SiteAllOfNameResolutionAssumedRoles {
	if o == nil || o.AssumedRoles == nil {
		var ret []SiteAllOfNameResolutionAssumedRoles
		return ret
	}
	return o.AssumedRoles
}

// GetAssumedRolesOk returns a tuple with the AssumedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) GetAssumedRolesOk() ([]SiteAllOfNameResolutionAssumedRoles, bool) {
	if o == nil || o.AssumedRoles == nil {
		return nil, false
	}
	return o.AssumedRoles, true
}

// HasAssumedRoles returns a boolean if a field has been set.
func (o *SiteAllOfNameResolutionAwsResolvers) HasAssumedRoles() bool {
	if o != nil && o.AssumedRoles != nil {
		return true
	}

	return false
}

// SetAssumedRoles gets a reference to the given []SiteAllOfNameResolutionAssumedRoles and assigns it to the AssumedRoles field.
func (o *SiteAllOfNameResolutionAwsResolvers) SetAssumedRoles(v []SiteAllOfNameResolutionAssumedRoles) {
	o.AssumedRoles = v
}

func (o SiteAllOfNameResolutionAwsResolvers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.UpdateInterval != nil {
		toSerialize["updateInterval"] = o.UpdateInterval
	}
	if o.Vpcs != nil {
		toSerialize["vpcs"] = o.Vpcs
	}
	if o.VpcAutoDiscovery != nil {
		toSerialize["vpcAutoDiscovery"] = o.VpcAutoDiscovery
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	if o.UseIAMRole != nil {
		toSerialize["useIAMRole"] = o.UseIAMRole
	}
	if o.AccessKeyId != nil {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	if o.SecretAccessKey != nil {
		toSerialize["secretAccessKey"] = o.SecretAccessKey
	}
	if o.HttpsProxy != nil {
		toSerialize["httpsProxy"] = o.HttpsProxy
	}
	if o.ResolveWithMasterCredentials != nil {
		toSerialize["resolveWithMasterCredentials"] = o.ResolveWithMasterCredentials
	}
	if o.Partition != nil {
		toSerialize["partition"] = o.Partition
	}
	if o.AssumedRoles != nil {
		toSerialize["assumedRoles"] = o.AssumedRoles
	}
	return json.Marshal(toSerialize)
}

type NullableSiteAllOfNameResolutionAwsResolvers struct {
	value *SiteAllOfNameResolutionAwsResolvers
	isSet bool
}

func (v NullableSiteAllOfNameResolutionAwsResolvers) Get() *SiteAllOfNameResolutionAwsResolvers {
	return v.value
}

func (v *NullableSiteAllOfNameResolutionAwsResolvers) Set(val *SiteAllOfNameResolutionAwsResolvers) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteAllOfNameResolutionAwsResolvers) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteAllOfNameResolutionAwsResolvers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteAllOfNameResolutionAwsResolvers(val *SiteAllOfNameResolutionAwsResolvers) *NullableSiteAllOfNameResolutionAwsResolvers {
	return &NullableSiteAllOfNameResolutionAwsResolvers{value: val, isSet: true}
}

func (v NullableSiteAllOfNameResolutionAwsResolvers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteAllOfNameResolutionAwsResolvers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
