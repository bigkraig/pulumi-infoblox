// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package infoblox

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type DHCPRange struct {
	s *pulumi.ResourceState
}

// NewDHCPRange registers a new resource with the given unique name, arguments, and options.
func NewDHCPRange(ctx *pulumi.Context,
	name string, args *DHCPRangeArgs, opts ...pulumi.ResourceOpt) (*DHCPRange, error) {
	if args == nil || args.EndAddr == nil {
		return nil, errors.New("missing required argument 'EndAddr'")
	}
	if args == nil || args.StartAddr == nil {
		return nil, errors.New("missing required argument 'StartAddr'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["comment"] = nil
		inputs["endAddr"] = nil
		inputs["members"] = nil
		inputs["name"] = nil
		inputs["network"] = nil
		inputs["networkView"] = nil
		inputs["restartIfNeeded"] = nil
		inputs["serverAssociationType"] = nil
		inputs["startAddr"] = nil
	} else {
		inputs["comment"] = args.Comment
		inputs["endAddr"] = args.EndAddr
		inputs["members"] = args.Members
		inputs["name"] = args.Name
		inputs["network"] = args.Network
		inputs["networkView"] = args.NetworkView
		inputs["restartIfNeeded"] = args.RestartIfNeeded
		inputs["serverAssociationType"] = args.ServerAssociationType
		inputs["startAddr"] = args.StartAddr
	}
	s, err := ctx.RegisterResource("infoblox:index/dHCPRange:DHCPRange", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DHCPRange{s: s}, nil
}

// GetDHCPRange gets an existing DHCPRange resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDHCPRange(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DHCPRangeState, opts ...pulumi.ResourceOpt) (*DHCPRange, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["comment"] = state.Comment
		inputs["endAddr"] = state.EndAddr
		inputs["members"] = state.Members
		inputs["name"] = state.Name
		inputs["network"] = state.Network
		inputs["networkView"] = state.NetworkView
		inputs["restartIfNeeded"] = state.RestartIfNeeded
		inputs["serverAssociationType"] = state.ServerAssociationType
		inputs["startAddr"] = state.StartAddr
	}
	s, err := ctx.ReadResource("infoblox:index/dHCPRange:DHCPRange", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DHCPRange{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DHCPRange) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DHCPRange) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *DHCPRange) Comment() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["comment"])
}

// The IPv4 Address end address of the range.
func (r *DHCPRange) EndAddr() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endAddr"])
}

// Infoblox DHCP member that serves this range
func (r *DHCPRange) Members() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["members"])
}

// This field contains the name of the Microsoft scope.
func (r *DHCPRange) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The network to which this range belongs, in IPv4 Address/CIDR format.
func (r *DHCPRange) Network() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["network"])
}

// The name of the network view in which this range resides.
func (r *DHCPRange) NetworkView() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkView"])
}

// Restarts any services if required by this change. Default: true.
func (r *DHCPRange) RestartIfNeeded() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["restartIfNeeded"])
}

// Must be set to 'MEMBER' if member is specified
func (r *DHCPRange) ServerAssociationType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serverAssociationType"])
}

// The IPv4 Address starting address of the range.
func (r *DHCPRange) StartAddr() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["startAddr"])
}

// Input properties used for looking up and filtering DHCPRange resources.
type DHCPRangeState struct {
	Comment interface{}
	// The IPv4 Address end address of the range.
	EndAddr interface{}
	// Infoblox DHCP member that serves this range
	Members interface{}
	// This field contains the name of the Microsoft scope.
	Name interface{}
	// The network to which this range belongs, in IPv4 Address/CIDR format.
	Network interface{}
	// The name of the network view in which this range resides.
	NetworkView interface{}
	// Restarts any services if required by this change. Default: true.
	RestartIfNeeded interface{}
	// Must be set to 'MEMBER' if member is specified
	ServerAssociationType interface{}
	// The IPv4 Address starting address of the range.
	StartAddr interface{}
}

// The set of arguments for constructing a DHCPRange resource.
type DHCPRangeArgs struct {
	Comment interface{}
	// The IPv4 Address end address of the range.
	EndAddr interface{}
	// Infoblox DHCP member that serves this range
	Members interface{}
	// This field contains the name of the Microsoft scope.
	Name interface{}
	// The network to which this range belongs, in IPv4 Address/CIDR format.
	Network interface{}
	// The name of the network view in which this range resides.
	NetworkView interface{}
	// Restarts any services if required by this change. Default: true.
	RestartIfNeeded interface{}
	// Must be set to 'MEMBER' if member is specified
	ServerAssociationType interface{}
	// The IPv4 Address starting address of the range.
	StartAddr interface{}
}
