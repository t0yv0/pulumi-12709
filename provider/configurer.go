package provider

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

type Configurer struct {
	pulumi.ResourceState

	AwsProvider aws.ProviderOutput `pulumi:"awsProvider"`
}

type ConfigurerArgs struct {
	Region  pulumi.StringInput `pulumi:"region"`
	Profile pulumi.StringInput `pulumi:"profile"`
}

func NewConfigurer(
	ctx *pulumi.Context,
	name string,
	args *ConfigurerArgs,
	opts ...pulumi.ResourceOption,
) (*Configurer, error) {
	resource := &Configurer{}
	if err := ctx.RegisterComponentResource(ConfigurerToken, name, resource, opts...); err != nil {
		return nil, err
	}

	awsProvName := fmt.Sprintf("%s-aws-provider", name)

	awsProvArgs := &aws.ProviderArgs{
		Region:  args.Region,
		Profile: args.Profile,
	}

	awsProv, err := aws.NewProvider(ctx, awsProvName, awsProvArgs)
	if err != nil {
		return nil, err
	}

	resource.AwsProvider = awsProv.ToProviderOutput()

	if err := ctx.RegisterResourceOutputs(resource, pulumi.Map{
		"awsProvider": resource.AwsProvider,
	}); err != nil {
		return nil, err
	}

	return resource, nil
}

func ConstructConfigurer(
	ctx *pulumi.Context,
	name string,
	inputs provider.ConstructInputs,
	opts ...pulumi.ResourceOption,
) (*provider.ConstructResult, error) {
	args := &ConfigurerArgs{}
	if err := inputs.CopyTo(args); err != nil {
		return nil, err
	}

	configurer, err := NewConfigurer(ctx, name, args, opts...)
	if err != nil {
		return nil, err
	}

	return provider.NewConstructResult(configurer)
}
