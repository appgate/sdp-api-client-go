/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v21+json**    An exception is made for the **_/admin/version** endpoint which instead expects an **application/json** Accept header. # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommended if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 21.1
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ApplianceStatsDeprecatedApiService ApplianceStatsDeprecatedApi service
type ApplianceStatsDeprecatedApiService service

type ApiStatsAppliancesGetRequest struct {
	ctx           context.Context
	ApiService    *ApplianceStatsDeprecatedApiService
	authorization *string
	query         *string
	range_        *string
	orderBy       *string
	descending    *string
	filterBy      *map[string]string
}

// The Token from the LoginResponse.
func (r ApiStatsAppliancesGetRequest) Authorization(authorization string) ApiStatsAppliancesGetRequest {
	r.authorization = &authorization
	return r
}

// Query string to filter the result list. It&#39;s used for various fields depending on the object type.  Send multiple query parameters to make the queries more specific.
func (r ApiStatsAppliancesGetRequest) Query(query string) ApiStatsAppliancesGetRequest {
	r.query = &query
	return r
}

// &#39;Range string to limit the result list. Format: -. 3-10 means he items between the (including) 3rd and the 10th  will be returned. Defaults to all objects.&#39;
func (r ApiStatsAppliancesGetRequest) Range_(range_ string) ApiStatsAppliancesGetRequest {
	r.range_ = &range_
	return r
}

// The field name to sort the result list. Supported fields vary from object to object. Defaults to certain field depending on the object type.
func (r ApiStatsAppliancesGetRequest) OrderBy(orderBy string) ApiStatsAppliancesGetRequest {
	r.orderBy = &orderBy
	return r
}

// Whether the sorting is applied descending or ascending. Defaults to certain field depending on the object type.
func (r ApiStatsAppliancesGetRequest) Descending(descending string) ApiStatsAppliancesGetRequest {
	r.descending = &descending
	return r
}

// Filters the result list by the given field and value. Supported fields vary from API to API. The filters can be combined with each other as well as the generic query parameter. The given value is checked for inclusion.
func (r ApiStatsAppliancesGetRequest) FilterBy(filterBy map[string]string) ApiStatsAppliancesGetRequest {
	r.filterBy = &filterBy
	return r
}

func (r ApiStatsAppliancesGetRequest) Execute() (*StatsAppliancesList, *http.Response, error) {
	return r.ApiService.StatsAppliancesGetExecute(r)
}

/*
StatsAppliancesGet Get Appliance Stats.

Get Stats and status of the active appliances. This API makes the controller to query every active appliance for status. The operation may take long if one or more appliances take long to respond.

**Note:** if identical filter keys are chained with '&', the first filter is applied and the rest are ignored.

Deprecated as of 6.3. Use `/appliances/status` instead.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStatsAppliancesGetRequest

Deprecated
*/
func (a *ApplianceStatsDeprecatedApiService) StatsAppliancesGet(ctx context.Context) ApiStatsAppliancesGetRequest {
	return ApiStatsAppliancesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StatsAppliancesList
//
// Deprecated
func (a *ApplianceStatsDeprecatedApiService) StatsAppliancesGetExecute(r ApiStatsAppliancesGetRequest) (*StatsAppliancesList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StatsAppliancesList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplianceStatsDeprecatedApiService.StatsAppliancesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stats/appliances"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}

	if r.query != nil {
		localVarQueryParams.Add("query", parameterToString(*r.query, ""))
	}
	if r.range_ != nil {
		localVarQueryParams.Add("range", parameterToString(*r.range_, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.descending != nil {
		localVarQueryParams.Add("descending", parameterToString(*r.descending, ""))
	}
	if r.filterBy != nil {
		localVarQueryParams.Add("filterBy", parameterToString(*r.filterBy, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 406 {
			var v LoginPost406Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
