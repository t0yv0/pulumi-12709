package main

import (
	"encoding/json"
	"fmt"

	rprovider "github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
	awsconf "github.com/t0yv0/pulumi-12709/provider"
)

const (
	version = "0.0.1"
)

func main() {
	err := rprovider.MainWithOptions(rprovider.Options{
		Name:      awsconf.ProviderName,
		Version:   version,
		Schema:    providerSchema(),
		Construct: constructFunc,
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}

func constructFunc(
	ctx *pulumi.Context,
	typ, name string,
	inputs provider.ConstructInputs,
	options pulumi.ResourceOption,
) (*provider.ConstructResult, error) {
	switch {
	case typ == awsconf.ConfigurerToken:
		return awsconf.ConstructConfigurer(ctx, name, inputs, options)
	default:
		return nil, fmt.Errorf("Cannot Construct a remote component resource: unknown type token %q", typ)
	}
}

func providerSchema() []byte {
	p := awsconf.PackageSpec()
	bytes, err := json.Marshal(p)
	contract.AssertNoError(err)
	return bytes
}
