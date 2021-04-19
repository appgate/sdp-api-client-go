package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/appgate/sdp-api-client-go/api/v13/openapi"
	"github.com/google/uuid"
)

const (
	version  = 13
	timeout  = 10
	insecure = true
	provider = "local"
)

func main() {
	distinguishedName := flag.String("distinguishedName", "", "a distinguished Name")
	address := flag.String("addr", "", "controller.appgate.com/admin")
	password := flag.String("password", "", "api password")
	username := flag.String("username", "", "api username")
	debug := flag.Bool("debug", false, "http debug")
	flag.Parse()

	timeoutDuration := time.Duration(timeout)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecure,
		},
		Dial: (&net.Dialer{
			Timeout: timeoutDuration * time.Second,
		}).Dial,
		TLSHandshakeTimeout: timeoutDuration * time.Second,
	}

	httpclient := &http.Client{
		Transport: tr,
		Timeout:   ((timeoutDuration * 2) * time.Second),
	}

	clientCfg := &openapi.Configuration{
		DefaultHeader: map[string]string{
			"Accept": fmt.Sprintf("application/vnd.appgate.peer-v%d+json", version),
		},
		Debug: *debug,
		Servers: []openapi.ServerConfiguration{
			{
				URL: *address,
			},
		},
		HTTPClient: httpclient,
	}

	client := openapi.NewAPIClient(clientCfg)
	ctx := context.TODO()

	loginOpts := openapi.LoginRequest{
		ProviderName: provider,
		Username:     username,
		Password:     password,
		DeviceId:     uuid.New().String(),
	}

	loginResponse, _, err := client.LoginApi.LoginPost(ctx).LoginRequest(loginOpts).Execute()
	if err != nil {
		fmt.Printf("Failed to get token %s\n", err)
		os.Exit(1)
	}

	token := fmt.Sprintf("Bearer %s", *openapi.PtrString(*loginResponse.Token))
	api := client.LicensedUsersApi

	fmt.Printf("Delete DN %q\n", *distinguishedName)
	request := api.LicenseUsersDistinguishedNameDelete(ctx, *distinguishedName)
	response, err := request.Authorization(token).Execute()
	if err != nil {
		fmt.Printf("Got error %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("response: %+v", response)
}
