// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package my8664

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConditionalAccessPolicy struct {
	pulumi.CustomResourceState

	Conditions ConditionalAccessPolicyConditionsOutput `pulumi:"conditions"`
}

// NewConditionalAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewConditionalAccessPolicy(ctx *pulumi.Context,
	name string, args *ConditionalAccessPolicyArgs, opts ...pulumi.ResourceOption) (*ConditionalAccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Conditions == nil {
		return nil, errors.New("invalid value for required argument 'Conditions'")
	}
	var resource ConditionalAccessPolicy
	err := ctx.RegisterResource("my8664::ConditionalAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConditionalAccessPolicy gets an existing ConditionalAccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConditionalAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConditionalAccessPolicyState, opts ...pulumi.ResourceOption) (*ConditionalAccessPolicy, error) {
	var resource ConditionalAccessPolicy
	err := ctx.ReadResource("my8664::ConditionalAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConditionalAccessPolicy resources.
type conditionalAccessPolicyState struct {
	Conditions *ConditionalAccessPolicyConditions `pulumi:"conditions"`
}

type ConditionalAccessPolicyState struct {
	Conditions ConditionalAccessPolicyConditionsPtrInput
}

func (ConditionalAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*conditionalAccessPolicyState)(nil)).Elem()
}

type conditionalAccessPolicyArgs struct {
	Conditions ConditionalAccessPolicyConditions `pulumi:"conditions"`
}

// The set of arguments for constructing a ConditionalAccessPolicy resource.
type ConditionalAccessPolicyArgs struct {
	Conditions ConditionalAccessPolicyConditionsInput
}

func (ConditionalAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*conditionalAccessPolicyArgs)(nil)).Elem()
}

type ConditionalAccessPolicyInput interface {
	pulumi.Input

	ToConditionalAccessPolicyOutput() ConditionalAccessPolicyOutput
	ToConditionalAccessPolicyOutputWithContext(ctx context.Context) ConditionalAccessPolicyOutput
}

func (*ConditionalAccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionalAccessPolicy)(nil)).Elem()
}

func (i *ConditionalAccessPolicy) ToConditionalAccessPolicyOutput() ConditionalAccessPolicyOutput {
	return i.ToConditionalAccessPolicyOutputWithContext(context.Background())
}

func (i *ConditionalAccessPolicy) ToConditionalAccessPolicyOutputWithContext(ctx context.Context) ConditionalAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionalAccessPolicyOutput)
}

type ConditionalAccessPolicyOutput struct{ *pulumi.OutputState }

func (ConditionalAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionalAccessPolicy)(nil)).Elem()
}

func (o ConditionalAccessPolicyOutput) ToConditionalAccessPolicyOutput() ConditionalAccessPolicyOutput {
	return o
}

func (o ConditionalAccessPolicyOutput) ToConditionalAccessPolicyOutputWithContext(ctx context.Context) ConditionalAccessPolicyOutput {
	return o
}

func (o ConditionalAccessPolicyOutput) Conditions() ConditionalAccessPolicyConditionsOutput {
	return o.ApplyT(func(v *ConditionalAccessPolicy) ConditionalAccessPolicyConditionsOutput { return v.Conditions }).(ConditionalAccessPolicyConditionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionalAccessPolicyInput)(nil)).Elem(), &ConditionalAccessPolicy{})
	pulumi.RegisterOutputType(ConditionalAccessPolicyOutput{})
}
