package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/appgate/sdp-api-client-go/api/v18/openapi"
)

const (
	// insecure toggle TLS verification
	insecure = true
	// provider represents the default authentication identity provider.
	provider = "local"
)

type Sdp struct {
	token  string
	client *openapi.APIClient
}

func main() {
	var (
		address  string
		username string
		password string
	)
	flag.StringVar(&address, "address", "", "appgate SDP API endpoint, example https://controller.appgate.devops:8443/admin")
	flag.StringVar(&username, "username", "", "username to local identity provider")
	flag.StringVar(&password, "password", "", "password to local identity provider")
	flag.Parse()
	clientCfg := &openapi.Configuration{
		DefaultHeader: map[string]string{
			// Always set expected Accept header
			// See specification:
			// https://github.com/appgate/sdp-api-specification/blob/3f41903e2c53675dcc70c1fe242cd402683969c0/api_specs.yml#L23-L72
			// Set the correct version based on your controller appliance version.
			//
			// The API version is backwards compatible 2 versions back.
			//
			// | Appliance Version 	| API Version 	|
			// |-------------------	|-------------	|
			// | 5.1.*             	| 12          	|
			// | 5.2.*             	| 13          	|
			// | 5.3.*             	| 14          	|
			// | 5.4.*             	| 15          	|
			// | 5.5.*             	| 16          	|
			// | 6.0.*             	| 17          	|
			// | 6.1.*             	| 18          	|
			// | 6.2.*             	| 19          	|

			"Accept": fmt.Sprintf("application/vnd.appgate.peer-v%d+json", 17),
		},
		// Toggle Debug to see Request/Response in stdout
		Debug:   false,
		Servers: []openapi.ServerConfiguration{{URL: address}},
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecure,
				},
			},
		},
	}
	sdp := Sdp{
		client: openapi.NewAPIClient(clientCfg),
	}
	ctx := context.Background()
	// First we need to login to get a bearer token that we can use in the consecutive requests.
	token, err := sdp.login(ctx, username, password)
	if err != nil {
		exitOnErr(err)
	}
	sdp.token = token

	// Get the default site, so we can reference it in our entitlement
	site, err := sdp.GetDefaultSite(ctx)
	if err != nil {
		exitOnErr(err)
	}

	entitlement, err := sdp.CreateEntitlement(ctx, site)
	if err != nil {
		exitOnErr(err)
	}
	fmt.Printf("Created entitlement - %s\n", entitlement.GetName())

	policy, err := sdp.CreatePolicy(ctx)
	if err != nil {
		exitOnErr(err)
	}
	fmt.Printf("Created policy - %s\n", policy.GetName())
}

func exitOnErr(err error) {
	if err != nil {
		log.Fatal(prettyPrintAPIError(err))
	}
}

type GenericErrorResponse struct {
	ID      string   `json:"id,omitempty"`
	Message string   `json:"message,omitempty"`
	Errors  []Errors `json:"errors,omitempty"`
}

type Errors struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

func prettyPrintAPIError(err error) error {
	if err, ok := err.(*openapi.GenericOpenAPIError); ok {
		model := err.Model()
		if err, ok := model.(openapi.Error); ok {
			return fmt.Errorf("%s - %s", err.GetId(), err.GetMessage())
		}
		if err, ok := model.(openapi.ValidationError); ok {
			var ValidationErrors string
			errorMessage := "Validation error"
			for _, ve := range err.GetErrors() {
				ValidationErrors = ValidationErrors + ve.GetField() + " " + ve.GetMessage() + "\n"
			}
			if msg, o := err.GetMessageOk(); o {
				errorMessage = fmt.Sprintf("%s %s", errorMessage, *msg)
			}
			return fmt.Errorf("%s \n %s", errorMessage, ValidationErrors)
		}

		errBody := GenericErrorResponse{}
		if errMarshal := json.Unmarshal(err.Body(), &errBody); errMarshal != nil {
			return err
		}
		if len(errBody.Message) > 0 {
			return errors.New(errBody.Message)
		}
		var genericErrors string
		for _, e := range errBody.Errors {
			genericErrors = genericErrors + e.Message + "\n"
		}

		return fmt.Errorf("%w %s", err, genericErrors)
	}
	return err
}

// login does basic auth to local provider
// This example does not include OTP or any other identity provider.
// For more examples, see:
// https://github.com/appgate/sdpctl/tree/2023.04.28/pkg/auth
// See:
// https://github.com/appgate/sdp-api-specification/blob/3f41903e2c53675dcc70c1fe242cd402683969c0/login.yml#L1-L35
func (s *Sdp) login(ctx context.Context, username, password string) (string, error) {
	loginOpts := openapi.LoginRequest{
		ProviderName: provider,
		Username:     openapi.PtrString(username),
		Password:     openapi.PtrString(password),
		// UUID to distinguish the Client device making the request.
		// It is supposed to be same for every login request from the same server.
		// the random UUID is just for demo purpose/
		DeviceId: "0fc572de-30aa-423a-a00a-577141bf9d8c",
	}
	response, _, err := s.client.LoginApi.LoginPost(ctx).LoginRequest(loginOpts).Execute()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Bearer %s", response.GetToken()), nil
}

// GetDefaultSite return the builtin site, "Default site" if it exists
func (s *Sdp) GetDefaultSite(ctx context.Context) (*openapi.Site, error) {
	response, _, err := s.client.SitesApi.SitesGet(ctx).Authorization(s.token).Execute()
	if err != nil {
		return nil, err
	}
	for _, site := range response.GetData() {
		if site.GetName() == "Default Site" {
			return &site, nil
		}
	}
	return nil, fmt.Errorf("could not find Default Site, found %d sites", len(response.GetData()))
}

// CreateEntitlement creates a new example Entitlement for appgate sdp
// See:
// https://github.com/appgate/sdp-api-specification/blob/3f41903e2c53675dcc70c1fe242cd402683969c0/entitlement.yml#L29-L42
func (s *Sdp) CreateEntitlement(ctx context.Context, site *openapi.Site) (*openapi.Entitlement, error) {
	// Required attributes:
	// - Name
	// - site
	// - conditions
	// - actions
	// https://github.com/appgate/sdp-api-specification/blob/3f41903e2c53675dcc70c1fe242cd402683969c0/entitlement.yml#L159-L367
	entitlement := openapi.Entitlement{
		Name:       "Example entitlement",
		Site:       site.GetId(),
		Conditions: []string{},
		Actions: []openapi.EntitlementAllOfActions{
			{
				Subtype: openapi.PtrString("icmp_up"),
				Action:  "allow",
				Hosts:   []string{"appgate.com"},
				Types:   []string{"0-255"},
			},
		},
	}

	response, _, err := s.client.EntitlementsApi.EntitlementsPost(ctx).Entitlement(entitlement).Authorization(s.token).Execute()
	if err != nil {
		return nil, err
	}
	return response, nil
}

// CreatePolicy creates a new example policy for appgate sdp
// See:
// https://github.com/appgate/sdp-api-specification/blob/3f41903e2c53675dcc70c1fe242cd402683969c0/policy.yml#L29-L58
func (s *Sdp) CreatePolicy(ctx context.Context) (*openapi.Policy, error) {
	// Required attributes:
	// - Name
	// - expression
	// See https://github.com/appgate/sdp-api-specification/blob/3f41903e2c53675dcc70c1fe242cd402683969c0/policy.yml#L159-L374
	policy := openapi.Policy{
		Name: "Example policy",
		Expression: `
		var result = false;
		if/*claims.user.groups*/(claims.user.groups && claims.user.groups.indexOf("developers") >= 0)/*end claims.user.groups*/ { return true; }
		if/*criteriaScript*/(admins(claims))/*end criteriaScript*/ { return true; }
		return result;
		`,
	}

	response, _, err := s.client.PoliciesApi.PoliciesPost(ctx).Policy(policy).Authorization(s.token).Execute()
	if err != nil {
		return nil, err
	}
	return response, nil
}
