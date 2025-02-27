// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mod1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A test for namespaces (mod 1)
type Typ struct {
	Val *string `pulumi:"val"`
}

// Defaults sets the appropriate defaults for Typ
func (val *Typ) Defaults() *Typ {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Val == nil {
		if d := getEnvOrDefault("mod1", nil, "PULUMI_EXAMPLE_MOD1_DEFAULT"); d != nil {
			val_ := d.(string)
			tmp.Val = &val_
		}
	}
	return &tmp
}

// TypInput is an input type that accepts TypArgs and TypOutput values.
// You can construct a concrete instance of `TypInput` via:
//
//	TypArgs{...}
type TypInput interface {
	pulumi.Input

	ToTypOutput() TypOutput
	ToTypOutputWithContext(context.Context) TypOutput
}

// A test for namespaces (mod 1)
type TypArgs struct {
	Val pulumi.StringPtrInput `pulumi:"val"`
}

// Defaults sets the appropriate defaults for TypArgs
func (val *TypArgs) Defaults() *TypArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Val == nil {
		if d := getEnvOrDefault("mod1", nil, "PULUMI_EXAMPLE_MOD1_DEFAULT"); d != nil {
			tmp.Val = pulumi.StringPtr(d.(string))
		}
	}
	return &tmp
}
func (TypArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Typ)(nil)).Elem()
}

func (i TypArgs) ToTypOutput() TypOutput {
	return i.ToTypOutputWithContext(context.Background())
}

func (i TypArgs) ToTypOutputWithContext(ctx context.Context) TypOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TypOutput)
}

func (i TypArgs) ToTypPtrOutput() TypPtrOutput {
	return i.ToTypPtrOutputWithContext(context.Background())
}

func (i TypArgs) ToTypPtrOutputWithContext(ctx context.Context) TypPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TypOutput).ToTypPtrOutputWithContext(ctx)
}

// TypPtrInput is an input type that accepts TypArgs, TypPtr and TypPtrOutput values.
// You can construct a concrete instance of `TypPtrInput` via:
//
//	        TypArgs{...}
//
//	or:
//
//	        nil
type TypPtrInput interface {
	pulumi.Input

	ToTypPtrOutput() TypPtrOutput
	ToTypPtrOutputWithContext(context.Context) TypPtrOutput
}

type typPtrType TypArgs

func TypPtr(v *TypArgs) TypPtrInput {
	return (*typPtrType)(v)
}

func (*typPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Typ)(nil)).Elem()
}

func (i *typPtrType) ToTypPtrOutput() TypPtrOutput {
	return i.ToTypPtrOutputWithContext(context.Background())
}

func (i *typPtrType) ToTypPtrOutputWithContext(ctx context.Context) TypPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TypPtrOutput)
}

// A test for namespaces (mod 1)
type TypOutput struct{ *pulumi.OutputState }

func (TypOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Typ)(nil)).Elem()
}

func (o TypOutput) ToTypOutput() TypOutput {
	return o
}

func (o TypOutput) ToTypOutputWithContext(ctx context.Context) TypOutput {
	return o
}

func (o TypOutput) ToTypPtrOutput() TypPtrOutput {
	return o.ToTypPtrOutputWithContext(context.Background())
}

func (o TypOutput) ToTypPtrOutputWithContext(ctx context.Context) TypPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Typ) *Typ {
		return &v
	}).(TypPtrOutput)
}

func (o TypOutput) Val() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Typ) *string { return v.Val }).(pulumi.StringPtrOutput)
}

type TypPtrOutput struct{ *pulumi.OutputState }

func (TypPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Typ)(nil)).Elem()
}

func (o TypPtrOutput) ToTypPtrOutput() TypPtrOutput {
	return o
}

func (o TypPtrOutput) ToTypPtrOutputWithContext(ctx context.Context) TypPtrOutput {
	return o
}

func (o TypPtrOutput) Elem() TypOutput {
	return o.ApplyT(func(v *Typ) Typ {
		if v != nil {
			return *v
		}
		var ret Typ
		return ret
	}).(TypOutput)
}

func (o TypPtrOutput) Val() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Typ) *string {
		if v == nil {
			return nil
		}
		return v.Val
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TypInput)(nil)).Elem(), TypArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TypPtrInput)(nil)).Elem(), TypArgs{})
	pulumi.RegisterOutputType(TypOutput{})
	pulumi.RegisterOutputType(TypPtrOutput{})
}
