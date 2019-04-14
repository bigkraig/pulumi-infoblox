// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package infoblox

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type NSGroupForward struct {
	s *pulumi.ResourceState
}

// NewNSGroupForward registers a new resource with the given unique name, arguments, and options.
func NewNSGroupForward(ctx *pulumi.Context,
	name string, args *NSGroupForwardArgs, opts ...pulumi.ResourceOpt) (*NSGroupForward, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["comment"] = nil
		inputs["forwardingServers"] = nil
		inputs["name"] = nil
	} else {
		inputs["comment"] = args.Comment
		inputs["forwardingServers"] = args.ForwardingServers
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("infoblox:index/nSGroupForward:NSGroupForward", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NSGroupForward{s: s}, nil
}

// GetNSGroupForward gets an existing NSGroupForward resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNSGroupForward(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NSGroupForwardState, opts ...pulumi.ResourceOpt) (*NSGroupForward, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["comment"] = state.Comment
		inputs["forwardingServers"] = state.ForwardingServers
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("infoblox:index/nSGroupForward:NSGroupForward", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NSGroupForward{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NSGroupForward) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NSGroupForward) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Comment field
func (r *NSGroupForward) Comment() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["comment"])
}

// The primary preference list with Grid member names and/or External Server structs for this member.
func (r *NSGroupForward) ForwardingServers() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["forwardingServers"])
}

// The name of the name server group
func (r *NSGroupForward) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering NSGroupForward resources.
type NSGroupForwardState struct {
	// Comment field
	Comment interface{}
	// The primary preference list with Grid member names and/or External Server structs for this member.
	ForwardingServers interface{}
	// The name of the name server group
	Name interface{}
}

// The set of arguments for constructing a NSGroupForward resource.
type NSGroupForwardArgs struct {
	// Comment field
	Comment interface{}
	// The primary preference list with Grid member names and/or External Server structs for this member.
	ForwardingServers interface{}
	// The name of the name server group
	Name interface{}
}