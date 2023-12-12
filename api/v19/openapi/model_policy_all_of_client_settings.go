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

// PolicyAllOfClientSettings Settings that admins can apply to the Client.
type PolicyAllOfClientSettings struct {
	// Enable Client Settings for this Policy.
	Enabled *bool `json:"enabled,omitempty"`
	// Show or hide Entitlement List on Client UI.
	EntitlementsList *string `json:"entitlementsList,omitempty"`
	// Set the Attention Level automatically on Client and hide the option. \"Show\" will leave option to the user.
	AttentionLevel *string `json:"attentionLevel,omitempty"`
	// Set the Autostart setting automatically on Client and hide the option. \"Show\" will leave option to the user.
	AutoStart *string `json:"autoStart,omitempty"`
	// Allow adding and removing profiles on Client.
	AddRemoveProfiles *string `json:"addRemoveProfiles,omitempty"`
	// Set the \"Keep me signed-in\" setting for credential providers automatically on Client and hide the option. \"Show\" will leave option to the user.
	KeepMeSignedIn *string `json:"keepMeSignedIn,omitempty"`
	// Set the \"SAML/Certificate auto sign-in\" setting automatically on Client and hide the option. \"Show\" will leave option the user.
	SamlAutoSignIn *string `json:"samlAutoSignIn,omitempty"`
	// Show or hide \"Quit\" on Client UI.
	Quit *string `json:"quit,omitempty"`
	// Show or hide \"Sign out\" on Client UI.
	SignOut *string `json:"signOut,omitempty"`
	// Show or hide \"Suspend\" feature on Client UI.
	Suspend *string `json:"suspend,omitempty"`
	// Show or hide the tooltips for new users on Client UI.
	NewUserOnboarding *string `json:"newUserOnboarding,omitempty"`
}

// NewPolicyAllOfClientSettings instantiates a new PolicyAllOfClientSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAllOfClientSettings() *PolicyAllOfClientSettings {
	this := PolicyAllOfClientSettings{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewPolicyAllOfClientSettingsWithDefaults instantiates a new PolicyAllOfClientSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAllOfClientSettingsWithDefaults() *PolicyAllOfClientSettings {
	this := PolicyAllOfClientSettings{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PolicyAllOfClientSettings) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOfClientSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PolicyAllOfClientSettings) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PolicyAllOfClientSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEntitlementsList returns the EntitlementsList field value if set, zero value otherwise.
func (o *PolicyAllOfClientSettings) GetEntitlementsList() string {
	if o == nil || o.EntitlementsList == nil {
		var ret string
		return ret
	}
	return *o.EntitlementsList
}

// GetEntitlementsListOk returns a tuple with the EntitlementsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOfClientSettings) GetEntitlementsListOk() (*string, bool) {
	if o == nil || o.EntitlementsList == nil {
		return nil, false
	}
	return o.EntitlementsList, true
}

// HasEntitlementsList returns a boolean if a field has been set.
func (o *PolicyAllOfClientSettings) HasEntitlementsList() bool {
	if o != nil && o.EntitlementsList != nil {
		return true
	}

	return false
}

// SetEntitlementsList gets a reference to the given string and assigns it to the EntitlementsList field.
func (o *PolicyAllOfClientSettings) SetEntitlementsList(v string) {
	o.EntitlementsList = &v
}

// GetAttentionLevel returns the AttentionLevel field value if set, zero value otherwise.
func (o *PolicyAllOfClientSettings) GetAttentionLevel() string {
	if o == nil || o.AttentionLevel == nil {
		var ret string
		return ret
	}
	return *o.AttentionLevel
}

// GetAttentionLevelOk returns a tuple with the AttentionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOfClientSettings) GetAttentionLevelOk() (*string, bool) {
	if o == nil || o.AttentionLevel == nil {
		return nil, false
	}
	return o.AttentionLevel, true
}

// HasAttentionLevel returns a boolean if a field has been set.
func (o *PolicyAllOfClientSettings) HasAttentionLevel() bool {
	if o != nil && o.AttentionLevel != nil {
		return true
	}

	return false
}

// SetAttentionLevel gets a reference to the given string and assigns it to the AttentionLevel field.
func (o *PolicyAllOfClientSettings) SetAttentionLevel(v string) {
	o.AttentionLevel = &v
}

// GetAutoStart returns the AutoStart field value if set, zero value otherwise.
func (o *PolicyAllOfClientSettings) GetAutoStart() string {
	if o == nil || o.AutoStart == nil {
		var ret string
		return ret
	}
	return *o.AutoStart
}

// GetAutoStartOk returns a tuple with the AutoStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOfClientSettings) GetAutoStartOk() (*string, bool) {
	if o == nil || o.AutoStart == nil {
		return nil, false
	}
	return o.AutoStart, true
}

// HasAutoStart returns a boolean if a field has been set.
func (o *PolicyAllOfClientSettings) HasAutoStart() bool {
	if o != nil && o.AutoStart != nil {
		return true
	}

	return false
}

// SetAutoStart gets a reference to the given string and assigns it to the AutoStart field.
func (o *PolicyAllOfClientSettings) SetAutoStart(v string) {
	o.AutoStart = &v
}

// GetAddRemoveProfiles returns the AddRemoveProfiles field value if set, zero value otherwise.
func (o *PolicyAllOfClientSettings) GetAddRemoveProfiles() string {
	if o == nil || o.AddRemoveProfiles == nil {
		var ret string
		return ret
	}
	return *o.AddRemoveProfiles
}

// GetAddRemoveProfilesOk returns a tuple with the AddRemoveProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOfClientSettings) GetAddRemoveProfilesOk() (*string, bool) {
	if o == nil || o.AddRemoveProfiles == nil {
		return nil, false
	}
	return o.AddRemoveProfiles, true
}

// HasAddRemoveProfiles returns a boolean if a field has been set.
func (o *PolicyAllOfClientSettings) HasAddRemoveProfiles() bool {
	if o != nil && o.AddRemoveProfiles != nil {
		return true
	}

	return false
}

// SetAddRemoveProfiles gets a reference to the given string and assigns it to the AddRemoveProfiles field.
func (o *PolicyAllOfClientSettings) SetAddRemoveProfiles(v string) {
	o.AddRemoveProfiles = &v
}

// GetKeepMeSignedIn returns the KeepMeSignedIn field value if set, zero value otherwise.
func (o *PolicyAllOfClientSettings) GetKeepMeSignedIn() string {
	if o == nil || o.KeepMeSignedIn == nil {
		var ret string
		return ret
	}
	return *o.KeepMeSignedIn
}

// GetKeepMeSignedInOk returns a tuple with the KeepMeSignedIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOfClientSettings) GetKeepMeSignedInOk() (*string, bool) {
	if o == nil || o.KeepMeSignedIn == nil {
		return nil, false
	}
	return o.KeepMeSignedIn, true
}

// HasKeepMeSignedIn returns a boolean if a field has been set.
func (o *PolicyAllOfClientSettings) HasKeepMeSignedIn() bool {
	if o != nil && o.KeepMeSignedIn != nil {
		return true
	}

	return false
}

// SetKeepMeSignedIn gets a reference to the given string and assigns it to the KeepMeSignedIn field.
func (o *PolicyAllOfClientSettings) SetKeepMeSignedIn(v string) {
	o.KeepMeSignedIn = &v
}

// GetSamlAutoSignIn returns the SamlAutoSignIn field value if set, zero value otherwise.
func (o *PolicyAllOfClientSettings) GetSamlAutoSignIn() string {
	if o == nil || o.SamlAutoSignIn == nil {
		var ret string
		return ret
	}
	return *o.SamlAutoSignIn
}

// GetSamlAutoSignInOk returns a tuple with the SamlAutoSignIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOfClientSettings) GetSamlAutoSignInOk() (*string, bool) {
	if o == nil || o.SamlAutoSignIn == nil {
		return nil, false
	}
	return o.SamlAutoSignIn, true
}

// HasSamlAutoSignIn returns a boolean if a field has been set.
func (o *PolicyAllOfClientSettings) HasSamlAutoSignIn() bool {
	if o != nil && o.SamlAutoSignIn != nil {
		return true
	}

	return false
}

// SetSamlAutoSignIn gets a reference to the given string and assigns it to the SamlAutoSignIn field.
func (o *PolicyAllOfClientSettings) SetSamlAutoSignIn(v string) {
	o.SamlAutoSignIn = &v
}

// GetQuit returns the Quit field value if set, zero value otherwise.
func (o *PolicyAllOfClientSettings) GetQuit() string {
	if o == nil || o.Quit == nil {
		var ret string
		return ret
	}
	return *o.Quit
}

// GetQuitOk returns a tuple with the Quit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOfClientSettings) GetQuitOk() (*string, bool) {
	if o == nil || o.Quit == nil {
		return nil, false
	}
	return o.Quit, true
}

// HasQuit returns a boolean if a field has been set.
func (o *PolicyAllOfClientSettings) HasQuit() bool {
	if o != nil && o.Quit != nil {
		return true
	}

	return false
}

// SetQuit gets a reference to the given string and assigns it to the Quit field.
func (o *PolicyAllOfClientSettings) SetQuit(v string) {
	o.Quit = &v
}

// GetSignOut returns the SignOut field value if set, zero value otherwise.
func (o *PolicyAllOfClientSettings) GetSignOut() string {
	if o == nil || o.SignOut == nil {
		var ret string
		return ret
	}
	return *o.SignOut
}

// GetSignOutOk returns a tuple with the SignOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOfClientSettings) GetSignOutOk() (*string, bool) {
	if o == nil || o.SignOut == nil {
		return nil, false
	}
	return o.SignOut, true
}

// HasSignOut returns a boolean if a field has been set.
func (o *PolicyAllOfClientSettings) HasSignOut() bool {
	if o != nil && o.SignOut != nil {
		return true
	}

	return false
}

// SetSignOut gets a reference to the given string and assigns it to the SignOut field.
func (o *PolicyAllOfClientSettings) SetSignOut(v string) {
	o.SignOut = &v
}

// GetSuspend returns the Suspend field value if set, zero value otherwise.
func (o *PolicyAllOfClientSettings) GetSuspend() string {
	if o == nil || o.Suspend == nil {
		var ret string
		return ret
	}
	return *o.Suspend
}

// GetSuspendOk returns a tuple with the Suspend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOfClientSettings) GetSuspendOk() (*string, bool) {
	if o == nil || o.Suspend == nil {
		return nil, false
	}
	return o.Suspend, true
}

// HasSuspend returns a boolean if a field has been set.
func (o *PolicyAllOfClientSettings) HasSuspend() bool {
	if o != nil && o.Suspend != nil {
		return true
	}

	return false
}

// SetSuspend gets a reference to the given string and assigns it to the Suspend field.
func (o *PolicyAllOfClientSettings) SetSuspend(v string) {
	o.Suspend = &v
}

// GetNewUserOnboarding returns the NewUserOnboarding field value if set, zero value otherwise.
func (o *PolicyAllOfClientSettings) GetNewUserOnboarding() string {
	if o == nil || o.NewUserOnboarding == nil {
		var ret string
		return ret
	}
	return *o.NewUserOnboarding
}

// GetNewUserOnboardingOk returns a tuple with the NewUserOnboarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAllOfClientSettings) GetNewUserOnboardingOk() (*string, bool) {
	if o == nil || o.NewUserOnboarding == nil {
		return nil, false
	}
	return o.NewUserOnboarding, true
}

// HasNewUserOnboarding returns a boolean if a field has been set.
func (o *PolicyAllOfClientSettings) HasNewUserOnboarding() bool {
	if o != nil && o.NewUserOnboarding != nil {
		return true
	}

	return false
}

// SetNewUserOnboarding gets a reference to the given string and assigns it to the NewUserOnboarding field.
func (o *PolicyAllOfClientSettings) SetNewUserOnboarding(v string) {
	o.NewUserOnboarding = &v
}

func (o PolicyAllOfClientSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.EntitlementsList != nil {
		toSerialize["entitlementsList"] = o.EntitlementsList
	}
	if o.AttentionLevel != nil {
		toSerialize["attentionLevel"] = o.AttentionLevel
	}
	if o.AutoStart != nil {
		toSerialize["autoStart"] = o.AutoStart
	}
	if o.AddRemoveProfiles != nil {
		toSerialize["addRemoveProfiles"] = o.AddRemoveProfiles
	}
	if o.KeepMeSignedIn != nil {
		toSerialize["keepMeSignedIn"] = o.KeepMeSignedIn
	}
	if o.SamlAutoSignIn != nil {
		toSerialize["samlAutoSignIn"] = o.SamlAutoSignIn
	}
	if o.Quit != nil {
		toSerialize["quit"] = o.Quit
	}
	if o.SignOut != nil {
		toSerialize["signOut"] = o.SignOut
	}
	if o.Suspend != nil {
		toSerialize["suspend"] = o.Suspend
	}
	if o.NewUserOnboarding != nil {
		toSerialize["newUserOnboarding"] = o.NewUserOnboarding
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyAllOfClientSettings struct {
	value *PolicyAllOfClientSettings
	isSet bool
}

func (v NullablePolicyAllOfClientSettings) Get() *PolicyAllOfClientSettings {
	return v.value
}

func (v *NullablePolicyAllOfClientSettings) Set(val *PolicyAllOfClientSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAllOfClientSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAllOfClientSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAllOfClientSettings(val *PolicyAllOfClientSettings) *NullablePolicyAllOfClientSettings {
	return &NullablePolicyAllOfClientSettings{value: val, isSet: true}
}

func (v NullablePolicyAllOfClientSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAllOfClientSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
