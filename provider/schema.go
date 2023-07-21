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
	self         = "__self__"
)

var (
	stringTS = schema.TypeSpec{Type: "string"}
)

func PackageSpec() schema.PackageSpec {
	return schema.PackageSpec{
		Name: ProviderName,

		Resources: map[string]schema.ResourceSpec{
			ConfigurerToken: {
				IsComponent: true,
				Methods: map[string]string{
					"configureAwsProvider": ConfigurerConfigureAwsMethodToken,
				},
			},
		},

		Functions: map[string]schema.FunctionSpec{
			ConfigurerConfigureAwsMethodToken: {
				XReturnPlainResource: true,
				Inputs: &schema.ObjectTypeSpec{
					Properties: map[string]schema.PropertySpec{
						self: {
							TypeSpec: schema.TypeSpec{
								Ref: localResourceRef(ConfigurerToken),
							},
						},
						"region":  {TypeSpec: stringTS},
						"profile": {TypeSpec: stringTS},
						"mode":    {TypeSpec: stringTS},
					},
					Required: []string{self, "region", "profile"},
				},
				ReturnType: &schema.ReturnTypeSpec{
					ObjectTypeSpec: &schema.ObjectTypeSpec{
						Properties: map[string]schema.PropertySpec{
							"awsProvider": {
								TypeSpec: schema.TypeSpec{Ref: awsRef("#/provider")},
							},
						},
						Required: []string{"awsProvider"},
					},
				},
			},
		},

		Language: map[string]schema.RawMessage{
			"go": rawMessage(map[string]interface{}{
				"generateResourceContainerTypes": true,
				"importBasePath":                 "github.com/t0yv0/pulumi-12709/sdk/go/awsconf",
			}),
			"nodejs": rawMessage(map[string]interface{}{
				"dependencies": map[string]interface{}{
					"@pulumi/aws": awsVersion,
				},
			}),
		},
	}
}

func localResourceRef(token string) string {
	return fmt.Sprintf("#/resources/%s", token)
}

func awsRef(ref string) string {
	return fmt.Sprintf("/aws/v%s/schema.json%s", awsVersion, ref)
}

func rawMessage(v interface{}) schema.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}
