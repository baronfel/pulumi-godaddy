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

package godaddy

import (
	"unicode"

	"github.com/n3integration/terraform-godaddy/plugin/terraform-godaddy"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/tokens"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "godaddy"
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
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

func refProviderLicense(license tfbridge.TFProviderLicense) *tfbridge.TFProviderLicense {
	return &license
}

func Provider() tfbridge.ProviderInfo {
	p := godaddy.Provider()
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "godaddy",
		GitHubOrg:         "godaddy",
		Description:       "A Pulumi package for creating and managing GoDaddy DNS resources.",
		Keywords:          []string{"pulumi", "godaddy"},
		License:           "Apache-2.0",
		TFProviderLicense: refProviderLicense(tfbridge.MITLicenseType),
		Homepage:          "https://pulumi.io",
		Repository:        "https://github.com/baronfel/pulumi-godaddy",
		Config: map[string]*tfbridge.SchemaInfo{
			"key": {
				Type: makeType("key", "Key"),
				Default: &tfBridge.DefaultInfo{
					EnvVars: []string{"GODADDY_API_KEY"},
				},
			},
			"secret": {
				Type: makeType("secret", "Secret"),
				Default: &tfBridge.DefaultInfo{
					EnvVars: []string{"GODADDY_API_SECRET"},
				},
			},
			"baseURL": {
				Type: makeType("baseURL", "BaseURL"),
				Default: &tfBridge.DefaultInfo{
					Value: string{"https://api.godaddy.com"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"godaddy_domain_record": {Tok: makeResource(mainMod, "DomainRecord")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"godaddy_domain_record": {Tok: makeDataSource(mainMod, "getDomainRecord")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^1.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=1.0.0,<2.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "1.5.0-*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				mainPkg: "GoDaddy",
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
