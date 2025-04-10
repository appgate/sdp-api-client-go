/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v22+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 22.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ClientProfile struct for ClientProfile
type ClientProfile struct {
	// ID of the object.
	Id *string `json:"id,omitempty"`
	// A name to identify the client profile. It will appear on the client UI.
	Name string `json:"name"`
	// Notes for the object. Used for documentation purposes.
	Notes *string `json:"notes,omitempty"`
	// Create date.
	Created *time.Time `json:"created,omitempty"`
	// Last update date.
	Updated *time.Time `json:"updated,omitempty"`
	// Array of tags.
	Tags []string `json:"tags,omitempty"`
	// Type of the client profile.
	Type *string `json:"type,omitempty"`
	// SPA key name to be used in the profile. Same key names in different profiles will have the same SPA key. SPA key is used by the client to connect to the controllers. If left empty, the name will be generated by the controller using the profile name.
	// Deprecated
	SpaKeyName *string `json:"spaKeyName,omitempty"`
	// Name of the Identity Provider to be used to authenticate.
	IdentityProviderName string `json:"identityProviderName"`
	// Overrides the Profile Hostname in global settings for this specific profile. Generated URLs will use this hostname instead.
	Hostname *string `json:"hostname,omitempty"`
	// The Global Profile Hostname as defined in global settings. Generated URLs will use this hostname.
	GlobalHostname *string `json:"globalHostname,omitempty"`
	// Exported time is the last time when the client profile is exported either using  the URL or Barcode.
	Exported *time.Time `json:"exported,omitempty"`
}

// NewClientProfile instantiates a new ClientProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientProfile(name string, identityProviderName string) *ClientProfile {
	this := ClientProfile{}
	this.Name = name
	var type_ string = "Profile"
	this.Type = &type_
	this.IdentityProviderName = identityProviderName
	return &this
}

// NewClientProfileWithDefaults instantiates a new ClientProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientProfileWithDefaults() *ClientProfile {
	this := ClientProfile{}
	var type_ string = "Profile"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClientProfile) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfile) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClientProfile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClientProfile) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ClientProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClientProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClientProfile) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ClientProfile) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfile) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ClientProfile) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ClientProfile) SetNotes(v string) {
	o.Notes = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ClientProfile) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfile) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ClientProfile) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ClientProfile) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ClientProfile) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfile) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ClientProfile) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *ClientProfile) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ClientProfile) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfile) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ClientProfile) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ClientProfile) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClientProfile) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfile) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClientProfile) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClientProfile) SetType(v string) {
	o.Type = &v
}

// GetSpaKeyName returns the SpaKeyName field value if set, zero value otherwise.
// Deprecated
func (o *ClientProfile) GetSpaKeyName() string {
	if o == nil || o.SpaKeyName == nil {
		var ret string
		return ret
	}
	return *o.SpaKeyName
}

// GetSpaKeyNameOk returns a tuple with the SpaKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ClientProfile) GetSpaKeyNameOk() (*string, bool) {
	if o == nil || o.SpaKeyName == nil {
		return nil, false
	}
	return o.SpaKeyName, true
}

// HasSpaKeyName returns a boolean if a field has been set.
func (o *ClientProfile) HasSpaKeyName() bool {
	if o != nil && o.SpaKeyName != nil {
		return true
	}

	return false
}

// SetSpaKeyName gets a reference to the given string and assigns it to the SpaKeyName field.
// Deprecated
func (o *ClientProfile) SetSpaKeyName(v string) {
	o.SpaKeyName = &v
}

// GetIdentityProviderName returns the IdentityProviderName field value
func (o *ClientProfile) GetIdentityProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityProviderName
}

// GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field value
// and a boolean to check if the value has been set.
func (o *ClientProfile) GetIdentityProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProviderName, true
}

// SetIdentityProviderName sets field value
func (o *ClientProfile) SetIdentityProviderName(v string) {
	o.IdentityProviderName = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ClientProfile) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfile) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ClientProfile) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ClientProfile) SetHostname(v string) {
	o.Hostname = &v
}

// GetGlobalHostname returns the GlobalHostname field value if set, zero value otherwise.
func (o *ClientProfile) GetGlobalHostname() string {
	if o == nil || o.GlobalHostname == nil {
		var ret string
		return ret
	}
	return *o.GlobalHostname
}

// GetGlobalHostnameOk returns a tuple with the GlobalHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfile) GetGlobalHostnameOk() (*string, bool) {
	if o == nil || o.GlobalHostname == nil {
		return nil, false
	}
	return o.GlobalHostname, true
}

// HasGlobalHostname returns a boolean if a field has been set.
func (o *ClientProfile) HasGlobalHostname() bool {
	if o != nil && o.GlobalHostname != nil {
		return true
	}

	return false
}

// SetGlobalHostname gets a reference to the given string and assigns it to the GlobalHostname field.
func (o *ClientProfile) SetGlobalHostname(v string) {
	o.GlobalHostname = &v
}

// GetExported returns the Exported field value if set, zero value otherwise.
func (o *ClientProfile) GetExported() time.Time {
	if o == nil || o.Exported == nil {
		var ret time.Time
		return ret
	}
	return *o.Exported
}

// GetExportedOk returns a tuple with the Exported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProfile) GetExportedOk() (*time.Time, bool) {
	if o == nil || o.Exported == nil {
		return nil, false
	}
	return o.Exported, true
}

// HasExported returns a boolean if a field has been set.
func (o *ClientProfile) HasExported() bool {
	if o != nil && o.Exported != nil {
		return true
	}

	return false
}

// SetExported gets a reference to the given time.Time and assigns it to the Exported field.
func (o *ClientProfile) SetExported(v time.Time) {
	o.Exported = &v
}

func (o ClientProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.SpaKeyName != nil {
		toSerialize["spaKeyName"] = o.SpaKeyName
	}
	if true {
		toSerialize["identityProviderName"] = o.IdentityProviderName
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.GlobalHostname != nil {
		toSerialize["globalHostname"] = o.GlobalHostname
	}
	if o.Exported != nil {
		toSerialize["exported"] = o.Exported
	}
	return json.Marshal(toSerialize)
}

type NullableClientProfile struct {
	value *ClientProfile
	isSet bool
}

func (v NullableClientProfile) Get() *ClientProfile {
	return v.value
}

func (v *NullableClientProfile) Set(val *ClientProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableClientProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableClientProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientProfile(val *ClientProfile) *NullableClientProfile {
	return &NullableClientProfile{value: val, isSet: true}
}

func (v NullableClientProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
