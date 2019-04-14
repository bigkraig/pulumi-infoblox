// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package infoblox

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type TXTRecord struct {
	s *pulumi.ResourceState
}

// NewTXTRecord registers a new resource with the given unique name, arguments, and options.
func NewTXTRecord(ctx *pulumi.Context,
	name string, args *TXTRecordArgs, opts ...pulumi.ResourceOpt) (*TXTRecord, error) {
	if args == nil || args.Text == nil {
		return nil, errors.New("missing required argument 'Text'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["comment"] = nil
		inputs["name"] = nil
		inputs["text"] = nil
		inputs["ttl"] = nil
		inputs["useTtl"] = nil
		inputs["view"] = nil
		inputs["zone"] = nil
	} else {
		inputs["comment"] = args.Comment
		inputs["name"] = args.Name
		inputs["text"] = args.Text
		inputs["ttl"] = args.Ttl
		inputs["useTtl"] = args.UseTtl
		inputs["view"] = args.View
		inputs["zone"] = args.Zone
	}
	s, err := ctx.RegisterResource("infoblox:index/tXTRecord:TXTRecord", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TXTRecord{s: s}, nil
}

// GetTXTRecord gets an existing TXTRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTXTRecord(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TXTRecordState, opts ...pulumi.ResourceOpt) (*TXTRecord, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["comment"] = state.Comment
		inputs["name"] = state.Name
		inputs["text"] = state.Text
		inputs["ttl"] = state.Ttl
		inputs["useTtl"] = state.UseTtl
		inputs["view"] = state.View
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("infoblox:index/tXTRecord:TXTRecord", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TXTRecord{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TXTRecord) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TXTRecord) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *TXTRecord) Comment() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["comment"])
}

func (r *TXTRecord) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *TXTRecord) Text() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["text"])
}

func (r *TXTRecord) Ttl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ttl"])
}

func (r *TXTRecord) UseTtl() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["useTtl"])
}

func (r *TXTRecord) View() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["view"])
}

func (r *TXTRecord) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering TXTRecord resources.
type TXTRecordState struct {
	Comment interface{}
	Name interface{}
	Text interface{}
	Ttl interface{}
	UseTtl interface{}
	View interface{}
	Zone interface{}
}

// The set of arguments for constructing a TXTRecord resource.
type TXTRecordArgs struct {
	Comment interface{}
	Name interface{}
	Text interface{}
	Ttl interface{}
	UseTtl interface{}
	View interface{}
	Zone interface{}
}