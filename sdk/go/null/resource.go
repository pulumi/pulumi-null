// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package null

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-null/sdk/go/null/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v1/go/aws"
//	"github.com/pulumi/pulumi-null/sdk/go/null"
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// var cluster []*index.Instance
//
//	for index := 0; index < 3; index++ {
//	    key0 := index
//	    _ := index
//
// __res, err := index.NewInstance(ctx, fmt.Sprintf("cluster-%v", key0), &index.InstanceArgs{
// Ami: "ami-0dcc1e21636832c5d",
// InstanceType: "m5.large",
// })
// if err != nil {
// return err
// }
// cluster = append(cluster, __res)
// }
// invokeJoin, err := std.Join(ctx, var splat0 []interface{}
// for _, val0 := range cluster {
// splat0 = append(splat0, val0.Id)
// }
// &std.JoinArgs{
// Separator: ",",
// Input: splat0,
// }, nil)
// if err != nil {
// return err
// }
// // The primary use-case for the null resource is as a do-nothing container
// // for arbitrary actions taken by a provisioner.
// //
// // In this example, three EC2 instances are created and then a
// // null_resource instance is used to gather data about all three
// // and execute a single action that affects them all. Due to the triggers
// // map, the null_resource will be replaced each time the instance ids
// // change, and thus the remote-exec provisioner will be re-run.
// _, err = null.NewResource(ctx, "cluster", &null.ResourceArgs{
// Triggers: pulumi.StringMap{
// "cluster_instance_ids": invokeJoin.Result,
// },
// })
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
type Resource struct {
	pulumi.CustomResourceState

	// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
	Triggers pulumi.StringMapOutput `pulumi:"triggers"`
}

// NewResource registers a new resource with the given unique name, arguments, and options.
func NewResource(ctx *pulumi.Context,
	name string, args *ResourceArgs, opts ...pulumi.ResourceOption) (*Resource, error) {
	if args == nil {
		args = &ResourceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Resource
	err := ctx.RegisterResource("null:index/resource:Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResource gets an existing Resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceState, opts ...pulumi.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("null:index/resource:Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resource resources.
type resourceState struct {
	// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
	Triggers map[string]string `pulumi:"triggers"`
}

type ResourceState struct {
	// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
	Triggers pulumi.StringMapInput
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
	Triggers map[string]string `pulumi:"triggers"`
}

// The set of arguments for constructing a Resource resource.
type ResourceArgs struct {
	// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
	Triggers pulumi.StringMapInput
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceInput interface {
	pulumi.Input

	ToResourceOutput() ResourceOutput
	ToResourceOutputWithContext(ctx context.Context) ResourceOutput
}

func (*Resource) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOutput)
}

// ResourceArrayInput is an input type that accepts ResourceArray and ResourceArrayOutput values.
// You can construct a concrete instance of `ResourceArrayInput` via:
//
//	ResourceArray{ ResourceArgs{...} }
type ResourceArrayInput interface {
	pulumi.Input

	ToResourceArrayOutput() ResourceArrayOutput
	ToResourceArrayOutputWithContext(context.Context) ResourceArrayOutput
}

type ResourceArray []ResourceInput

func (ResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Resource)(nil)).Elem()
}

func (i ResourceArray) ToResourceArrayOutput() ResourceArrayOutput {
	return i.ToResourceArrayOutputWithContext(context.Background())
}

func (i ResourceArray) ToResourceArrayOutputWithContext(ctx context.Context) ResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceArrayOutput)
}

// ResourceMapInput is an input type that accepts ResourceMap and ResourceMapOutput values.
// You can construct a concrete instance of `ResourceMapInput` via:
//
//	ResourceMap{ "key": ResourceArgs{...} }
type ResourceMapInput interface {
	pulumi.Input

	ToResourceMapOutput() ResourceMapOutput
	ToResourceMapOutputWithContext(context.Context) ResourceMapOutput
}

type ResourceMap map[string]ResourceInput

func (ResourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Resource)(nil)).Elem()
}

func (i ResourceMap) ToResourceMapOutput() ResourceMapOutput {
	return i.ToResourceMapOutputWithContext(context.Background())
}

func (i ResourceMap) ToResourceMapOutputWithContext(ctx context.Context) ResourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMapOutput)
}

type ResourceOutput struct{ *pulumi.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

// A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
func (o ResourceOutput) Triggers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringMapOutput { return v.Triggers }).(pulumi.StringMapOutput)
}

type ResourceArrayOutput struct{ *pulumi.OutputState }

func (ResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Resource)(nil)).Elem()
}

func (o ResourceArrayOutput) ToResourceArrayOutput() ResourceArrayOutput {
	return o
}

func (o ResourceArrayOutput) ToResourceArrayOutputWithContext(ctx context.Context) ResourceArrayOutput {
	return o
}

func (o ResourceArrayOutput) Index(i pulumi.IntInput) ResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Resource {
		return vs[0].([]*Resource)[vs[1].(int)]
	}).(ResourceOutput)
}

type ResourceMapOutput struct{ *pulumi.OutputState }

func (ResourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Resource)(nil)).Elem()
}

func (o ResourceMapOutput) ToResourceMapOutput() ResourceMapOutput {
	return o
}

func (o ResourceMapOutput) ToResourceMapOutputWithContext(ctx context.Context) ResourceMapOutput {
	return o
}

func (o ResourceMapOutput) MapIndex(k pulumi.StringInput) ResourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Resource {
		return vs[0].(map[string]*Resource)[vs[1].(string)]
	}).(ResourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInput)(nil)).Elem(), &Resource{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceArrayInput)(nil)).Elem(), ResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceMapInput)(nil)).Elem(), ResourceMap{})
	pulumi.RegisterOutputType(ResourceOutput{})
	pulumi.RegisterOutputType(ResourceArrayOutput{})
	pulumi.RegisterOutputType(ResourceMapOutput{})
}
