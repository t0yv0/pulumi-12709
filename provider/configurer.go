package provider

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

type Configurer struct {
	pulumi.ResourceState

	AwsProvider aws.ProviderOutput `pulumi:"awsProvider"`
}

type ConfigurerArgs struct{}

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

	if err := ctx.RegisterResourceOutputs(resource, pulumi.Map{}); err != nil {
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

type ConfigureAwsMethodArgs struct {
	Region  pulumi.StringInput `pulumi:"region"`
	Profile pulumi.StringInput `pulumi:"profile"`
}

type ConfigureAwsMethodResult struct {
	AwsProvider *aws.Provider `pulumi:"awsProvider"`
}

func CallConfigureAwsMethod(ctx *pulumi.Context, inputs provider.CallArgs) (*provider.CallResult, error) {
	args := &ConfigureAwsMethodArgs{}
	self, err := inputs.CopyTo(args)
	if err != nil {
		return nil, err
	}

	awsProvArgs := &aws.ProviderArgs{
		Region:  args.Region,
		Profile: args.Profile,
	}

	awsProv, err := aws.NewProvider(ctx, "aws-p", awsProvArgs, pulumi.Parent(self))
	if err != nil {
		return nil, err
	}

	return provider.NewCallResult(&ConfigureAwsMethodResult{
		AwsProvider: awsProv,
	})
}
