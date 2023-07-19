// Code generated by pulumi-gen-awsconf DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/t0yv0/pulumi-12709/sdk/go/awsconf/internal"
)

type Configurer struct {
	pulumi.ResourceState

	AwsProvider aws.ProviderOutput `pulumi:"awsProvider"`
}

// NewConfigurer registers a new resource with the given unique name, arguments, and options.
func NewConfigurer(ctx *pulumi.Context,
	name string, args *ConfigurerArgs, opts ...pulumi.ResourceOption) (*Configurer, error) {
	if args == nil {
		args = &ConfigurerArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Configurer
	err := ctx.RegisterRemoteComponentResource("awsconf:index:Configurer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type configurerArgs struct {
	Profile *string `pulumi:"profile"`
	Region  *string `pulumi:"region"`
}

// The set of arguments for constructing a Configurer resource.
type ConfigurerArgs struct {
	Profile pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
}

func (ConfigurerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurerArgs)(nil)).Elem()
}

func (r *Configurer) ConfigureAwsProvider(ctx *pulumi.Context) (aws.ProviderOutput, error) {
	out, err := ctx.Call("awsconf:index:Configurer/awsMethod", nil, configurerConfigureAwsProviderResultOutput{}, r)
	if err != nil {
		return aws.ProviderOutput{}, err
	}
	return out.(configurerConfigureAwsProviderResultOutput).AwsProvider(), nil
}

type configurerConfigureAwsProviderResult struct {
	AwsProvider *aws.Provider `pulumi:"awsProvider"`
}

type configurerConfigureAwsProviderResultOutput struct{ *pulumi.OutputState }

func (configurerConfigureAwsProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*configurerConfigureAwsProviderResult)(nil)).Elem()
}

func (o configurerConfigureAwsProviderResultOutput) AwsProvider() aws.ProviderOutput {
	return o.ApplyT(func(v configurerConfigureAwsProviderResult) *aws.Provider { return v.AwsProvider }).(aws.ProviderOutput)
}

type ConfigurerInput interface {
	pulumi.Input

	ToConfigurerOutput() ConfigurerOutput
	ToConfigurerOutputWithContext(ctx context.Context) ConfigurerOutput
}

func (*Configurer) ElementType() reflect.Type {
	return reflect.TypeOf((**Configurer)(nil)).Elem()
}

func (i *Configurer) ToConfigurerOutput() ConfigurerOutput {
	return i.ToConfigurerOutputWithContext(context.Background())
}

func (i *Configurer) ToConfigurerOutputWithContext(ctx context.Context) ConfigurerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurerOutput)
}

// ConfigurerArrayInput is an input type that accepts ConfigurerArray and ConfigurerArrayOutput values.
// You can construct a concrete instance of `ConfigurerArrayInput` via:
//
//	ConfigurerArray{ ConfigurerArgs{...} }
type ConfigurerArrayInput interface {
	pulumi.Input

	ToConfigurerArrayOutput() ConfigurerArrayOutput
	ToConfigurerArrayOutputWithContext(context.Context) ConfigurerArrayOutput
}

type ConfigurerArray []ConfigurerInput

func (ConfigurerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Configurer)(nil)).Elem()
}

func (i ConfigurerArray) ToConfigurerArrayOutput() ConfigurerArrayOutput {
	return i.ToConfigurerArrayOutputWithContext(context.Background())
}

func (i ConfigurerArray) ToConfigurerArrayOutputWithContext(ctx context.Context) ConfigurerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurerArrayOutput)
}

// ConfigurerMapInput is an input type that accepts ConfigurerMap and ConfigurerMapOutput values.
// You can construct a concrete instance of `ConfigurerMapInput` via:
//
//	ConfigurerMap{ "key": ConfigurerArgs{...} }
type ConfigurerMapInput interface {
	pulumi.Input

	ToConfigurerMapOutput() ConfigurerMapOutput
	ToConfigurerMapOutputWithContext(context.Context) ConfigurerMapOutput
}

type ConfigurerMap map[string]ConfigurerInput

func (ConfigurerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Configurer)(nil)).Elem()
}

func (i ConfigurerMap) ToConfigurerMapOutput() ConfigurerMapOutput {
	return i.ToConfigurerMapOutputWithContext(context.Background())
}

func (i ConfigurerMap) ToConfigurerMapOutputWithContext(ctx context.Context) ConfigurerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurerMapOutput)
}

type ConfigurerOutput struct{ *pulumi.OutputState }

func (ConfigurerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Configurer)(nil)).Elem()
}

func (o ConfigurerOutput) ToConfigurerOutput() ConfigurerOutput {
	return o
}

func (o ConfigurerOutput) ToConfigurerOutputWithContext(ctx context.Context) ConfigurerOutput {
	return o
}

func (o ConfigurerOutput) AwsProvider() aws.ProviderOutput {
	return o.ApplyT(func(v *Configurer) aws.ProviderOutput { return v.AwsProvider }).(aws.ProviderOutput)
}

type ConfigurerArrayOutput struct{ *pulumi.OutputState }

func (ConfigurerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Configurer)(nil)).Elem()
}

func (o ConfigurerArrayOutput) ToConfigurerArrayOutput() ConfigurerArrayOutput {
	return o
}

func (o ConfigurerArrayOutput) ToConfigurerArrayOutputWithContext(ctx context.Context) ConfigurerArrayOutput {
	return o
}

func (o ConfigurerArrayOutput) Index(i pulumi.IntInput) ConfigurerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Configurer {
		return vs[0].([]*Configurer)[vs[1].(int)]
	}).(ConfigurerOutput)
}

type ConfigurerMapOutput struct{ *pulumi.OutputState }

func (ConfigurerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Configurer)(nil)).Elem()
}

func (o ConfigurerMapOutput) ToConfigurerMapOutput() ConfigurerMapOutput {
	return o
}

func (o ConfigurerMapOutput) ToConfigurerMapOutputWithContext(ctx context.Context) ConfigurerMapOutput {
	return o
}

func (o ConfigurerMapOutput) MapIndex(k pulumi.StringInput) ConfigurerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Configurer {
		return vs[0].(map[string]*Configurer)[vs[1].(string)]
	}).(ConfigurerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurerInput)(nil)).Elem(), &Configurer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurerArrayInput)(nil)).Elem(), ConfigurerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurerMapInput)(nil)).Elem(), ConfigurerMap{})
	pulumi.RegisterOutputType(ConfigurerOutput{})
	pulumi.RegisterOutputType(configurerConfigureAwsProviderResultOutput{})
	pulumi.RegisterOutputType(ConfigurerArrayOutput{})
	pulumi.RegisterOutputType(ConfigurerMapOutput{})
}
