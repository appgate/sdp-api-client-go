/*
 * Appgate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the Integration chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-functions-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v15+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 15
 * Contact: appgatesdp.support@appgate.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplianceAllOf Represents an Appliance.
type ApplianceAllOf struct {
	// Whether the Appliance is activated or not. If it is not activated, it won't be accessible by the Clients.
	Activated *bool `json:"activated,omitempty"`
	// Whether the Appliance is pending certificate renewal or not. Should be true for a very short period on certificate renewal.
	PendingCertificateRenewal *bool `json:"pendingCertificateRenewal,omitempty"`
	// Peer version of the Appliance.
	Version *int32 `json:"version,omitempty"`
	// Generic hostname of the appliance. Used as linux hostname and to identify within logs. If peerInterface.hostname is changed, this field's value is cleared. When empty, peerInterface.hostname will be used to generate it.
	Hostname *string `json:"hostname,omitempty"`
	// Site served by the Appliance. Entitlements on this Site will be included in the Entitlement Token for this Appliance. Not useful if Gateway role is not enabled.
	Site *string `json:"site,omitempty"`
	// Name of the Site for this Appliance. For convenience only.
	SiteName *string `json:"siteName,omitempty"`
	// Customization assigned to this Appliance.
	Customization *string `json:"customization,omitempty"`
	// Makes the Appliance to connect to Controller/LogServer/LogForwarders using their clientInterface.httpsPort instead of peerInterface.httpsPort. The Appliance uses SPA to connect. This field is deprecated as of 5.4. It will always be enabled when the support for peerInterface is removed.
	ConnectToPeersUsingClientPortWithSpa *bool                             `json:"connectToPeersUsingClientPortWithSpa,omitempty"`
	ClientInterface                      ApplianceAllOfClientInterface     `json:"clientInterface"`
	PeerInterface                        ApplianceAllOfPeerInterface       `json:"peerInterface"`
	AdminInterface                       *ApplianceAllOfAdminInterface     `json:"adminInterface,omitempty"`
	Networking                           ApplianceAllOfNetworking          `json:"networking"`
	Ntp                                  *ApplianceAllOfNtp                `json:"ntp,omitempty"`
	SshServer                            *ApplianceAllOfSshServer          `json:"sshServer,omitempty"`
	SnmpServer                           *ApplianceAllOfSnmpServer         `json:"snmpServer,omitempty"`
	HealthcheckServer                    *ApplianceAllOfHealthcheckServer  `json:"healthcheckServer,omitempty"`
	PrometheusExporter                   *ApplianceAllOfPrometheusExporter `json:"prometheusExporter,omitempty"`
	Ping                                 *ApplianceAllOfPing               `json:"ping,omitempty"`
	LogServer                            *ApplianceAllOfLogServer          `json:"logServer,omitempty"`
	Controller                           *ApplianceAllOfController         `json:"controller,omitempty"`
	Gateway                              *ApplianceAllOfGateway            `json:"gateway,omitempty"`
	LogForwarder                         *ApplianceAllOfLogForwarder       `json:"logForwarder,omitempty"`
	Connector                            *ApplianceAllOfConnector          `json:"connector,omitempty"`
	Portal                               *ApplianceAllOfPortal             `json:"portal,omitempty"`
	// Rsyslog destination settings to forward appliance logs.
	RsyslogDestinations *[]ApplianceAllOfRsyslogDestinations `json:"rsyslogDestinations,omitempty"`
	// Hostname aliases. They are added to the Appliance certificate as Subject Alternative Names so it is trusted using different IPs or hostnames. Requires manual certificate renewal to apply changes to the certificate.
	HostnameAliases *[]string `json:"hostnameAliases,omitempty"`
}

// NewApplianceAllOf instantiates a new ApplianceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOf(clientInterface ApplianceAllOfClientInterface, peerInterface ApplianceAllOfPeerInterface, networking ApplianceAllOfNetworking) *ApplianceAllOf {
	this := ApplianceAllOf{}
	var connectToPeersUsingClientPortWithSpa bool = true
	this.ConnectToPeersUsingClientPortWithSpa = &connectToPeersUsingClientPortWithSpa
	this.ClientInterface = clientInterface
	this.PeerInterface = peerInterface
	this.Networking = networking
	return &this
}

// NewApplianceAllOfWithDefaults instantiates a new ApplianceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfWithDefaults() *ApplianceAllOf {
	this := ApplianceAllOf{}
	var connectToPeersUsingClientPortWithSpa bool = true
	this.ConnectToPeersUsingClientPortWithSpa = &connectToPeersUsingClientPortWithSpa
	return &this
}

// GetActivated returns the Activated field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetActivated() bool {
	if o == nil || o.Activated == nil {
		var ret bool
		return ret
	}
	return *o.Activated
}

// GetActivatedOk returns a tuple with the Activated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetActivatedOk() (*bool, bool) {
	if o == nil || o.Activated == nil {
		return nil, false
	}
	return o.Activated, true
}

// HasActivated returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasActivated() bool {
	if o != nil && o.Activated != nil {
		return true
	}

	return false
}

// SetActivated gets a reference to the given bool and assigns it to the Activated field.
func (o *ApplianceAllOf) SetActivated(v bool) {
	o.Activated = &v
}

// GetPendingCertificateRenewal returns the PendingCertificateRenewal field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetPendingCertificateRenewal() bool {
	if o == nil || o.PendingCertificateRenewal == nil {
		var ret bool
		return ret
	}
	return *o.PendingCertificateRenewal
}

// GetPendingCertificateRenewalOk returns a tuple with the PendingCertificateRenewal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetPendingCertificateRenewalOk() (*bool, bool) {
	if o == nil || o.PendingCertificateRenewal == nil {
		return nil, false
	}
	return o.PendingCertificateRenewal, true
}

// HasPendingCertificateRenewal returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasPendingCertificateRenewal() bool {
	if o != nil && o.PendingCertificateRenewal != nil {
		return true
	}

	return false
}

// SetPendingCertificateRenewal gets a reference to the given bool and assigns it to the PendingCertificateRenewal field.
func (o *ApplianceAllOf) SetPendingCertificateRenewal(v bool) {
	o.PendingCertificateRenewal = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ApplianceAllOf) SetVersion(v int32) {
	o.Version = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApplianceAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetSite() string {
	if o == nil || o.Site == nil {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetSiteOk() (*string, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *ApplianceAllOf) SetSite(v string) {
	o.Site = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *ApplianceAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetCustomization returns the Customization field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetCustomization() string {
	if o == nil || o.Customization == nil {
		var ret string
		return ret
	}
	return *o.Customization
}

// GetCustomizationOk returns a tuple with the Customization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetCustomizationOk() (*string, bool) {
	if o == nil || o.Customization == nil {
		return nil, false
	}
	return o.Customization, true
}

// HasCustomization returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasCustomization() bool {
	if o != nil && o.Customization != nil {
		return true
	}

	return false
}

// SetCustomization gets a reference to the given string and assigns it to the Customization field.
func (o *ApplianceAllOf) SetCustomization(v string) {
	o.Customization = &v
}

// GetConnectToPeersUsingClientPortWithSpa returns the ConnectToPeersUsingClientPortWithSpa field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetConnectToPeersUsingClientPortWithSpa() bool {
	if o == nil || o.ConnectToPeersUsingClientPortWithSpa == nil {
		var ret bool
		return ret
	}
	return *o.ConnectToPeersUsingClientPortWithSpa
}

// GetConnectToPeersUsingClientPortWithSpaOk returns a tuple with the ConnectToPeersUsingClientPortWithSpa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetConnectToPeersUsingClientPortWithSpaOk() (*bool, bool) {
	if o == nil || o.ConnectToPeersUsingClientPortWithSpa == nil {
		return nil, false
	}
	return o.ConnectToPeersUsingClientPortWithSpa, true
}

// HasConnectToPeersUsingClientPortWithSpa returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasConnectToPeersUsingClientPortWithSpa() bool {
	if o != nil && o.ConnectToPeersUsingClientPortWithSpa != nil {
		return true
	}

	return false
}

// SetConnectToPeersUsingClientPortWithSpa gets a reference to the given bool and assigns it to the ConnectToPeersUsingClientPortWithSpa field.
func (o *ApplianceAllOf) SetConnectToPeersUsingClientPortWithSpa(v bool) {
	o.ConnectToPeersUsingClientPortWithSpa = &v
}

// GetClientInterface returns the ClientInterface field value
func (o *ApplianceAllOf) GetClientInterface() ApplianceAllOfClientInterface {
	if o == nil {
		var ret ApplianceAllOfClientInterface
		return ret
	}

	return o.ClientInterface
}

// GetClientInterfaceOk returns a tuple with the ClientInterface field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetClientInterfaceOk() (*ApplianceAllOfClientInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientInterface, true
}

// SetClientInterface sets field value
func (o *ApplianceAllOf) SetClientInterface(v ApplianceAllOfClientInterface) {
	o.ClientInterface = v
}

// GetPeerInterface returns the PeerInterface field value
func (o *ApplianceAllOf) GetPeerInterface() ApplianceAllOfPeerInterface {
	if o == nil {
		var ret ApplianceAllOfPeerInterface
		return ret
	}

	return o.PeerInterface
}

// GetPeerInterfaceOk returns a tuple with the PeerInterface field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetPeerInterfaceOk() (*ApplianceAllOfPeerInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeerInterface, true
}

// SetPeerInterface sets field value
func (o *ApplianceAllOf) SetPeerInterface(v ApplianceAllOfPeerInterface) {
	o.PeerInterface = v
}

// GetAdminInterface returns the AdminInterface field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetAdminInterface() ApplianceAllOfAdminInterface {
	if o == nil || o.AdminInterface == nil {
		var ret ApplianceAllOfAdminInterface
		return ret
	}
	return *o.AdminInterface
}

// GetAdminInterfaceOk returns a tuple with the AdminInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetAdminInterfaceOk() (*ApplianceAllOfAdminInterface, bool) {
	if o == nil || o.AdminInterface == nil {
		return nil, false
	}
	return o.AdminInterface, true
}

// HasAdminInterface returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasAdminInterface() bool {
	if o != nil && o.AdminInterface != nil {
		return true
	}

	return false
}

// SetAdminInterface gets a reference to the given ApplianceAllOfAdminInterface and assigns it to the AdminInterface field.
func (o *ApplianceAllOf) SetAdminInterface(v ApplianceAllOfAdminInterface) {
	o.AdminInterface = &v
}

// GetNetworking returns the Networking field value
func (o *ApplianceAllOf) GetNetworking() ApplianceAllOfNetworking {
	if o == nil {
		var ret ApplianceAllOfNetworking
		return ret
	}

	return o.Networking
}

// GetNetworkingOk returns a tuple with the Networking field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetNetworkingOk() (*ApplianceAllOfNetworking, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Networking, true
}

// SetNetworking sets field value
func (o *ApplianceAllOf) SetNetworking(v ApplianceAllOfNetworking) {
	o.Networking = v
}

// GetNtp returns the Ntp field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetNtp() ApplianceAllOfNtp {
	if o == nil || o.Ntp == nil {
		var ret ApplianceAllOfNtp
		return ret
	}
	return *o.Ntp
}

// GetNtpOk returns a tuple with the Ntp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetNtpOk() (*ApplianceAllOfNtp, bool) {
	if o == nil || o.Ntp == nil {
		return nil, false
	}
	return o.Ntp, true
}

// HasNtp returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasNtp() bool {
	if o != nil && o.Ntp != nil {
		return true
	}

	return false
}

// SetNtp gets a reference to the given ApplianceAllOfNtp and assigns it to the Ntp field.
func (o *ApplianceAllOf) SetNtp(v ApplianceAllOfNtp) {
	o.Ntp = &v
}

// GetSshServer returns the SshServer field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetSshServer() ApplianceAllOfSshServer {
	if o == nil || o.SshServer == nil {
		var ret ApplianceAllOfSshServer
		return ret
	}
	return *o.SshServer
}

// GetSshServerOk returns a tuple with the SshServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetSshServerOk() (*ApplianceAllOfSshServer, bool) {
	if o == nil || o.SshServer == nil {
		return nil, false
	}
	return o.SshServer, true
}

// HasSshServer returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasSshServer() bool {
	if o != nil && o.SshServer != nil {
		return true
	}

	return false
}

// SetSshServer gets a reference to the given ApplianceAllOfSshServer and assigns it to the SshServer field.
func (o *ApplianceAllOf) SetSshServer(v ApplianceAllOfSshServer) {
	o.SshServer = &v
}

// GetSnmpServer returns the SnmpServer field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetSnmpServer() ApplianceAllOfSnmpServer {
	if o == nil || o.SnmpServer == nil {
		var ret ApplianceAllOfSnmpServer
		return ret
	}
	return *o.SnmpServer
}

// GetSnmpServerOk returns a tuple with the SnmpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetSnmpServerOk() (*ApplianceAllOfSnmpServer, bool) {
	if o == nil || o.SnmpServer == nil {
		return nil, false
	}
	return o.SnmpServer, true
}

// HasSnmpServer returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasSnmpServer() bool {
	if o != nil && o.SnmpServer != nil {
		return true
	}

	return false
}

// SetSnmpServer gets a reference to the given ApplianceAllOfSnmpServer and assigns it to the SnmpServer field.
func (o *ApplianceAllOf) SetSnmpServer(v ApplianceAllOfSnmpServer) {
	o.SnmpServer = &v
}

// GetHealthcheckServer returns the HealthcheckServer field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetHealthcheckServer() ApplianceAllOfHealthcheckServer {
	if o == nil || o.HealthcheckServer == nil {
		var ret ApplianceAllOfHealthcheckServer
		return ret
	}
	return *o.HealthcheckServer
}

// GetHealthcheckServerOk returns a tuple with the HealthcheckServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetHealthcheckServerOk() (*ApplianceAllOfHealthcheckServer, bool) {
	if o == nil || o.HealthcheckServer == nil {
		return nil, false
	}
	return o.HealthcheckServer, true
}

// HasHealthcheckServer returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasHealthcheckServer() bool {
	if o != nil && o.HealthcheckServer != nil {
		return true
	}

	return false
}

// SetHealthcheckServer gets a reference to the given ApplianceAllOfHealthcheckServer and assigns it to the HealthcheckServer field.
func (o *ApplianceAllOf) SetHealthcheckServer(v ApplianceAllOfHealthcheckServer) {
	o.HealthcheckServer = &v
}

// GetPrometheusExporter returns the PrometheusExporter field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetPrometheusExporter() ApplianceAllOfPrometheusExporter {
	if o == nil || o.PrometheusExporter == nil {
		var ret ApplianceAllOfPrometheusExporter
		return ret
	}
	return *o.PrometheusExporter
}

// GetPrometheusExporterOk returns a tuple with the PrometheusExporter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetPrometheusExporterOk() (*ApplianceAllOfPrometheusExporter, bool) {
	if o == nil || o.PrometheusExporter == nil {
		return nil, false
	}
	return o.PrometheusExporter, true
}

// HasPrometheusExporter returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasPrometheusExporter() bool {
	if o != nil && o.PrometheusExporter != nil {
		return true
	}

	return false
}

// SetPrometheusExporter gets a reference to the given ApplianceAllOfPrometheusExporter and assigns it to the PrometheusExporter field.
func (o *ApplianceAllOf) SetPrometheusExporter(v ApplianceAllOfPrometheusExporter) {
	o.PrometheusExporter = &v
}

// GetPing returns the Ping field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetPing() ApplianceAllOfPing {
	if o == nil || o.Ping == nil {
		var ret ApplianceAllOfPing
		return ret
	}
	return *o.Ping
}

// GetPingOk returns a tuple with the Ping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetPingOk() (*ApplianceAllOfPing, bool) {
	if o == nil || o.Ping == nil {
		return nil, false
	}
	return o.Ping, true
}

// HasPing returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasPing() bool {
	if o != nil && o.Ping != nil {
		return true
	}

	return false
}

// SetPing gets a reference to the given ApplianceAllOfPing and assigns it to the Ping field.
func (o *ApplianceAllOf) SetPing(v ApplianceAllOfPing) {
	o.Ping = &v
}

// GetLogServer returns the LogServer field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetLogServer() ApplianceAllOfLogServer {
	if o == nil || o.LogServer == nil {
		var ret ApplianceAllOfLogServer
		return ret
	}
	return *o.LogServer
}

// GetLogServerOk returns a tuple with the LogServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetLogServerOk() (*ApplianceAllOfLogServer, bool) {
	if o == nil || o.LogServer == nil {
		return nil, false
	}
	return o.LogServer, true
}

// HasLogServer returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasLogServer() bool {
	if o != nil && o.LogServer != nil {
		return true
	}

	return false
}

// SetLogServer gets a reference to the given ApplianceAllOfLogServer and assigns it to the LogServer field.
func (o *ApplianceAllOf) SetLogServer(v ApplianceAllOfLogServer) {
	o.LogServer = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetController() ApplianceAllOfController {
	if o == nil || o.Controller == nil {
		var ret ApplianceAllOfController
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetControllerOk() (*ApplianceAllOfController, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given ApplianceAllOfController and assigns it to the Controller field.
func (o *ApplianceAllOf) SetController(v ApplianceAllOfController) {
	o.Controller = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetGateway() ApplianceAllOfGateway {
	if o == nil || o.Gateway == nil {
		var ret ApplianceAllOfGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetGatewayOk() (*ApplianceAllOfGateway, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given ApplianceAllOfGateway and assigns it to the Gateway field.
func (o *ApplianceAllOf) SetGateway(v ApplianceAllOfGateway) {
	o.Gateway = &v
}

// GetLogForwarder returns the LogForwarder field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetLogForwarder() ApplianceAllOfLogForwarder {
	if o == nil || o.LogForwarder == nil {
		var ret ApplianceAllOfLogForwarder
		return ret
	}
	return *o.LogForwarder
}

// GetLogForwarderOk returns a tuple with the LogForwarder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetLogForwarderOk() (*ApplianceAllOfLogForwarder, bool) {
	if o == nil || o.LogForwarder == nil {
		return nil, false
	}
	return o.LogForwarder, true
}

// HasLogForwarder returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasLogForwarder() bool {
	if o != nil && o.LogForwarder != nil {
		return true
	}

	return false
}

// SetLogForwarder gets a reference to the given ApplianceAllOfLogForwarder and assigns it to the LogForwarder field.
func (o *ApplianceAllOf) SetLogForwarder(v ApplianceAllOfLogForwarder) {
	o.LogForwarder = &v
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetConnector() ApplianceAllOfConnector {
	if o == nil || o.Connector == nil {
		var ret ApplianceAllOfConnector
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetConnectorOk() (*ApplianceAllOfConnector, bool) {
	if o == nil || o.Connector == nil {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasConnector() bool {
	if o != nil && o.Connector != nil {
		return true
	}

	return false
}

// SetConnector gets a reference to the given ApplianceAllOfConnector and assigns it to the Connector field.
func (o *ApplianceAllOf) SetConnector(v ApplianceAllOfConnector) {
	o.Connector = &v
}

// GetPortal returns the Portal field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetPortal() ApplianceAllOfPortal {
	if o == nil || o.Portal == nil {
		var ret ApplianceAllOfPortal
		return ret
	}
	return *o.Portal
}

// GetPortalOk returns a tuple with the Portal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetPortalOk() (*ApplianceAllOfPortal, bool) {
	if o == nil || o.Portal == nil {
		return nil, false
	}
	return o.Portal, true
}

// HasPortal returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasPortal() bool {
	if o != nil && o.Portal != nil {
		return true
	}

	return false
}

// SetPortal gets a reference to the given ApplianceAllOfPortal and assigns it to the Portal field.
func (o *ApplianceAllOf) SetPortal(v ApplianceAllOfPortal) {
	o.Portal = &v
}

// GetRsyslogDestinations returns the RsyslogDestinations field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetRsyslogDestinations() []ApplianceAllOfRsyslogDestinations {
	if o == nil || o.RsyslogDestinations == nil {
		var ret []ApplianceAllOfRsyslogDestinations
		return ret
	}
	return *o.RsyslogDestinations
}

// GetRsyslogDestinationsOk returns a tuple with the RsyslogDestinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetRsyslogDestinationsOk() (*[]ApplianceAllOfRsyslogDestinations, bool) {
	if o == nil || o.RsyslogDestinations == nil {
		return nil, false
	}
	return o.RsyslogDestinations, true
}

// HasRsyslogDestinations returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasRsyslogDestinations() bool {
	if o != nil && o.RsyslogDestinations != nil {
		return true
	}

	return false
}

// SetRsyslogDestinations gets a reference to the given []ApplianceAllOfRsyslogDestinations and assigns it to the RsyslogDestinations field.
func (o *ApplianceAllOf) SetRsyslogDestinations(v []ApplianceAllOfRsyslogDestinations) {
	o.RsyslogDestinations = &v
}

// GetHostnameAliases returns the HostnameAliases field value if set, zero value otherwise.
func (o *ApplianceAllOf) GetHostnameAliases() []string {
	if o == nil || o.HostnameAliases == nil {
		var ret []string
		return ret
	}
	return *o.HostnameAliases
}

// GetHostnameAliasesOk returns a tuple with the HostnameAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOf) GetHostnameAliasesOk() (*[]string, bool) {
	if o == nil || o.HostnameAliases == nil {
		return nil, false
	}
	return o.HostnameAliases, true
}

// HasHostnameAliases returns a boolean if a field has been set.
func (o *ApplianceAllOf) HasHostnameAliases() bool {
	if o != nil && o.HostnameAliases != nil {
		return true
	}

	return false
}

// SetHostnameAliases gets a reference to the given []string and assigns it to the HostnameAliases field.
func (o *ApplianceAllOf) SetHostnameAliases(v []string) {
	o.HostnameAliases = &v
}

func (o ApplianceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Activated != nil {
		toSerialize["activated"] = o.Activated
	}
	if o.PendingCertificateRenewal != nil {
		toSerialize["pendingCertificateRenewal"] = o.PendingCertificateRenewal
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	if o.SiteName != nil {
		toSerialize["siteName"] = o.SiteName
	}
	if o.Customization != nil {
		toSerialize["customization"] = o.Customization
	}
	if o.ConnectToPeersUsingClientPortWithSpa != nil {
		toSerialize["connectToPeersUsingClientPortWithSpa"] = o.ConnectToPeersUsingClientPortWithSpa
	}
	if true {
		toSerialize["clientInterface"] = o.ClientInterface
	}
	if true {
		toSerialize["peerInterface"] = o.PeerInterface
	}
	if o.AdminInterface != nil {
		toSerialize["adminInterface"] = o.AdminInterface
	}
	if true {
		toSerialize["networking"] = o.Networking
	}
	if o.Ntp != nil {
		toSerialize["ntp"] = o.Ntp
	}
	if o.SshServer != nil {
		toSerialize["sshServer"] = o.SshServer
	}
	if o.SnmpServer != nil {
		toSerialize["snmpServer"] = o.SnmpServer
	}
	if o.HealthcheckServer != nil {
		toSerialize["healthcheckServer"] = o.HealthcheckServer
	}
	if o.PrometheusExporter != nil {
		toSerialize["prometheusExporter"] = o.PrometheusExporter
	}
	if o.Ping != nil {
		toSerialize["ping"] = o.Ping
	}
	if o.LogServer != nil {
		toSerialize["logServer"] = o.LogServer
	}
	if o.Controller != nil {
		toSerialize["controller"] = o.Controller
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.LogForwarder != nil {
		toSerialize["logForwarder"] = o.LogForwarder
	}
	if o.Connector != nil {
		toSerialize["connector"] = o.Connector
	}
	if o.Portal != nil {
		toSerialize["portal"] = o.Portal
	}
	if o.RsyslogDestinations != nil {
		toSerialize["rsyslogDestinations"] = o.RsyslogDestinations
	}
	if o.HostnameAliases != nil {
		toSerialize["hostnameAliases"] = o.HostnameAliases
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOf struct {
	value *ApplianceAllOf
	isSet bool
}

func (v NullableApplianceAllOf) Get() *ApplianceAllOf {
	return v.value
}

func (v *NullableApplianceAllOf) Set(val *ApplianceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOf(val *ApplianceAllOf) *NullableApplianceAllOf {
	return &NullableApplianceAllOf{value: val, isSet: true}
}

func (v NullableApplianceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
