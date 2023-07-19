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
	Region  pulumi.StringInput    `pulumi:"region"`
	Profile pulumi.StringInput    `pulumi:"profile"`
	Mode    pulumi.StringPtrInput `pulumi:"mode"`
}

type ConfigureAwsMethodResult struct {
	AwsProvider *aws.Provider       `pulumi:"awsProvider"`
	SomeString  pulumi.StringOutput `pulumi:"someString"`
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

	// The following code resolves to Unknown.

	result := &ConfigureAwsMethodResult{
		AwsProvider: awsProv,
	}

	if ctx.DryRun() {
		result.SomeString = awsProv.HttpProxy.ToStringPtrOutput().ApplyT(func(x *string) string {
			if x != nil {
				return fmt.Sprintf("OK: mode was %q", *x)
			}
			return "OK: mode was nil"
		}).(pulumi.StringOutput)
	} else {
		result.SomeString = pulumi.String("OK").ToStringOutput()
	}

	return provider.NewCallResult(result)
}
