/*
Appgate SDP Controller REST API

# About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the REST API chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin/API TLS Connection (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliance-function-configure.html?anchor=admin-api)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v16+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.

API version: API version 16.2
Contact: appgatesdp.support@appgate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ActiveDevicesApiService ActiveDevicesApi service
type ActiveDevicesApiService service

type ApiTokenRecordsDnGetRequest struct {
	ctx           _context.Context
	ApiService    *ActiveDevicesApiService
	authorization *string
	query         *string
	range_        *string
	orderBy       *string
	descending    *string
	filterBy      *map[string]string
}

// The Token from the LoginResponse.
func (r ApiTokenRecordsDnGetRequest) Authorization(authorization string) ApiTokenRecordsDnGetRequest {
	r.authorization = &authorization
	return r
}

// Query string to filter the result list. It&#39;s used for various fields depending on the object type.
func (r ApiTokenRecordsDnGetRequest) Query(query string) ApiTokenRecordsDnGetRequest {
	r.query = &query
	return r
}

// &#39;Range string to limit the result list. Format: -. 3-10 means he items between the (including) 3rd and the 10th  will be returned. Defaults to all objects.&#39;
func (r ApiTokenRecordsDnGetRequest) Range_(range_ string) ApiTokenRecordsDnGetRequest {
	r.range_ = &range_
	return r
}

// The field name to sort the result list. Supported fields vary from object to object. Defaults to certain field depending on the object type.
func (r ApiTokenRecordsDnGetRequest) OrderBy(orderBy string) ApiTokenRecordsDnGetRequest {
	r.orderBy = &orderBy
	return r
}

// Whether the sorting is applied descending or ascending. Defaults to certain field depending on the object type.
func (r ApiTokenRecordsDnGetRequest) Descending(descending string) ApiTokenRecordsDnGetRequest {
	r.descending = &descending
	return r
}

// Filters the result list by the given field and value. Supported fields vary from object to object. The filters can be combined with each other as well as the generic query field. The given value is checked for inclusion. The representation of the dynamic query parameters is not correct at the moment. See the example for getting a better idea.
func (r ApiTokenRecordsDnGetRequest) FilterBy(filterBy map[string]string) ApiTokenRecordsDnGetRequest {
	r.filterBy = &filterBy
	return r
}

func (r ApiTokenRecordsDnGetRequest) Execute() (DistinguishedNameList, *_nethttp.Response, error) {
	return r.ApiService.TokenRecordsDnGetExecute(r)
}

/*
TokenRecordsDnGet List all Distinguished Names active in the past 24 hour.

List all Distinguished Names active in the past 24 hour. Includes the users who has at least one token that has not expired past 24 hours. If a token was created 30 hours ago and it has 10 hours expiration time, it will be in this list.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTokenRecordsDnGetRequest
*/
func (a *ActiveDevicesApiService) TokenRecordsDnGet(ctx _context.Context) ApiTokenRecordsDnGetRequest {
	return ApiTokenRecordsDnGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return DistinguishedNameList
func (a *ActiveDevicesApiService) TokenRecordsDnGetExecute(r ApiTokenRecordsDnGetRequest) (DistinguishedNameList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DistinguishedNameList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActiveDevicesApiService.TokenRecordsDnGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/token-records/dn"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTokenRecordsReevalByDnDistinguishedNamePostRequest struct {
	ctx               _context.Context
	ApiService        *ActiveDevicesApiService
	authorization     *string
	distinguishedName string
}

// The Token from the LoginResponse.
func (r ApiTokenRecordsReevalByDnDistinguishedNamePostRequest) Authorization(authorization string) ApiTokenRecordsReevalByDnDistinguishedNamePostRequest {
	r.authorization = &authorization
	return r
}

func (r ApiTokenRecordsReevalByDnDistinguishedNamePostRequest) Execute() (InlineResponse20016, *_nethttp.Response, error) {
	return r.ApiService.TokenRecordsReevalByDnDistinguishedNamePostExecute(r)
}

/*
TokenRecordsReevalByDnDistinguishedNamePost Reevaluate all sessions with given Distinguished Name substring.

Reevaluate all sessions belongs to the user&devices ending with the given Distinguished Name substring.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distinguishedName Distinguished name of the user&devices which will be affected by the operation. Format: 'CN=\\<device ID\\>,CN=\\<username\\>,OU=\\<provider name\\>'
 @return ApiTokenRecordsReevalByDnDistinguishedNamePostRequest
*/
func (a *ActiveDevicesApiService) TokenRecordsReevalByDnDistinguishedNamePost(ctx _context.Context, distinguishedName string) ApiTokenRecordsReevalByDnDistinguishedNamePostRequest {
	return ApiTokenRecordsReevalByDnDistinguishedNamePostRequest{
		ApiService:        a,
		ctx:               ctx,
		distinguishedName: distinguishedName,
	}
}

// Execute executes the request
//  @return InlineResponse20016
func (a *ActiveDevicesApiService) TokenRecordsReevalByDnDistinguishedNamePostExecute(r ApiTokenRecordsReevalByDnDistinguishedNamePostRequest) (InlineResponse20016, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20016
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActiveDevicesApiService.TokenRecordsReevalByDnDistinguishedNamePost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/token-records/reeval/by-dn/{distinguished-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"distinguished-name"+"}", _neturl.PathEscape(parameterToString(r.distinguishedName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest struct {
	ctx                    _context.Context
	ApiService             *ActiveDevicesApiService
	authorization          *string
	distinguishedName      string
	siteId                 *string
	tokenType              *string
	tokenRevocationRequest *TokenRevocationRequest
}

// The Token from the LoginResponse.
func (r ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest) Authorization(authorization string) ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest {
	r.authorization = &authorization
	return r
}

// Optional query parameter to revoke only entitlement tokens for the given Site ID.
func (r ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest) SiteId(siteId string) ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest {
	r.siteId = &siteId
	return r
}

// Optional query parameter to revoke only certain types of tokens.
func (r ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest) TokenType(tokenType string) ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest {
	r.tokenType = &tokenType
	return r
}

// Token revocation details.
func (r ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest) TokenRevocationRequest(tokenRevocationRequest TokenRevocationRequest) ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest {
	r.tokenRevocationRequest = &tokenRevocationRequest
	return r
}

func (r ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest) Execute() (TokenRevocationResponse, *_nethttp.Response, error) {
	return r.ApiService.TokenRecordsRevokedByDnDistinguishedNamePutExecute(r)
}

/*
TokenRecordsRevokedByDnDistinguishedNamePut Revoke all Tokens ending with the given Distinguished Name substring.

Revoke all Tokens belong to the user&devices ending with the given Distinguished Name substring.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distinguishedName Distinguished name of the user&devices which will be affected by the operation. Format: 'CN=\\<device ID\\>,CN=\\<username\\>,OU=\\<provider name\\>'
 @return ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest
*/
func (a *ActiveDevicesApiService) TokenRecordsRevokedByDnDistinguishedNamePut(ctx _context.Context, distinguishedName string) ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest {
	return ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest{
		ApiService:        a,
		ctx:               ctx,
		distinguishedName: distinguishedName,
	}
}

// Execute executes the request
//  @return TokenRevocationResponse
func (a *ActiveDevicesApiService) TokenRecordsRevokedByDnDistinguishedNamePutExecute(r ApiTokenRecordsRevokedByDnDistinguishedNamePutRequest) (TokenRevocationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TokenRevocationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActiveDevicesApiService.TokenRecordsRevokedByDnDistinguishedNamePut")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/token-records/revoked/by-dn/{distinguished-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"distinguished-name"+"}", _neturl.PathEscape(parameterToString(r.distinguishedName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}

	if r.siteId != nil {
		localVarQueryParams.Add("siteId", parameterToString(*r.siteId, ""))
	}
	if r.tokenType != nil {
		localVarQueryParams.Add("tokenType", parameterToString(*r.tokenType, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.tokenRevocationRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTokenRecordsRevokedByTypeTokenTypePutRequest struct {
	ctx                    _context.Context
	ApiService             *ActiveDevicesApiService
	authorization          *string
	tokenType              string
	siteId                 *string
	tokenRevocationRequest *TokenRevocationRequest
}

// The Token from the LoginResponse.
func (r ApiTokenRecordsRevokedByTypeTokenTypePutRequest) Authorization(authorization string) ApiTokenRecordsRevokedByTypeTokenTypePutRequest {
	r.authorization = &authorization
	return r
}

// Optional query parameter to revoke only entitlement tokens for the given Site ID.
func (r ApiTokenRecordsRevokedByTypeTokenTypePutRequest) SiteId(siteId string) ApiTokenRecordsRevokedByTypeTokenTypePutRequest {
	r.siteId = &siteId
	return r
}

// Token revocation details.
func (r ApiTokenRecordsRevokedByTypeTokenTypePutRequest) TokenRevocationRequest(tokenRevocationRequest TokenRevocationRequest) ApiTokenRecordsRevokedByTypeTokenTypePutRequest {
	r.tokenRevocationRequest = &tokenRevocationRequest
	return r
}

func (r ApiTokenRecordsRevokedByTypeTokenTypePutRequest) Execute() (TokenRevocationResponse, *_nethttp.Response, error) {
	return r.ApiService.TokenRecordsRevokedByTypeTokenTypePutExecute(r)
}

/*
TokenRecordsRevokedByTypeTokenTypePut Revoke all Tokens with given type.

Revoke all Tokens with given type.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tokenType The type of the tokens.
 @return ApiTokenRecordsRevokedByTypeTokenTypePutRequest
*/
func (a *ActiveDevicesApiService) TokenRecordsRevokedByTypeTokenTypePut(ctx _context.Context, tokenType string) ApiTokenRecordsRevokedByTypeTokenTypePutRequest {
	return ApiTokenRecordsRevokedByTypeTokenTypePutRequest{
		ApiService: a,
		ctx:        ctx,
		tokenType:  tokenType,
	}
}

// Execute executes the request
//  @return TokenRevocationResponse
func (a *ActiveDevicesApiService) TokenRecordsRevokedByTypeTokenTypePutExecute(r ApiTokenRecordsRevokedByTypeTokenTypePutRequest) (TokenRevocationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TokenRevocationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActiveDevicesApiService.TokenRecordsRevokedByTypeTokenTypePut")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/token-records/revoked/by-type/{token-type}"
	localVarPath = strings.Replace(localVarPath, "{"+"token-type"+"}", _neturl.PathEscape(parameterToString(r.tokenType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}

	if r.siteId != nil {
		localVarQueryParams.Add("siteId", parameterToString(*r.siteId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.tokenRevocationRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
