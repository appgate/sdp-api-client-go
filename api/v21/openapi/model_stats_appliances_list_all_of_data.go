/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v21+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 21.0
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StatsAppliancesListAllOfData An active Appliance status details.
type StatsAppliancesListAllOfData struct {
	// ID of the Appliance.
	Id *string `json:"id,omitempty"`
	// Name of the Appliance.
	Name *string `json:"name,omitempty"`
	// Name of the Site for this Appliance.
	SiteName *string `json:"siteName,omitempty"`
	// Tags of the Appliance.
	Tags []string `json:"tags,omitempty"`
	// Whether the Appliance is reachable by the Controller or not. Deprecated as of 5.1. \"status\" field will be set to \"offline\" instead.
	// Deprecated
	Online *bool `json:"online,omitempty"`
	// The Appliance build version.
	Version *string `json:"version,omitempty"`
	// State of the Appliance. For internal use.
	State *string `json:"state,omitempty"`
	// The volume number in use.
	VolumeNumber *float32 `json:"volumeNumber,omitempty"`
	// Aggregated status of the Appliance.
	Status *string `json:"status,omitempty"`
	// Comma-separated list of functions enabled on this Appliance in short format.
	Function     *string                    `json:"function,omitempty"`
	Controller   *ControllerRole            `json:"controller,omitempty"`
	LogServer    *ApplianceRole             `json:"logServer,omitempty"`
	LogForwarder *ApplianceRole             `json:"logForwarder,omitempty"`
	Gateway      *ApplianceWithSessionsRole `json:"gateway,omitempty"`
	Connector    *ApplianceWithSessionsRole `json:"connector,omitempty"`
	Portal       *ApplianceWithSessionsRole `json:"portal,omitempty"`
	Appliance    *ApplianceBaseRole         `json:"appliance,omitempty"`
	// Number of sessions on Gateway, Portal or Connector.
	NumberOfSessions *int32 `json:"numberOfSessions,omitempty"`
	// Current CPU utilization % on the Appliance.
	Cpu *float32 `json:"cpu,omitempty"`
	// Current memory reserved % on the Appliance.
	Memory *float32 `json:"memory,omitempty"`
	// Current disk usage % on the Appliance.
	Disk     *float32                          `json:"disk,omitempty"`
	DiskInfo *StatsAppliancesListAllOfDiskInfo `json:"diskInfo,omitempty"`
	Network  *StatsAppliancesListAllOfNetwork  `json:"network,omitempty"`
	Upgrade  *StatsAppliancesListAllOfUpgrade  `json:"upgrade,omitempty"`
	// The name of the Appliance Customization if one is assigned to this Appliance.
	CustomizationName *string `json:"customizationName,omitempty"`
}

// NewStatsAppliancesListAllOfData instantiates a new StatsAppliancesListAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsAppliancesListAllOfData() *StatsAppliancesListAllOfData {
	this := StatsAppliancesListAllOfData{}
	return &this
}

// NewStatsAppliancesListAllOfDataWithDefaults instantiates a new StatsAppliancesListAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsAppliancesListAllOfDataWithDefaults() *StatsAppliancesListAllOfData {
	this := StatsAppliancesListAllOfData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StatsAppliancesListAllOfData) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StatsAppliancesListAllOfData) SetName(v string) {
	o.Name = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *StatsAppliancesListAllOfData) SetSiteName(v string) {
	o.SiteName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *StatsAppliancesListAllOfData) SetTags(v []string) {
	o.Tags = v
}

// GetOnline returns the Online field value if set, zero value otherwise.
// Deprecated
func (o *StatsAppliancesListAllOfData) GetOnline() bool {
	if o == nil || o.Online == nil {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *StatsAppliancesListAllOfData) GetOnlineOk() (*bool, bool) {
	if o == nil || o.Online == nil {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasOnline() bool {
	if o != nil && o.Online != nil {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
// Deprecated
func (o *StatsAppliancesListAllOfData) SetOnline(v bool) {
	o.Online = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StatsAppliancesListAllOfData) SetVersion(v string) {
	o.Version = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StatsAppliancesListAllOfData) SetState(v string) {
	o.State = &v
}

// GetVolumeNumber returns the VolumeNumber field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetVolumeNumber() float32 {
	if o == nil || o.VolumeNumber == nil {
		var ret float32
		return ret
	}
	return *o.VolumeNumber
}

// GetVolumeNumberOk returns a tuple with the VolumeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetVolumeNumberOk() (*float32, bool) {
	if o == nil || o.VolumeNumber == nil {
		return nil, false
	}
	return o.VolumeNumber, true
}

// HasVolumeNumber returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasVolumeNumber() bool {
	if o != nil && o.VolumeNumber != nil {
		return true
	}

	return false
}

// SetVolumeNumber gets a reference to the given float32 and assigns it to the VolumeNumber field.
func (o *StatsAppliancesListAllOfData) SetVolumeNumber(v float32) {
	o.VolumeNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StatsAppliancesListAllOfData) SetStatus(v string) {
	o.Status = &v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetFunction() string {
	if o == nil || o.Function == nil {
		var ret string
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetFunctionOk() (*string, bool) {
	if o == nil || o.Function == nil {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasFunction() bool {
	if o != nil && o.Function != nil {
		return true
	}

	return false
}

// SetFunction gets a reference to the given string and assigns it to the Function field.
func (o *StatsAppliancesListAllOfData) SetFunction(v string) {
	o.Function = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetController() ControllerRole {
	if o == nil || o.Controller == nil {
		var ret ControllerRole
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetControllerOk() (*ControllerRole, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given ControllerRole and assigns it to the Controller field.
func (o *StatsAppliancesListAllOfData) SetController(v ControllerRole) {
	o.Controller = &v
}

// GetLogServer returns the LogServer field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetLogServer() ApplianceRole {
	if o == nil || o.LogServer == nil {
		var ret ApplianceRole
		return ret
	}
	return *o.LogServer
}

// GetLogServerOk returns a tuple with the LogServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetLogServerOk() (*ApplianceRole, bool) {
	if o == nil || o.LogServer == nil {
		return nil, false
	}
	return o.LogServer, true
}

// HasLogServer returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasLogServer() bool {
	if o != nil && o.LogServer != nil {
		return true
	}

	return false
}

// SetLogServer gets a reference to the given ApplianceRole and assigns it to the LogServer field.
func (o *StatsAppliancesListAllOfData) SetLogServer(v ApplianceRole) {
	o.LogServer = &v
}

// GetLogForwarder returns the LogForwarder field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetLogForwarder() ApplianceRole {
	if o == nil || o.LogForwarder == nil {
		var ret ApplianceRole
		return ret
	}
	return *o.LogForwarder
}

// GetLogForwarderOk returns a tuple with the LogForwarder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetLogForwarderOk() (*ApplianceRole, bool) {
	if o == nil || o.LogForwarder == nil {
		return nil, false
	}
	return o.LogForwarder, true
}

// HasLogForwarder returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasLogForwarder() bool {
	if o != nil && o.LogForwarder != nil {
		return true
	}

	return false
}

// SetLogForwarder gets a reference to the given ApplianceRole and assigns it to the LogForwarder field.
func (o *StatsAppliancesListAllOfData) SetLogForwarder(v ApplianceRole) {
	o.LogForwarder = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetGateway() ApplianceWithSessionsRole {
	if o == nil || o.Gateway == nil {
		var ret ApplianceWithSessionsRole
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetGatewayOk() (*ApplianceWithSessionsRole, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given ApplianceWithSessionsRole and assigns it to the Gateway field.
func (o *StatsAppliancesListAllOfData) SetGateway(v ApplianceWithSessionsRole) {
	o.Gateway = &v
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetConnector() ApplianceWithSessionsRole {
	if o == nil || o.Connector == nil {
		var ret ApplianceWithSessionsRole
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetConnectorOk() (*ApplianceWithSessionsRole, bool) {
	if o == nil || o.Connector == nil {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasConnector() bool {
	if o != nil && o.Connector != nil {
		return true
	}

	return false
}

// SetConnector gets a reference to the given ApplianceWithSessionsRole and assigns it to the Connector field.
func (o *StatsAppliancesListAllOfData) SetConnector(v ApplianceWithSessionsRole) {
	o.Connector = &v
}

// GetPortal returns the Portal field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetPortal() ApplianceWithSessionsRole {
	if o == nil || o.Portal == nil {
		var ret ApplianceWithSessionsRole
		return ret
	}
	return *o.Portal
}

// GetPortalOk returns a tuple with the Portal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetPortalOk() (*ApplianceWithSessionsRole, bool) {
	if o == nil || o.Portal == nil {
		return nil, false
	}
	return o.Portal, true
}

// HasPortal returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasPortal() bool {
	if o != nil && o.Portal != nil {
		return true
	}

	return false
}

// SetPortal gets a reference to the given ApplianceWithSessionsRole and assigns it to the Portal field.
func (o *StatsAppliancesListAllOfData) SetPortal(v ApplianceWithSessionsRole) {
	o.Portal = &v
}

// GetAppliance returns the Appliance field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetAppliance() ApplianceBaseRole {
	if o == nil || o.Appliance == nil {
		var ret ApplianceBaseRole
		return ret
	}
	return *o.Appliance
}

// GetApplianceOk returns a tuple with the Appliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetApplianceOk() (*ApplianceBaseRole, bool) {
	if o == nil || o.Appliance == nil {
		return nil, false
	}
	return o.Appliance, true
}

// HasAppliance returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasAppliance() bool {
	if o != nil && o.Appliance != nil {
		return true
	}

	return false
}

// SetAppliance gets a reference to the given ApplianceBaseRole and assigns it to the Appliance field.
func (o *StatsAppliancesListAllOfData) SetAppliance(v ApplianceBaseRole) {
	o.Appliance = &v
}

// GetNumberOfSessions returns the NumberOfSessions field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetNumberOfSessions() int32 {
	if o == nil || o.NumberOfSessions == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfSessions
}

// GetNumberOfSessionsOk returns a tuple with the NumberOfSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetNumberOfSessionsOk() (*int32, bool) {
	if o == nil || o.NumberOfSessions == nil {
		return nil, false
	}
	return o.NumberOfSessions, true
}

// HasNumberOfSessions returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasNumberOfSessions() bool {
	if o != nil && o.NumberOfSessions != nil {
		return true
	}

	return false
}

// SetNumberOfSessions gets a reference to the given int32 and assigns it to the NumberOfSessions field.
func (o *StatsAppliancesListAllOfData) SetNumberOfSessions(v int32) {
	o.NumberOfSessions = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetCpu() float32 {
	if o == nil || o.Cpu == nil {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetCpuOk() (*float32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *StatsAppliancesListAllOfData) SetCpu(v float32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetMemory() float32 {
	if o == nil || o.Memory == nil {
		var ret float32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetMemoryOk() (*float32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given float32 and assigns it to the Memory field.
func (o *StatsAppliancesListAllOfData) SetMemory(v float32) {
	o.Memory = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetDisk() float32 {
	if o == nil || o.Disk == nil {
		var ret float32
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetDiskOk() (*float32, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given float32 and assigns it to the Disk field.
func (o *StatsAppliancesListAllOfData) SetDisk(v float32) {
	o.Disk = &v
}

// GetDiskInfo returns the DiskInfo field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetDiskInfo() StatsAppliancesListAllOfDiskInfo {
	if o == nil || o.DiskInfo == nil {
		var ret StatsAppliancesListAllOfDiskInfo
		return ret
	}
	return *o.DiskInfo
}

// GetDiskInfoOk returns a tuple with the DiskInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetDiskInfoOk() (*StatsAppliancesListAllOfDiskInfo, bool) {
	if o == nil || o.DiskInfo == nil {
		return nil, false
	}
	return o.DiskInfo, true
}

// HasDiskInfo returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasDiskInfo() bool {
	if o != nil && o.DiskInfo != nil {
		return true
	}

	return false
}

// SetDiskInfo gets a reference to the given StatsAppliancesListAllOfDiskInfo and assigns it to the DiskInfo field.
func (o *StatsAppliancesListAllOfData) SetDiskInfo(v StatsAppliancesListAllOfDiskInfo) {
	o.DiskInfo = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetNetwork() StatsAppliancesListAllOfNetwork {
	if o == nil || o.Network == nil {
		var ret StatsAppliancesListAllOfNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetNetworkOk() (*StatsAppliancesListAllOfNetwork, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given StatsAppliancesListAllOfNetwork and assigns it to the Network field.
func (o *StatsAppliancesListAllOfData) SetNetwork(v StatsAppliancesListAllOfNetwork) {
	o.Network = &v
}

// GetUpgrade returns the Upgrade field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetUpgrade() StatsAppliancesListAllOfUpgrade {
	if o == nil || o.Upgrade == nil {
		var ret StatsAppliancesListAllOfUpgrade
		return ret
	}
	return *o.Upgrade
}

// GetUpgradeOk returns a tuple with the Upgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetUpgradeOk() (*StatsAppliancesListAllOfUpgrade, bool) {
	if o == nil || o.Upgrade == nil {
		return nil, false
	}
	return o.Upgrade, true
}

// HasUpgrade returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasUpgrade() bool {
	if o != nil && o.Upgrade != nil {
		return true
	}

	return false
}

// SetUpgrade gets a reference to the given StatsAppliancesListAllOfUpgrade and assigns it to the Upgrade field.
func (o *StatsAppliancesListAllOfData) SetUpgrade(v StatsAppliancesListAllOfUpgrade) {
	o.Upgrade = &v
}

// GetCustomizationName returns the CustomizationName field value if set, zero value otherwise.
func (o *StatsAppliancesListAllOfData) GetCustomizationName() string {
	if o == nil || o.CustomizationName == nil {
		var ret string
		return ret
	}
	return *o.CustomizationName
}

// GetCustomizationNameOk returns a tuple with the CustomizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsAppliancesListAllOfData) GetCustomizationNameOk() (*string, bool) {
	if o == nil || o.CustomizationName == nil {
		return nil, false
	}
	return o.CustomizationName, true
}

// HasCustomizationName returns a boolean if a field has been set.
func (o *StatsAppliancesListAllOfData) HasCustomizationName() bool {
	if o != nil && o.CustomizationName != nil {
		return true
	}

	return false
}

// SetCustomizationName gets a reference to the given string and assigns it to the CustomizationName field.
func (o *StatsAppliancesListAllOfData) SetCustomizationName(v string) {
	o.CustomizationName = &v
}

func (o StatsAppliancesListAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SiteName != nil {
		toSerialize["siteName"] = o.SiteName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Online != nil {
		toSerialize["online"] = o.Online
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.VolumeNumber != nil {
		toSerialize["volumeNumber"] = o.VolumeNumber
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Function != nil {
		toSerialize["function"] = o.Function
	}
	if o.Controller != nil {
		toSerialize["controller"] = o.Controller
	}
	if o.LogServer != nil {
		toSerialize["logServer"] = o.LogServer
	}
	if o.LogForwarder != nil {
		toSerialize["logForwarder"] = o.LogForwarder
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.Connector != nil {
		toSerialize["connector"] = o.Connector
	}
	if o.Portal != nil {
		toSerialize["portal"] = o.Portal
	}
	if o.Appliance != nil {
		toSerialize["appliance"] = o.Appliance
	}
	if o.NumberOfSessions != nil {
		toSerialize["numberOfSessions"] = o.NumberOfSessions
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Disk != nil {
		toSerialize["disk"] = o.Disk
	}
	if o.DiskInfo != nil {
		toSerialize["diskInfo"] = o.DiskInfo
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Upgrade != nil {
		toSerialize["upgrade"] = o.Upgrade
	}
	if o.CustomizationName != nil {
		toSerialize["customizationName"] = o.CustomizationName
	}
	return json.Marshal(toSerialize)
}

type NullableStatsAppliancesListAllOfData struct {
	value *StatsAppliancesListAllOfData
	isSet bool
}

func (v NullableStatsAppliancesListAllOfData) Get() *StatsAppliancesListAllOfData {
	return v.value
}

func (v *NullableStatsAppliancesListAllOfData) Set(val *StatsAppliancesListAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsAppliancesListAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsAppliancesListAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsAppliancesListAllOfData(val *StatsAppliancesListAllOfData) *NullableStatsAppliancesListAllOfData {
	return &NullableStatsAppliancesListAllOfData{value: val, isSet: true}
}

func (v NullableStatsAppliancesListAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsAppliancesListAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
