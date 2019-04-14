// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package infoblox

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type NSRecord struct {
	s *pulumi.ResourceState
}

// NewNSRecord registers a new resource with the given unique name, arguments, and options.
func NewNSRecord(ctx *pulumi.Context,
	name string, args *NSRecordArgs, opts ...pulumi.ResourceOpt) (*NSRecord, error) {
	if args == nil || args.Addresses == nil {
		return nil, errors.New("missing required argument 'Addresses'")
	}
	if args == nil || args.Nameserver == nil {
		return nil, errors.New("missing required argument 'Nameserver'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["addresses"] = nil
		inputs["creator"] = nil
		inputs["dnsName"] = nil
		inputs["msDelegationName"] = nil
		inputs["name"] = nil
		inputs["nameserver"] = nil
		inputs["policy"] = nil
		inputs["view"] = nil
	} else {
		inputs["addresses"] = args.Addresses
		inputs["creator"] = args.Creator
		inputs["dnsName"] = args.DnsName
		inputs["msDelegationName"] = args.MsDelegationName
		inputs["name"] = args.Name
		inputs["nameserver"] = args.Nameserver
		inputs["policy"] = args.Policy
		inputs["view"] = args.View
	}
	inputs["zone"] = nil
	s, err := ctx.RegisterResource("infoblox:index/nSRecord:NSRecord", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NSRecord{s: s}, nil
}

// GetNSRecord gets an existing NSRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNSRecord(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NSRecordState, opts ...pulumi.ResourceOpt) (*NSRecord, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["addresses"] = state.Addresses
		inputs["creator"] = state.Creator
		inputs["dnsName"] = state.DnsName
		inputs["msDelegationName"] = state.MsDelegationName
		inputs["name"] = state.Name
		inputs["nameserver"] = state.Nameserver
		inputs["policy"] = state.Policy
		inputs["view"] = state.View
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("infoblox:index/nSRecord:NSRecord", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NSRecord{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NSRecord) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NSRecord) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The list of zone name servers
func (r *NSRecord) Addresses() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["addresses"])
}

// The record creator.Valid values:DYNAMIC,STATIC,SYSTEM.Defaults to STATIC
func (r *NSRecord) Creator() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creator"])
}

// The name for the NS record in punycode format. Values with leading or trailing white space are not valid for this field.
// Cannot be written nor updated.
func (r *NSRecord) DnsName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dnsName"])
}

// The MS delegation point name
func (r *NSRecord) MsDelegationName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["msDelegationName"])
}

// The name of the ns record where the record should reside. Cannot be updated
func (r *NSRecord) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The FQDN of the name server
func (r *NSRecord) Nameserver() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["nameserver"])
}

func (r *NSRecord) Policy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policy"])
}

// The name of the DNS view in which the record resides
func (r *NSRecord) View() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["view"])
}

// DNS Zone for the record
func (r *NSRecord) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering NSRecord resources.
type NSRecordState struct {
	// The list of zone name servers
	Addresses interface{}
	// The record creator.Valid values:DYNAMIC,STATIC,SYSTEM.Defaults to STATIC
	Creator interface{}
	// The name for the NS record in punycode format. Values with leading or trailing white space are not valid for this
	// field. Cannot be written nor updated.
	DnsName interface{}
	// The MS delegation point name
	MsDelegationName interface{}
	// The name of the ns record where the record should reside. Cannot be updated
	Name interface{}
	// The FQDN of the name server
	Nameserver interface{}
	Policy interface{}
	// The name of the DNS view in which the record resides
	View interface{}
	// DNS Zone for the record
	Zone interface{}
}

// The set of arguments for constructing a NSRecord resource.
type NSRecordArgs struct {
	// The list of zone name servers
	Addresses interface{}
	// The record creator.Valid values:DYNAMIC,STATIC,SYSTEM.Defaults to STATIC
	Creator interface{}
	// The name for the NS record in punycode format. Values with leading or trailing white space are not valid for this
	// field. Cannot be written nor updated.
	DnsName interface{}
	// The MS delegation point name
	MsDelegationName interface{}
	// The name of the ns record where the record should reside. Cannot be updated
	Name interface{}
	// The FQDN of the name server
	Nameserver interface{}
	Policy interface{}
	// The name of the DNS view in which the record resides
	View interface{}
}
