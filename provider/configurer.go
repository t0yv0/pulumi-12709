package provider

import (
	"context"
	"sync"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/internals"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

type Configurer struct {
	pulumi.ResourceState

	AwsRegion  string
	AwsProfile string
}

type ConfigurerArgs struct {
	AwsRegion  string                `pulumi:"awsRegion"`
	AwsProfile string                `pulumi:"awsProfile"`
	Mode       pulumi.StringPtrInput `pulumi:"mode"`
}

func NewConfigurer(
	ctx *pulumi.Context,
	name string,
	args *ConfigurerArgs,
	opts ...pulumi.ResourceOption,
) (*Configurer, error) {
	resource := &Configurer{
		AwsRegion:  args.AwsRegion,
		AwsProfile: args.AwsProfile,
	}
	if err := ctx.RegisterComponentResource(ConfigurerToken, name, resource, opts...); err != nil {
		return nil, err
	}

	if err := ctx.RegisterResourceOutputs(resource, pulumi.Map{}); err != nil {
		return nil, err
	}

	registerConfigurer(ctx.Context(), resource)

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

type ConfigureAwsMethodArgs struct{}

type ConfigureAwsMethodResult struct {
	AwsProvider *aws.Provider `pulumi:"awsProvider"`
}

func CallConfigureAwsMethod(ctx *pulumi.Context, inputs provider.CallArgs) (*provider.CallResult, error) {
	args := &ConfigureAwsMethodArgs{}
	res, err := inputs.CopyTo(args)
	if err != nil {
		return nil, err
	}

	self := lookupConfigurer(ctx.Context(), res.URN())

	awsProv, err := aws.NewProvider(ctx, "aws-p", &aws.ProviderArgs{
		Region:  pulumi.String(self.AwsRegion),
		Profile: pulumi.String(self.AwsProfile),
	}, pulumi.Parent(self))
	if err != nil {
		return nil, err
	}

	result := &ConfigureAwsMethodResult{
		AwsProvider: awsProv,
	}

	// The following code resolves to Unknown.
	// if ctx.DryRun() {
	// 	result.SomeString = awsProv.HttpProxy.ToStringPtrOutput().ApplyT(func(x *string) string {
	// 		if x != nil {
	// 			return fmt.Sprintf("OK: mode was %q", *x)
	// 		}
	// 		return "OK: mode was nil"
	// 	}).(pulumi.StringOutput)
	// } else {
	// 	result.SomeString = pulumi.String("OK").ToStringOutput()
	// }

	return provider.NewCallResult(result)
}

var (
	configurerByURN sync.Map
)

func registerConfigurer(ctx context.Context, c *Configurer) {
	key := forceURN(ctx, c.URN())
	configurerByURN.Store(key, c)
}

func lookupConfigurer(ctx context.Context, urn pulumi.URNOutput) *Configurer {
	theURN := forceURN(ctx, urn)
	v, ok := configurerByURN.Load(theURN)
	contract.Assertf(ok, "lookupConfigurer: unknown URN %q", theURN)
	c, ok := v.(*Configurer)
	contract.Assertf(ok, "lookupConfigurer: type mismatch")
	return c
}

func forceURN(ctx context.Context, urnOutput pulumi.Output) pulumi.URN {
	r, err := internals.UnsafeAwaitOutput(ctx, urnOutput)
	contract.AssertNoErrorf(err, "forceURN: UnsafeAwaitOutput failed")
	contract.Assertf(r.Known, "forceURN: URN should be known")
	urn, ok := r.Value.(pulumi.URN)
	contract.Assertf(ok, "forceURN: r.Value should be a pulumi.URN")
	return urn
}
