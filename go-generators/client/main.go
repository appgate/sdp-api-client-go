//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	verbose = flag.Bool("v", false, "Print verbose log messages")
)

type fileTemplate struct {
	name       string
	minVersion int
}

var files = []fileTemplate{
	// bigInt issues
	// modified the original generated code and replaced IpPool.Total from int64 to bigInt
	// workaround for https://github.com/OpenAPITools/openapi-generator/issues/10620
	{
		name:       "model_ip_pool.template",
		minVersion: 16,
	},
	{
		name:       "model_ip_pool_all_of.template",
		minVersion: 16,
	},

	// Custom model structs for Each identity providers
	{
		name:       "model_connector_identity_provider_list.template",
		minVersion: 16,
	},
	{
		name:       "model_ldap_identity_provider_list.template",
		minVersion: 16,
	},
	{
		name:       "model_radius_identity_provider_list.template",
		minVersion: 16,
	},
	{
		name:       "model_ldap_certificate_identity_provider_list.template",
		minVersion: 16,
	},
	{
		name:       "model_local_database_identity_provider_list.template",
		minVersion: 16,
	},
	{
		name:       "model_saml_identity_provider_list.template",
		minVersion: 16,
	},
	{
		name:       "model_oidc_identity_provider_list.template",
		minVersion: 19,
	},

	// Unit test for client to test custom patches
	{
		name:       "client_test.template",
		minVersion: 16,
	},
}

type identityProvider struct {
	name, filename string
	minVersion     int
}

// the templates for identity_providers is because of how the openapi-generator treat discriminator objects.
// The default behavior within openapi-generator is to create 1 api service for all, but that does not really work well
// in practice because we need a way in the code to have separate models for each type of identity providers,
// that is why we create a copy for each type in the discriminator object based on the default api_identity_providers.go
// the only difference is the service name and the models used in the arguments, the patch template-patches/go/identity_providers.patch
// will append the extra services to the client.
var identityProviders = []identityProvider{
	{
		name:       "Connector",
		filename:   "api_connector_identity_providers",
		minVersion: 18,
	},
	{
		name:       "Ldap",
		filename:   "api_ldap_identity_providers",
		minVersion: 18,
	},
	{
		name:       "LdapCertificate",
		filename:   "api_ldap_certificate_identity_providers",
		minVersion: 18,
	},
	{
		name:       "LocalDatabase",
		filename:   "api_local_database_identity_providers",
		minVersion: 18,
	},
	{
		name:       "Radius",
		filename:   "api_radius_identity_providers",
		minVersion: 18,
	},
	{
		name:       "Saml",
		filename:   "api_saml_identity_providers",
		minVersion: 18,
	},
	{
		name:       "Oidc",
		filename:   "api_oidc_identity_providers",
		minVersion: 19,
	},
}

func logf(fmt string, args ...interface{}) {
	if *verbose {
		log.Printf(fmt, args...)
	}
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	var version int
	flag.IntVar(&version, "version", 0, "API Version")
	flag.Parse()
	if version == 0 {
		die(errors.New("Requires version flag, got None"))
	}
	logf("Generating for version %+v", version)
	for _, tmpl := range files {
		if tmpl.minVersion <= version {
			name := strings.TrimSuffix(tmpl.name, ".template")
			data := struct {
				Version int
			}{
				Version: version,
			}

			f, err := writeFile(name, tmpl.name, version, data)
			if err != nil {
				die(err)
			}
			logf("created %s for v%d", f.Name(), version)
		}
	}
	for _, ip := range identityProviders {
		data := struct {
			Version              int
			IdentityProviderName string
		}{
			Version:              version,
			IdentityProviderName: ip.name,
		}
		if ip.minVersion <= version {
			f, err := writeFile(ip.filename, "api_identity_providers.template", version, data)
			if err != nil {
				die(err)
			}
			logf("created %s", f.Name())
		}
	}
	logf("done")
}

func writeFile(name, tmpl string, version int, data interface{}) (*os.File, error) {
	f, err := os.Create(fmt.Sprintf("api/v%d/openapi/%s.go", version, name))
	if err != nil {
		return f, err
	}
	defer f.Close()
	cwd, err := os.Getwd()
	if err != nil {
		return f, err
	}
	body, err := os.ReadFile(filepath.Join(cwd, "go-generators/client", tmpl))
	if err != nil {
		return f, err
	}
	goTemplate, err := template.New("").Parse(string(body))
	if err != nil {
		return f, fmt.Errorf("template New err %w", err)
	}

	var buf bytes.Buffer
	if err := goTemplate.Execute(&buf, data); err != nil {
		return f, fmt.Errorf("template err %w", err)
	}
	p, err := format.Source(buf.Bytes())
	if err != nil {
		return f, fmt.Errorf("format err %w", err)
	}
	if _, err := f.Write(p); err != nil {
		return f, fmt.Errorf("write err %w", err)
	}

	return f, nil
}
