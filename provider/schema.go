package provider

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
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

		Language: map[string]schema.RawMessage{
			"go": rawMessage(map[string]interface{}{
				"generateResourceContainerTypes": true,
				"importBasePath":                 "github.com/t0yv0/pulumi-12709/sdk/go/awsconf",
				"liftSingleValueMethodReturns":   true,
			}),
		},
	}
}

func awsRef(ref string) string {
	return fmt.Sprintf("/aws/v%s/schema.json%s", awsVersion, ref)
}

func rawMessage(v interface{}) schema.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}
