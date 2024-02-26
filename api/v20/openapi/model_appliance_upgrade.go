/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v20+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 20.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceUpgrade Appliance Upgrade request.
type ApplianceUpgrade struct {
	// Disclaimer of Warranty - internal use only. This is not a supported property. Use the development keyring to verify the upgrade image.
	DevKeyring *bool `json:"devKeyring,omitempty"`
	// The URL for the Appliance the download the Upgrade image from. The URL may be a public HTTP URL or it could be a file uploaded to the Controller. See \"files\" APIs for uploading to Controller. In order to use a Controller based file, set this field to \"controller://\\<controller-peer-hostname:port\\>/{filename}\". The Appliance will authenticate itself to the Controller and download the file.
	ImageUrl string `json:"imageUrl"`
}

// NewApplianceUpgrade instantiates a new ApplianceUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceUpgrade(imageUrl string) *ApplianceUpgrade {
	this := ApplianceUpgrade{}
	this.ImageUrl = imageUrl
	return &this
}

// NewApplianceUpgradeWithDefaults instantiates a new ApplianceUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceUpgradeWithDefaults() *ApplianceUpgrade {
	this := ApplianceUpgrade{}
	return &this
}

// GetDevKeyring returns the DevKeyring field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetDevKeyring() bool {
	if o == nil || o.DevKeyring == nil {
		var ret bool
		return ret
	}
	return *o.DevKeyring
}

// GetDevKeyringOk returns a tuple with the DevKeyring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetDevKeyringOk() (*bool, bool) {
	if o == nil || o.DevKeyring == nil {
		return nil, false
	}
	return o.DevKeyring, true
}

// HasDevKeyring returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasDevKeyring() bool {
	if o != nil && o.DevKeyring != nil {
		return true
	}

	return false
}

// SetDevKeyring gets a reference to the given bool and assigns it to the DevKeyring field.
func (o *ApplianceUpgrade) SetDevKeyring(v bool) {
	o.DevKeyring = &v
}

// GetImageUrl returns the ImageUrl field value
func (o *ApplianceUpgrade) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *ApplianceUpgrade) SetImageUrl(v string) {
	o.ImageUrl = v
}

func (o ApplianceUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DevKeyring != nil {
		toSerialize["devKeyring"] = o.DevKeyring
	}
	if true {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceUpgrade struct {
	value *ApplianceUpgrade
	isSet bool
}

func (v NullableApplianceUpgrade) Get() *ApplianceUpgrade {
	return v.value
}

func (v *NullableApplianceUpgrade) Set(val *ApplianceUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceUpgrade(val *ApplianceUpgrade) *NullableApplianceUpgrade {
	return &NullableApplianceUpgrade{value: val, isSet: true}
}

func (v NullableApplianceUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
