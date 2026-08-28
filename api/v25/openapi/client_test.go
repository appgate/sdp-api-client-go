package openapi

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func client() *APIClient {
	return NewAPIClient(NewConfiguration())
}

func TestConfigCustomUserAgent(t *testing.T) {
	cfg := &Configuration{
		UserAgent: "appgate",
		DefaultHeader: map[string]string{
			"Accept": "Default accept",
		},
	}
	client := NewAPIClient(cfg)

	request, err := client.prepareRequest(
		context.TODO(),
		"appgate.com/get",
		http.MethodGet,
		nil,
		map[string]string{},
		nil,
		nil,
		nil)
	if err != nil {
		t.Fatalf("got error, expected none %s", err)
	}
	if v, ok := request.Header["Accept"]; ok && v[0] != "Default accept" {
		t.Fatalf("Expected default accept to be Default accept, got %s", v)
	}
	if v, ok := request.Header["User-Agent"]; ok && v[0] != "appgate" {
		t.Fatalf("Expected User-Agent to be appgate, got %s", v)
	}

	secondRequest, err := client.prepareRequest(
		context.WithValue(context.TODO(), ContextAcceptHeader, "NEW"),
		"appgate.com/get",
		http.MethodGet,
		nil,
		map[string]string{},
		nil,
		nil,
		nil)
	if err != nil {
		t.Fatalf("got error, expected none %s", err)
	}
	if v, ok := secondRequest.Header["Accept"]; ok && v[0] != "NEW" {
		t.Fatalf("Expected default accept to be NEW, got %s", v)
	}
}

func TestAPIClientHeaders(t *testing.T) {
	c := client()
	tests := []struct {
		name         string
		contextKey   contextKey
		contextValue string
		wantHeaders  http.Header
		wantErr      bool
	}{
		{
			name:         "Test overwrite Accept header",
			contextKey:   ContextAcceptHeader,
			contextValue: "application/vnd.appgate.peer-v99+json",
			wantHeaders: http.Header{
				"Accept":     []string{"application/vnd.appgate.peer-v99+json"},
				"User-Agent": []string{"OpenAPI/0.0.1/go"},
			},
			wantErr: false,
		},
		{
			name:         "basic auth",
			contextKey:   ContextAccessToken,
			contextValue: "b64string",
			wantHeaders: http.Header{
				"User-Agent":    []string{"OpenAPI/0.0.1/go"},
				"Authorization": []string{"Bearer b64string"},
			},
			wantErr: false,
		},
		{
			name: "default",
			wantHeaders: http.Header{
				"User-Agent": []string{"OpenAPI/0.0.1/go"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLocalVarRequest, err := c.prepareRequest(
				context.WithValue(context.TODO(), tt.contextKey, tt.contextValue),
				"appgate.com/get",
				http.MethodGet,
				nil,
				map[string]string{},
				nil,
				nil,
				nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("APIClient.prepareRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLocalVarRequest.Header, tt.wantHeaders) {
				t.Errorf("APIClient.prepareRequest() = %v, want  %v", gotLocalVarRequest.Header, tt.wantHeaders)
			}
		})
	}
}
