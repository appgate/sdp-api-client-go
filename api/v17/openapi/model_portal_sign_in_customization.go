/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v17+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 17.4
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PortalSignInCustomization Visual customizations to make on the Portal sign-in page.
type PortalSignInCustomization struct {
	// Changes the background color on the sign-in page. In hexadecimal format.
	BackgroundColor *string `json:"backgroundColor,omitempty"`
	// Changes the background image on the sign-in page. Must be in PNG, JPEG or GIF format.
	BackgroundImage *string `json:"backgroundImage,omitempty"`
	// Changes the logo on the sign-in page. Must be in PNG, JPEG or GIF format.
	Logo *string `json:"logo,omitempty"`
	// Adds a text to the sign-in page.
	Text *string `json:"text,omitempty"`
	// Changes the text color on the sign-in page. In hexadecimal format.
	TextColor *string `json:"textColor,omitempty"`
	// If enabled and the user lands on the Portal sign-in page by entering an endpoint URL on the browser, it will be redirected to the endpoint automatically after successfully signing in instead of the Portal Client overview page.
	AutoRedirect *bool `json:"autoRedirect,omitempty"`
}

// NewPortalSignInCustomization instantiates a new PortalSignInCustomization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortalSignInCustomization() *PortalSignInCustomization {
	this := PortalSignInCustomization{}
	return &this
}

// NewPortalSignInCustomizationWithDefaults instantiates a new PortalSignInCustomization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortalSignInCustomizationWithDefaults() *PortalSignInCustomization {
	this := PortalSignInCustomization{}
	return &this
}

// GetBackgroundColor returns the BackgroundColor field value if set, zero value otherwise.
func (o *PortalSignInCustomization) GetBackgroundColor() string {
	if o == nil || o.BackgroundColor == nil {
		var ret string
		return ret
	}
	return *o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortalSignInCustomization) GetBackgroundColorOk() (*string, bool) {
	if o == nil || o.BackgroundColor == nil {
		return nil, false
	}
	return o.BackgroundColor, true
}

// HasBackgroundColor returns a boolean if a field has been set.
func (o *PortalSignInCustomization) HasBackgroundColor() bool {
	if o != nil && o.BackgroundColor != nil {
		return true
	}

	return false
}

// SetBackgroundColor gets a reference to the given string and assigns it to the BackgroundColor field.
func (o *PortalSignInCustomization) SetBackgroundColor(v string) {
	o.BackgroundColor = &v
}

// GetBackgroundImage returns the BackgroundImage field value if set, zero value otherwise.
func (o *PortalSignInCustomization) GetBackgroundImage() string {
	if o == nil || o.BackgroundImage == nil {
		var ret string
		return ret
	}
	return *o.BackgroundImage
}

// GetBackgroundImageOk returns a tuple with the BackgroundImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortalSignInCustomization) GetBackgroundImageOk() (*string, bool) {
	if o == nil || o.BackgroundImage == nil {
		return nil, false
	}
	return o.BackgroundImage, true
}

// HasBackgroundImage returns a boolean if a field has been set.
func (o *PortalSignInCustomization) HasBackgroundImage() bool {
	if o != nil && o.BackgroundImage != nil {
		return true
	}

	return false
}

// SetBackgroundImage gets a reference to the given string and assigns it to the BackgroundImage field.
func (o *PortalSignInCustomization) SetBackgroundImage(v string) {
	o.BackgroundImage = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *PortalSignInCustomization) GetLogo() string {
	if o == nil || o.Logo == nil {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortalSignInCustomization) GetLogoOk() (*string, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *PortalSignInCustomization) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *PortalSignInCustomization) SetLogo(v string) {
	o.Logo = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *PortalSignInCustomization) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortalSignInCustomization) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *PortalSignInCustomization) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *PortalSignInCustomization) SetText(v string) {
	o.Text = &v
}

// GetTextColor returns the TextColor field value if set, zero value otherwise.
func (o *PortalSignInCustomization) GetTextColor() string {
	if o == nil || o.TextColor == nil {
		var ret string
		return ret
	}
	return *o.TextColor
}

// GetTextColorOk returns a tuple with the TextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortalSignInCustomization) GetTextColorOk() (*string, bool) {
	if o == nil || o.TextColor == nil {
		return nil, false
	}
	return o.TextColor, true
}

// HasTextColor returns a boolean if a field has been set.
func (o *PortalSignInCustomization) HasTextColor() bool {
	if o != nil && o.TextColor != nil {
		return true
	}

	return false
}

// SetTextColor gets a reference to the given string and assigns it to the TextColor field.
func (o *PortalSignInCustomization) SetTextColor(v string) {
	o.TextColor = &v
}

// GetAutoRedirect returns the AutoRedirect field value if set, zero value otherwise.
func (o *PortalSignInCustomization) GetAutoRedirect() bool {
	if o == nil || o.AutoRedirect == nil {
		var ret bool
		return ret
	}
	return *o.AutoRedirect
}

// GetAutoRedirectOk returns a tuple with the AutoRedirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortalSignInCustomization) GetAutoRedirectOk() (*bool, bool) {
	if o == nil || o.AutoRedirect == nil {
		return nil, false
	}
	return o.AutoRedirect, true
}

// HasAutoRedirect returns a boolean if a field has been set.
func (o *PortalSignInCustomization) HasAutoRedirect() bool {
	if o != nil && o.AutoRedirect != nil {
		return true
	}

	return false
}

// SetAutoRedirect gets a reference to the given bool and assigns it to the AutoRedirect field.
func (o *PortalSignInCustomization) SetAutoRedirect(v bool) {
	o.AutoRedirect = &v
}

func (o PortalSignInCustomization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackgroundColor != nil {
		toSerialize["backgroundColor"] = o.BackgroundColor
	}
	if o.BackgroundImage != nil {
		toSerialize["backgroundImage"] = o.BackgroundImage
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.TextColor != nil {
		toSerialize["textColor"] = o.TextColor
	}
	if o.AutoRedirect != nil {
		toSerialize["autoRedirect"] = o.AutoRedirect
	}
	return json.Marshal(toSerialize)
}

type NullablePortalSignInCustomization struct {
	value *PortalSignInCustomization
	isSet bool
}

func (v NullablePortalSignInCustomization) Get() *PortalSignInCustomization {
	return v.value
}

func (v *NullablePortalSignInCustomization) Set(val *PortalSignInCustomization) {
	v.value = val
	v.isSet = true
}

func (v NullablePortalSignInCustomization) IsSet() bool {
	return v.isSet
}

func (v *NullablePortalSignInCustomization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortalSignInCustomization(val *PortalSignInCustomization) *NullablePortalSignInCustomization {
	return &NullablePortalSignInCustomization{value: val, isSet: true}
}

func (v NullablePortalSignInCustomization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortalSignInCustomization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
