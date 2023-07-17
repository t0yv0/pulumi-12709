package provider

import (
	"fmt"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

const (
	ProviderName = "awsconf"
	awsVersion   = "5.31.0"
)

func PackageSpec() schema.PackageSpec {
	return schema.PackageSpec{
		Name: ProviderName,
		Resources: map[string]schema.ResourceSpec{
			"awsconf:index:Configurer": {
				IsComponent: true,
				InputProperties: map[string]schema.PropertySpec{
					"region":  {TypeSpec: schema.TypeSpec{Type: "string"}},
					"profile": {TypeSpec: schema.TypeSpec{Type: "string"}},
				},
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Properties: map[string]schema.PropertySpec{
						"awsProvider": {
							TypeSpec: schema.TypeSpec{Ref: awsRef("#/provider")},
						},
					},
				},
			},
		},
	}
}

func awsRef(ref string) string {
	return fmt.Sprintf("/aws/v%s/schema.json%s", awsVersion, ref)
}
