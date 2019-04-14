// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package infoblox

import (
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "infoblox"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
// func makeDataSource(mod string, res string) tokens.ModuleMember {
// 	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
// 	return makeMember(mod+"/"+fn, res)
// }

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
// func boolRef(b bool) *bool {
// 	return &b
// }

// stringValue gets a string value from a property map if present, else ""
// func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
// 	val, ok := vars[prop]
// 	if ok && val.IsString() {
// 		return val.StringValue()
// 	}
// 	return ""
// }

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
// var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := infoblox.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "infoblox",
		Description: "A Pulumi package for creating and managing infoblox resources.",
		Keywords:    []string{"pulumi", "infoblox"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-infoblox",
		GitHubOrg:   "sky-uk",
		Config: map[string]*tfbridge.SchemaInfo{

			"username": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"INFOBLOX_USERNAME"},
				},
			},
			// "username": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Required:    true,
			// 	DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_USERNAME", nil),
			// 	Description: "User to authenticate with Infoblox appliance",
			// },

			"password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"INFOBLOX_PASSWORD"},
				},
			},
			// "password": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Required:    true,
			// 	DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_PASSWORD", nil),
			// 	Description: "Password to authenticate with Infoblox appliance",
			// },

			"server": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"INFOBLOX_SERVER"},
				},
			},
			// "server": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Required:    true,
			// 	DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_SERVER", nil),
			// 	Description: "Infoblox appliance to connect to eg https://192.168.0.1",
			// },

			"allow_unverified_ssl": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"INFOBLOX_ALLOW_UNVERIFIED_SSL"},
				},
			},
			// "allow_unverified_ssl": &schema.Schema{
			// 	Type:        schema.TypeBool,
			// 	Optional:    true,
			// 	DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_ALLOW_UNVERIFIED_SSL", false),
			// 	Description: "If set, Infoblox client will permit unverifiable SSL certificates.",
			// },

			"wapi_version": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"INFOBLOX_WAPI_VERSION"},
				},
			},
			// "wapi_version": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Description: "Infoblox WAPI server version, defaults to v2.6.1",
			// 	Default:     "v2.6.1",
			// },

			"timeout": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"INFOBLOX_CLIENT_TIMEOUT"},
				},
			},
			// "timeout": &schema.Schema{
			// 	Type:        schema.TypeInt,
			// 	Optional:    true,
			// 	Description: "http response timeout, in seconds",
			// 	DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_CLIENT_TIMEOUT", 0),
			// },

			"client_debug": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"INFOBLOX_CLIENT_DEBUG"},
				},
			},
			// "client_debug": &schema.Schema{
			// 	Type:        schema.TypeBool,
			// 	Optional:    true,
			// 	DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_CLIENT_DEBUG", false),
			// 	Description: "infoblox client debug",
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// https://github.com/sky-uk/terraform-provider-infoblox/blob/master/infoblox/provider.go
			"infoblox_cname_record":          {Tok: makeResource(mainMod, "CNameRecord")},
			"infoblox_arecord":               {Tok: makeResource(mainMod, "ARecord")},
			"infoblox_srv_record":            {Tok: makeResource(mainMod, "SRVRecord")},
			"infoblox_txtrecord":             {Tok: makeResource(mainMod, "TXTRecord")},
			"infoblox_network":               {Tok: makeResource(mainMod, "Network")},
			"infoblox_zone_auth":             {Tok: makeResource(mainMod, "ZoneAuth")},
			"infoblox_dhcp_range":            {Tok: makeResource(mainMod, "DHCPRange")},
			"infoblox_admin_user":            {Tok: makeResource(mainMod, "AdminUser")},
			"infoblox_admin_group":           {Tok: makeResource(mainMod, "AdminGroup")},
			"infoblox_admin_role":            {Tok: makeResource(mainMod, "AdminRole")},
			"infoblox_ns_record":             {Tok: makeResource(mainMod, "NSRecord")},
			"infoblox_zone_delegated":        {Tok: makeResource(mainMod, "ZoneDelegated")},
			"infoblox_permission":            {Tok: makeResource(mainMod, "Permission")},
			"infoblox_zone_stub":             {Tok: makeResource(mainMod, "ZoneStub")},
			"infoblox_zone_forward":          {Tok: makeResource(mainMod, "ZoneForward")},
			"infoblox_ns_group_delegation":   {Tok: makeResource(mainMod, "NSGroupDelegation")},
			"infoblox_ns_group_forward":      {Tok: makeResource(mainMod, "NSGroupForward")},
			"infoblox_ns_group_stub":         {Tok: makeResource(mainMod, "NSGroupStub")},
			"infoblox_ns_group_forward_stub": {Tok: makeResource(mainMod, "NSGroupForwardStub")},
			"infoblox_mx_record":             {Tok: makeResource(mainMod, "MXRecord")},
			"infoblox_dns_view":              {Tok: makeResource(mainMod, "DNSView")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: makeDataSource(mainMod, "getAmi")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "latest",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			// Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=0.16.4,<0.17.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
