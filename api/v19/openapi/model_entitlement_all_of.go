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

// EntitlementAllOf Represents an Entitlement.
type EntitlementAllOf struct {
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

// NewEntitlementAllOf instantiates a new EntitlementAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementAllOf(site string, conditions []string, actions []EntitlementAllOfActions) *EntitlementAllOf {
	this := EntitlementAllOf{}
	var disabled bool = false
	this.Disabled = &disabled
	this.Site = site
	var conditionLogic string = "and"
	this.ConditionLogic = &conditionLogic
	this.Conditions = conditions
	this.Actions = actions
	return &this
}

// NewEntitlementAllOfWithDefaults instantiates a new EntitlementAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementAllOfWithDefaults() *EntitlementAllOf {
	this := EntitlementAllOf{}
	var disabled bool = false
	this.Disabled = &disabled
	var conditionLogic string = "and"
	this.ConditionLogic = &conditionLogic
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *EntitlementAllOf) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAllOf) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *EntitlementAllOf) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *EntitlementAllOf) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetSite returns the Site field value
func (o *EntitlementAllOf) GetSite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *EntitlementAllOf) GetSiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Site, true
}

// SetSite sets field value
func (o *EntitlementAllOf) SetSite(v string) {
	o.Site = v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *EntitlementAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *EntitlementAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *EntitlementAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRiskSensitivity returns the RiskSensitivity field value if set, zero value otherwise.
func (o *EntitlementAllOf) GetRiskSensitivity() string {
	if o == nil || o.RiskSensitivity == nil {
		var ret string
		return ret
	}
	return *o.RiskSensitivity
}

// GetRiskSensitivityOk returns a tuple with the RiskSensitivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAllOf) GetRiskSensitivityOk() (*string, bool) {
	if o == nil || o.RiskSensitivity == nil {
		return nil, false
	}
	return o.RiskSensitivity, true
}

// HasRiskSensitivity returns a boolean if a field has been set.
func (o *EntitlementAllOf) HasRiskSensitivity() bool {
	if o != nil && o.RiskSensitivity != nil {
		return true
	}

	return false
}

// SetRiskSensitivity gets a reference to the given string and assigns it to the RiskSensitivity field.
func (o *EntitlementAllOf) SetRiskSensitivity(v string) {
	o.RiskSensitivity = &v
}

// GetConditionLogic returns the ConditionLogic field value if set, zero value otherwise.
func (o *EntitlementAllOf) GetConditionLogic() string {
	if o == nil || o.ConditionLogic == nil {
		var ret string
		return ret
	}
	return *o.ConditionLogic
}

// GetConditionLogicOk returns a tuple with the ConditionLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAllOf) GetConditionLogicOk() (*string, bool) {
	if o == nil || o.ConditionLogic == nil {
		return nil, false
	}
	return o.ConditionLogic, true
}

// HasConditionLogic returns a boolean if a field has been set.
func (o *EntitlementAllOf) HasConditionLogic() bool {
	if o != nil && o.ConditionLogic != nil {
		return true
	}

	return false
}

// SetConditionLogic gets a reference to the given string and assigns it to the ConditionLogic field.
func (o *EntitlementAllOf) SetConditionLogic(v string) {
	o.ConditionLogic = &v
}

// GetConditions returns the Conditions field value
func (o *EntitlementAllOf) GetConditions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *EntitlementAllOf) GetConditionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *EntitlementAllOf) SetConditions(v []string) {
	o.Conditions = v
}

// GetActions returns the Actions field value
func (o *EntitlementAllOf) GetActions() []EntitlementAllOfActions {
	if o == nil {
		var ret []EntitlementAllOfActions
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *EntitlementAllOf) GetActionsOk() ([]EntitlementAllOfActions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *EntitlementAllOf) SetActions(v []EntitlementAllOfActions) {
	o.Actions = v
}

// GetAppShortcuts returns the AppShortcuts field value if set, zero value otherwise.
func (o *EntitlementAllOf) GetAppShortcuts() []AppShortcut {
	if o == nil || o.AppShortcuts == nil {
		var ret []AppShortcut
		return ret
	}
	return o.AppShortcuts
}

// GetAppShortcutsOk returns a tuple with the AppShortcuts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAllOf) GetAppShortcutsOk() ([]AppShortcut, bool) {
	if o == nil || o.AppShortcuts == nil {
		return nil, false
	}
	return o.AppShortcuts, true
}

// HasAppShortcuts returns a boolean if a field has been set.
func (o *EntitlementAllOf) HasAppShortcuts() bool {
	if o != nil && o.AppShortcuts != nil {
		return true
	}

	return false
}

// SetAppShortcuts gets a reference to the given []AppShortcut and assigns it to the AppShortcuts field.
func (o *EntitlementAllOf) SetAppShortcuts(v []AppShortcut) {
	o.AppShortcuts = v
}

// GetAppShortcutScripts returns the AppShortcutScripts field value if set, zero value otherwise.
func (o *EntitlementAllOf) GetAppShortcutScripts() []string {
	if o == nil || o.AppShortcutScripts == nil {
		var ret []string
		return ret
	}
	return o.AppShortcutScripts
}

// GetAppShortcutScriptsOk returns a tuple with the AppShortcutScripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAllOf) GetAppShortcutScriptsOk() ([]string, bool) {
	if o == nil || o.AppShortcutScripts == nil {
		return nil, false
	}
	return o.AppShortcutScripts, true
}

// HasAppShortcutScripts returns a boolean if a field has been set.
func (o *EntitlementAllOf) HasAppShortcutScripts() bool {
	if o != nil && o.AppShortcutScripts != nil {
		return true
	}

	return false
}

// SetAppShortcutScripts gets a reference to the given []string and assigns it to the AppShortcutScripts field.
func (o *EntitlementAllOf) SetAppShortcutScripts(v []string) {
	o.AppShortcutScripts = v
}

func (o EntitlementAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableEntitlementAllOf struct {
	value *EntitlementAllOf
	isSet bool
}

func (v NullableEntitlementAllOf) Get() *EntitlementAllOf {
	return v.value
}

func (v *NullableEntitlementAllOf) Set(val *EntitlementAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementAllOf(val *EntitlementAllOf) *NullableEntitlementAllOf {
	return &NullableEntitlementAllOf{value: val, isSet: true}
}

func (v NullableEntitlementAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
