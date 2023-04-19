/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v18+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 18.5
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Entitlement struct for Entitlement
type Entitlement struct {
	// ID of the object.
	Id *string `json:"id,omitempty"`
	// Name of the object.
	Name string `json:"name"`
	// Notes for the object. Used for documentation purposes.
	Notes *string `json:"notes,omitempty"`
	// Create date.
	Created *time.Time `json:"created,omitempty"`
	// Last update date.
	Updated *time.Time `json:"updated,omitempty"`
	// Array of tags.
	Tags []string `json:"tags,omitempty"`
	// If true, the Entitlement will be disregarded during authorization.
	Disabled *bool `json:"disabled,omitempty"`
	// ID of the Site for this Entitlement.
	Site string `json:"site"`
	// Name of the Site for this Entitlement. For convenience only.
	SiteName *string `json:"siteName,omitempty"`
	// Generate Conditions for the Entitlement based on the Risk Model. Cannot be combined with other Conditions.
	RiskSensitivity *string `json:"riskSensitivity,omitempty"`
	// Whether all the Conditions must succeed to have access to this Entitlement or just one.
	ConditionLogic *string `json:"conditionLogic,omitempty"`
	// List of Condition IDs applies to this Entitlement.
	Conditions []string `json:"conditions"`
	// List of all IP Access actions in this Entitlement.
	Actions []EntitlementAllOfActions `json:"actions"`
	// Array of App Shortcuts.
	AppShortcuts []AppShortcut `json:"appShortcuts,omitempty"`
	// List of Entitlement Script IDs used for creating App Shortcuts dynamically.
	AppShortcutScripts []string `json:"appShortcutScripts,omitempty"`
}

// NewEntitlement instantiates a new Entitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlement(name string, site string, conditions []string, actions []EntitlementAllOfActions) *Entitlement {
	this := Entitlement{}
	this.Name = name
	var disabled bool = false
	this.Disabled = &disabled
	this.Site = site
	var conditionLogic string = "and"
	this.ConditionLogic = &conditionLogic
	this.Conditions = conditions
	this.Actions = actions
	return &this
}

// NewEntitlementWithDefaults instantiates a new Entitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementWithDefaults() *Entitlement {
	this := Entitlement{}
	var disabled bool = false
	this.Disabled = &disabled
	var conditionLogic string = "and"
	this.ConditionLogic = &conditionLogic
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Entitlement) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Entitlement) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Entitlement) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Entitlement) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Entitlement) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *Entitlement) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *Entitlement) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *Entitlement) SetNotes(v string) {
	o.Notes = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Entitlement) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Entitlement) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Entitlement) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Entitlement) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Entitlement) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *Entitlement) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Entitlement) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Entitlement) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Entitlement) SetTags(v []string) {
	o.Tags = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *Entitlement) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *Entitlement) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *Entitlement) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetSite returns the Site field value
func (o *Entitlement) GetSite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetSiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Site, true
}

// SetSite sets field value
func (o *Entitlement) SetSite(v string) {
	o.Site = v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *Entitlement) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *Entitlement) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *Entitlement) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRiskSensitivity returns the RiskSensitivity field value if set, zero value otherwise.
func (o *Entitlement) GetRiskSensitivity() string {
	if o == nil || o.RiskSensitivity == nil {
		var ret string
		return ret
	}
	return *o.RiskSensitivity
}

// GetRiskSensitivityOk returns a tuple with the RiskSensitivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetRiskSensitivityOk() (*string, bool) {
	if o == nil || o.RiskSensitivity == nil {
		return nil, false
	}
	return o.RiskSensitivity, true
}

// HasRiskSensitivity returns a boolean if a field has been set.
func (o *Entitlement) HasRiskSensitivity() bool {
	if o != nil && o.RiskSensitivity != nil {
		return true
	}

	return false
}

// SetRiskSensitivity gets a reference to the given string and assigns it to the RiskSensitivity field.
func (o *Entitlement) SetRiskSensitivity(v string) {
	o.RiskSensitivity = &v
}

// GetConditionLogic returns the ConditionLogic field value if set, zero value otherwise.
func (o *Entitlement) GetConditionLogic() string {
	if o == nil || o.ConditionLogic == nil {
		var ret string
		return ret
	}
	return *o.ConditionLogic
}

// GetConditionLogicOk returns a tuple with the ConditionLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetConditionLogicOk() (*string, bool) {
	if o == nil || o.ConditionLogic == nil {
		return nil, false
	}
	return o.ConditionLogic, true
}

// HasConditionLogic returns a boolean if a field has been set.
func (o *Entitlement) HasConditionLogic() bool {
	if o != nil && o.ConditionLogic != nil {
		return true
	}

	return false
}

// SetConditionLogic gets a reference to the given string and assigns it to the ConditionLogic field.
func (o *Entitlement) SetConditionLogic(v string) {
	o.ConditionLogic = &v
}

// GetConditions returns the Conditions field value
func (o *Entitlement) GetConditions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetConditionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *Entitlement) SetConditions(v []string) {
	o.Conditions = v
}

// GetActions returns the Actions field value
func (o *Entitlement) GetActions() []EntitlementAllOfActions {
	if o == nil {
		var ret []EntitlementAllOfActions
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetActionsOk() ([]EntitlementAllOfActions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *Entitlement) SetActions(v []EntitlementAllOfActions) {
	o.Actions = v
}

// GetAppShortcuts returns the AppShortcuts field value if set, zero value otherwise.
func (o *Entitlement) GetAppShortcuts() []AppShortcut {
	if o == nil || o.AppShortcuts == nil {
		var ret []AppShortcut
		return ret
	}
	return o.AppShortcuts
}

// GetAppShortcutsOk returns a tuple with the AppShortcuts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetAppShortcutsOk() ([]AppShortcut, bool) {
	if o == nil || o.AppShortcuts == nil {
		return nil, false
	}
	return o.AppShortcuts, true
}

// HasAppShortcuts returns a boolean if a field has been set.
func (o *Entitlement) HasAppShortcuts() bool {
	if o != nil && o.AppShortcuts != nil {
		return true
	}

	return false
}

// SetAppShortcuts gets a reference to the given []AppShortcut and assigns it to the AppShortcuts field.
func (o *Entitlement) SetAppShortcuts(v []AppShortcut) {
	o.AppShortcuts = v
}

// GetAppShortcutScripts returns the AppShortcutScripts field value if set, zero value otherwise.
func (o *Entitlement) GetAppShortcutScripts() []string {
	if o == nil || o.AppShortcutScripts == nil {
		var ret []string
		return ret
	}
	return o.AppShortcutScripts
}

// GetAppShortcutScriptsOk returns a tuple with the AppShortcutScripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlement) GetAppShortcutScriptsOk() ([]string, bool) {
	if o == nil || o.AppShortcutScripts == nil {
		return nil, false
	}
	return o.AppShortcutScripts, true
}

// HasAppShortcutScripts returns a boolean if a field has been set.
func (o *Entitlement) HasAppShortcutScripts() bool {
	if o != nil && o.AppShortcutScripts != nil {
		return true
	}

	return false
}

// SetAppShortcutScripts gets a reference to the given []string and assigns it to the AppShortcutScripts field.
func (o *Entitlement) SetAppShortcutScripts(v []string) {
	o.AppShortcutScripts = v
}

func (o Entitlement) MarshalJSON() ([]byte, error) {
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
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if true {
		toSerialize["site"] = o.Site
	}
	if o.SiteName != nil {
		toSerialize["siteName"] = o.SiteName
	}
	if o.RiskSensitivity != nil {
		toSerialize["riskSensitivity"] = o.RiskSensitivity
	}
	if o.ConditionLogic != nil {
		toSerialize["conditionLogic"] = o.ConditionLogic
	}
	if true {
		toSerialize["conditions"] = o.Conditions
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	if o.AppShortcuts != nil {
		toSerialize["appShortcuts"] = o.AppShortcuts
	}
	if o.AppShortcutScripts != nil {
		toSerialize["appShortcutScripts"] = o.AppShortcutScripts
	}
	return json.Marshal(toSerialize)
}

type NullableEntitlement struct {
	value *Entitlement
	isSet bool
}

func (v NullableEntitlement) Get() *Entitlement {
	return v.value
}

func (v *NullableEntitlement) Set(val *Entitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlement(val *Entitlement) *NullableEntitlement {
	return &NullableEntitlement{value: val, isSet: true}
}

func (v NullableEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
