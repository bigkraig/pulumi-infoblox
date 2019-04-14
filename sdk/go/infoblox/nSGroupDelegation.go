// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package infoblox

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type NSGroupDelegation struct {
	s *pulumi.ResourceState
}

// NewNSGroupDelegation registers a new resource with the given unique name, arguments, and options.
func NewNSGroupDelegation(ctx *pulumi.Context,
	name string, args *NSGroupDelegationArgs, opts ...pulumi.ResourceOpt) (*NSGroupDelegation, error) {
	if args == nil || args.DelegateTos == nil {
		return nil, errors.New("missing required argument 'DelegateTos'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["comment"] = nil
		inputs["delegateTos"] = nil
		inputs["name"] = nil
	} else {
		inputs["comment"] = args.Comment
		inputs["delegateTos"] = args.DelegateTos
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("infoblox:index/nSGroupDelegation:NSGroupDelegation", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NSGroupDelegation{s: s}, nil
}

// GetNSGroupDelegation gets an existing NSGroupDelegation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNSGroupDelegation(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NSGroupDelegationState, opts ...pulumi.ResourceOpt) (*NSGroupDelegation, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["comment"] = state.Comment
		inputs["delegateTos"] = state.DelegateTos
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("infoblox:index/nSGroupDelegation:NSGroupDelegation", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NSGroupDelegation{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NSGroupDelegation) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NSGroupDelegation) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Comment field
func (r *NSGroupDelegation) Comment() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["comment"])
}

// The primary preference set with Grid member names and/or External Server structs for this member.
func (r *NSGroupDelegation) DelegateTos() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["delegateTos"])
}

// The name of the name server group
func (r *NSGroupDelegation) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering NSGroupDelegation resources.
type NSGroupDelegationState struct {
	// Comment field
	Comment interface{}
	// The primary preference set with Grid member names and/or External Server structs for this member.
	DelegateTos interface{}
	// The name of the name server group
	Name interface{}
}

// The set of arguments for constructing a NSGroupDelegation resource.
type NSGroupDelegationArgs struct {
	// Comment field
	Comment interface{}
	// The primary preference set with Grid member names and/or External Server structs for this member.
	DelegateTos interface{}
	// The name of the name server group
	Name interface{}
}