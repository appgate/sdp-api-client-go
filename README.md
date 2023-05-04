# Appgate SDP golang api client

this repository contains a Go API client for the Appgate API.
The code is generated using [openapi-generator](https://github.com/OpenAPITools/openapi-generator) and [apigentools](https://github.com/DataDog/apigentools).



## Version compatibility matrix

Each appliance version has a minimum of compatibility with 2 API versions back.

| Appliance Version 	| API Version 	|
|-------------------	|-------------	|
| 5.1.*             	| 12          	|
| 5.2.*             	| 13          	|
| 5.3.*             	| 14          	|
| 5.4.*             	| 15          	|
| 5.5.*             	| 16          	|
| 6.0.*             	| 17          	|
| 6.1.*             	| 18          	|
| 6.2.*             	| 19          	|




## Getting started

### Example usage

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


### Code generation

#### Requirements
	- go 1.20
	- python
	- [apigentools](https://github.com/DataDog/apigentools) v1.6.0
	- docker (with userns enabled)

```sh
bash generate_spec.bash
```

