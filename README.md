# Appgate SDP golang api client

this repository contains a Go API client for the Appgate API.
The code is generated using [openapi-generator](https://github.com/OpenAPITools/openapi-generator)



|                         	|  client version 12 	| client version 13 	| client version 14 	| client version 15 	|
|-------------------------	|--------------------	|-------------------	|-------------------	|-------------------	|
| Appliance version 5.1.* 	| Full support       	|                   	|                   	|                   	|
| Appliance version 5.2.* 	| Partial support    	| Full support      	| Partial support      	|                   	|
| Appliance version 5.3.* 	| Partial support   	| Partial support   	| Full support      	| Partial support      	|
| Appliance version 5.4.* 	| Partial support   	| Partial support   	| Partial support      	| Full support      	|

## Getting started

Example usage how to create the client and get auth token.

```golang
package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/appgate/sdp-api-client-go/api/v14/openapi"
	"github.com/google/uuid"
)

const (
	version  = 14
	timeout  = 10
	insecure = true
	provider = "local"
	username = "admin"
	password = "admin"
	address  = "https://controller.appgate.com/admin"
)

func main() {
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
		Debug: true,
		Servers: []openapi.ServerConfiguration{
			{
				URL: address,
			},
		},
		HTTPClient: httpclient,
	}

	client := openapi.NewAPIClient(clientCfg)
	ctx := context.TODO()

	loginOpts := openapi.LoginRequest{
		ProviderName: provider,
		Username:     openapi.PtrString(username),
		Password:     openapi.PtrString(password),
		DeviceId:     uuid.New().String(),
	}

	loginResponse, _, err := client.LoginApi.LoginPost(ctx).LoginRequest(loginOpts).Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Bearer %s", *openapi.PtrString(*loginResponse.Token))
}
```
