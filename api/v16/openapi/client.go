/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.5
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// APIClient manages communication with the Appgate SDP Controller REST API API vAPI version 16.5
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	ActiveDevicesApi *ActiveDevicesApiService

	ActiveSessionsApi *ActiveSessionsApiService

	AdminMessagesApi *AdminMessagesApiService

	AdministrativeRolesApi *AdministrativeRolesApiService

	ApplianceApi *ApplianceApiService

	ApplianceBackupApi *ApplianceBackupApiService

	ApplianceChangeApi *ApplianceChangeApiService

	ApplianceCustomizationsApi *ApplianceCustomizationsApiService

	ApplianceMaintenanceApi *ApplianceMaintenanceApiService

	ApplianceMetricsApi *ApplianceMetricsApiService

	ApplianceStatsApi *ApplianceStatsApiService

	ApplianceUpgradeApi *ApplianceUpgradeApiService

	AppliancesApi *AppliancesApiService

	BlacklistedUsersApi *BlacklistedUsersApiService

	CAApi *CAApiService

	ClientAutoUpdateApi *ClientAutoUpdateApiService

	ClientConnectionsApi *ClientConnectionsApiService

	ConditionsApi *ConditionsApiService

	CriteriaScriptsApi *CriteriaScriptsApiService

	DNSClassificationsApi *DNSClassificationsApiService

	DNSRulesApi *DNSRulesApiService

	DefaultTimeBasedOTPProviderSeedsApi *DefaultTimeBasedOTPProviderSeedsApiService

	DeviceClaimScriptsApi *DeviceClaimScriptsApiService

	DevicesOnBoardedPerHourApi *DevicesOnBoardedPerHourApiService

	DiscoveredAppsApi *DiscoveredAppsApiService

	EntitlementScriptsApi *EntitlementScriptsApiService

	EntitlementsApi *EntitlementsApiService

	FIDO2DevicesApi *FIDO2DevicesApiService

	FailedAuthenticationsPerHourApi *FailedAuthenticationsPerHourApiService

	GlobalSettingsApi *GlobalSettingsApiService

	IPPoolsApi *IPPoolsApiService

	IdentityProvidersApi *IdentityProvidersApiService

	LicenseApi *LicenseApiService

	LicensedUsersApi *LicensedUsersApiService

	LocalUsersApi *LocalUsersApiService

	LoginApi *LoginApiService

	MFAForAdminsApi *MFAForAdminsApiService

	MFAProvidersApi *MFAProvidersApiService

	OnBoardedDevicesApi *OnBoardedDevicesApiService

	PoliciesApi *PoliciesApiService

	RingfenceRulesApi *RingfenceRulesApiService

	SitesApi *SitesApiService

	TopEntitlementsApi *TopEntitlementsApiService

	TrustedCertificatesApi *TrustedCertificatesApiService

	UserClaimScriptsApi *UserClaimScriptsApiService

	UserLoginsPerHourApi *UserLoginsPerHourApiService
	// PATCH seperate API services for each discriminator identity provider
	LdapIdentityProvidersApi *LdapIdentityProvidersApiService

	RadiusIdentityProvidersApi *RadiusIdentityProvidersApiService

	SamlIdentityProvidersApi *SamlIdentityProvidersApiService

	LocalDatabaseIdentityProvidersApi *LocalDatabaseIdentityProvidersApiService

	LdapCertificateIdentityProvidersApi *LdapCertificateIdentityProvidersApiService

	ConnectorIdentityProvidersApi *ConnectorIdentityProvidersApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.ActiveDevicesApi = (*ActiveDevicesApiService)(&c.common)
	c.ActiveSessionsApi = (*ActiveSessionsApiService)(&c.common)
	c.AdminMessagesApi = (*AdminMessagesApiService)(&c.common)
	c.AdministrativeRolesApi = (*AdministrativeRolesApiService)(&c.common)
	c.ApplianceApi = (*ApplianceApiService)(&c.common)
	c.ApplianceBackupApi = (*ApplianceBackupApiService)(&c.common)
	c.ApplianceChangeApi = (*ApplianceChangeApiService)(&c.common)
	c.ApplianceCustomizationsApi = (*ApplianceCustomizationsApiService)(&c.common)
	c.ApplianceMaintenanceApi = (*ApplianceMaintenanceApiService)(&c.common)
	c.ApplianceMetricsApi = (*ApplianceMetricsApiService)(&c.common)
	c.ApplianceStatsApi = (*ApplianceStatsApiService)(&c.common)
	c.ApplianceUpgradeApi = (*ApplianceUpgradeApiService)(&c.common)
	c.AppliancesApi = (*AppliancesApiService)(&c.common)
	c.BlacklistedUsersApi = (*BlacklistedUsersApiService)(&c.common)
	c.CAApi = (*CAApiService)(&c.common)
	c.ClientAutoUpdateApi = (*ClientAutoUpdateApiService)(&c.common)
	c.ClientConnectionsApi = (*ClientConnectionsApiService)(&c.common)
	c.ConditionsApi = (*ConditionsApiService)(&c.common)
	c.CriteriaScriptsApi = (*CriteriaScriptsApiService)(&c.common)
	c.DNSClassificationsApi = (*DNSClassificationsApiService)(&c.common)
	c.DNSRulesApi = (*DNSRulesApiService)(&c.common)
	c.DefaultTimeBasedOTPProviderSeedsApi = (*DefaultTimeBasedOTPProviderSeedsApiService)(&c.common)
	c.DeviceClaimScriptsApi = (*DeviceClaimScriptsApiService)(&c.common)
	c.DevicesOnBoardedPerHourApi = (*DevicesOnBoardedPerHourApiService)(&c.common)
	c.DiscoveredAppsApi = (*DiscoveredAppsApiService)(&c.common)
	c.EntitlementScriptsApi = (*EntitlementScriptsApiService)(&c.common)
	c.EntitlementsApi = (*EntitlementsApiService)(&c.common)
	c.FIDO2DevicesApi = (*FIDO2DevicesApiService)(&c.common)
	c.FailedAuthenticationsPerHourApi = (*FailedAuthenticationsPerHourApiService)(&c.common)
	c.GlobalSettingsApi = (*GlobalSettingsApiService)(&c.common)
	c.IPPoolsApi = (*IPPoolsApiService)(&c.common)
	c.IdentityProvidersApi = (*IdentityProvidersApiService)(&c.common)
	c.LicenseApi = (*LicenseApiService)(&c.common)
	c.LicensedUsersApi = (*LicensedUsersApiService)(&c.common)
	c.LocalUsersApi = (*LocalUsersApiService)(&c.common)
	c.LoginApi = (*LoginApiService)(&c.common)
	c.MFAForAdminsApi = (*MFAForAdminsApiService)(&c.common)
	c.MFAProvidersApi = (*MFAProvidersApiService)(&c.common)
	c.OnBoardedDevicesApi = (*OnBoardedDevicesApiService)(&c.common)
	c.PoliciesApi = (*PoliciesApiService)(&c.common)
	c.RingfenceRulesApi = (*RingfenceRulesApiService)(&c.common)
	c.SitesApi = (*SitesApiService)(&c.common)
	c.TopEntitlementsApi = (*TopEntitlementsApiService)(&c.common)
	c.TrustedCertificatesApi = (*TrustedCertificatesApiService)(&c.common)
	c.UserClaimScriptsApi = (*UserClaimScriptsApiService)(&c.common)
	c.UserLoginsPerHourApi = (*UserLoginsPerHourApiService)(&c.common)
	// PATCH manually added to replace IdentityProvidersApiService
	// since openapi.generator does not play well with discriminator from the open api spec.
	c.LdapIdentityProvidersApi = (*LdapIdentityProvidersApiService)(&c.common)
	c.RadiusIdentityProvidersApi = (*RadiusIdentityProvidersApiService)(&c.common)
	c.SamlIdentityProvidersApi = (*SamlIdentityProvidersApiService)(&c.common)
	c.LocalDatabaseIdentityProvidersApi = (*LocalDatabaseIdentityProvidersApiService)(&c.common)
	c.LdapCertificateIdentityProvidersApi = (*LdapCertificateIdentityProvidersApiService)(&c.common)
	c.ConnectorIdentityProvidersApi = (*ConnectorIdentityProvidersApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFileName string,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile(formFileName, filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)
	for header, value := range c.cfg.DefaultHeader {
		// For the sake of convenience how appgate SDP Accept header works,
		// We will replace the default apppend (Add()) with Set()
		// so that we can configure the client as:
		//
		// clientCfg := &openapi.Configuration{
		//     DefaultHeader: map[string]string{
		//         "Accept": fmt.Sprintf("application/vnd.appgate.peer-v%d+json", version),
		//     },
		//     Debug: true,
		//     Servers: []openapi.ServerConfiguration{
		//         {
		//             URL: address,
		//         },
		//     },
		//     HTTPClient: httpclient,
		// }
		// Since there is already a Accept header present, we want to replace it, and we can not do that with Add()
		localVarRequest.Header.Set(header, value)
	}
	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
		// Accept Header, Overwrites any existing Accept headers
		if accept, ok := ctx.Value(ContextAcceptHeader).(string); ok {
			localVarRequest.Header.Set("Accept", accept)
		}

	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(**os.File); ok {
		*f, err = ioutil.TempFile("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	// https://github.com/OpenAPITools/openapi-generator/issues/11965
	if file, ok := v.(***os.File); ok {
		var tmp *os.File
		tmp, err = ioutil.TempFile("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = tmp.Write(b)
		if err != nil {
			return
		}
		_, err = tmp.Seek(0, io.SeekStart)
		*file = &tmp
		return
	}
	if xmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if jsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(**os.File); ok {
		_, err = bodyBuf.ReadFrom(*fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}
