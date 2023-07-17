package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	rprovider "github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

const (
	providerName = "aws-configurer"
	version      = "0.0.1"
	awsVersion   = "5.31.0"
)

func main() {
	err := rprovider.MainWithOptions(rprovider.Options{
		Name:      providerName,
		Version:   version,
		Schema:    providerSchema(),
		Construct: construct,
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}

func construct(
	ctx *pulumi.Context,
	typ,
	name string,
	inputs provider.ConstructInputs,
	options pulumi.ResourceOption,
) (*provider.ConstructResult, error) {
	return nil, fmt.Errorf("TODO: Construct")
}

func providerSchema() []byte {
	p := packageSpec()
	bytes, err := json.Marshal(p)
	contract.AssertNoError(err)
	return bytes
}

func packageSpec() schema.PackageSpec {
	return schema.PackageSpec{
		Name: providerName,
		Resources: map[string]schema.ResourceSpec{
			"cfg:index:Configurer": {
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
