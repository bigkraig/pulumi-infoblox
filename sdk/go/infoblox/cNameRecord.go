// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package infoblox

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type CNameRecord struct {
	s *pulumi.ResourceState
}

// NewCNameRecord registers a new resource with the given unique name, arguments, and options.
func NewCNameRecord(ctx *pulumi.Context,
	name string, args *CNameRecordArgs, opts ...pulumi.ResourceOpt) (*CNameRecord, error) {
	if args == nil || args.Canonical == nil {
		return nil, errors.New("missing required argument 'Canonical'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["canonical"] = nil
		inputs["comment"] = nil
		inputs["creator"] = nil
		inputs["ddnsPrincipal"] = nil
		inputs["ddnsProtected"] = nil
		inputs["disable"] = nil
		inputs["dnsName"] = nil
		inputs["forbidReclamation"] = nil
		inputs["name"] = nil
		inputs["reclaimable"] = nil
		inputs["sharedRecordGroup"] = nil
		inputs["ttl"] = nil
		inputs["useTtl"] = nil
		inputs["view"] = nil
	} else {
		inputs["canonical"] = args.Canonical
		inputs["comment"] = args.Comment
		inputs["creator"] = args.Creator
		inputs["ddnsPrincipal"] = args.DdnsPrincipal
		inputs["ddnsProtected"] = args.DdnsProtected
		inputs["disable"] = args.Disable
		inputs["dnsName"] = args.DnsName
		inputs["forbidReclamation"] = args.ForbidReclamation
		inputs["name"] = args.Name
		inputs["reclaimable"] = args.Reclaimable
		inputs["sharedRecordGroup"] = args.SharedRecordGroup
		inputs["ttl"] = args.Ttl
		inputs["useTtl"] = args.UseTtl
		inputs["view"] = args.View
	}
	inputs["creationTime"] = nil
	inputs["zone"] = nil
	s, err := ctx.RegisterResource("infoblox:index/cNameRecord:CNameRecord", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CNameRecord{s: s}, nil
}

// GetCNameRecord gets an existing CNameRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCNameRecord(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CNameRecordState, opts ...pulumi.ResourceOpt) (*CNameRecord, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["canonical"] = state.Canonical
		inputs["comment"] = state.Comment
		inputs["creationTime"] = state.CreationTime
		inputs["creator"] = state.Creator
		inputs["ddnsPrincipal"] = state.DdnsPrincipal
		inputs["ddnsProtected"] = state.DdnsProtected
		inputs["disable"] = state.Disable
		inputs["dnsName"] = state.DnsName
		inputs["forbidReclamation"] = state.ForbidReclamation
		inputs["name"] = state.Name
		inputs["reclaimable"] = state.Reclaimable
		inputs["sharedRecordGroup"] = state.SharedRecordGroup
		inputs["ttl"] = state.Ttl
		inputs["useTtl"] = state.UseTtl
		inputs["view"] = state.View
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("infoblox:index/cNameRecord:CNameRecord", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CNameRecord{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *CNameRecord) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *CNameRecord) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Canonical name in FQDN format
func (r *CNameRecord) Canonical() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["canonical"])
}

// Comment for the record; maximum 256 characters
func (r *CNameRecord) Comment() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["comment"])
}

// The time of the record creation in Epoch seconds format.
func (r *CNameRecord) CreationTime() *pulumi.Float64Output {
	return (*pulumi.Float64Output)(r.s.State["creationTime"])
}

// The record creator.Valid values:DYNAMIC,STATIC,SYSTEM.Defaults to STATIC
func (r *CNameRecord) Creator() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creator"])
}

// The GSS-TSIG principal that owns this record
func (r *CNameRecord) DdnsPrincipal() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ddnsPrincipal"])
}

// Determines if the DDNS updates for this record are allowed or not
func (r *CNameRecord) DdnsProtected() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["ddnsProtected"])
}

// Determines if the record is disabled or not. False means that the record is enabled.
func (r *CNameRecord) Disable() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disable"])
}

// The name for the CNAME record in punycode format. Values with leading or trailing white space are not valid for this
// field. Cannot be written nor updated.
func (r *CNameRecord) DnsName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dnsName"])
}

// Determines if the reclamation is allowed for the record or not.
func (r *CNameRecord) ForbidReclamation() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["forbidReclamation"])
}

// The name for a CNAME record in FQDN format
func (r *CNameRecord) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Determines if the record is reclaimable or not. Cannot be updated/written
func (r *CNameRecord) Reclaimable() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["reclaimable"])
}

// The name of the shared record group in which the record resides. This field exists only on db_objects if this record is
// a shared record. Cannot be updated/written
func (r *CNameRecord) SharedRecordGroup() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sharedRecordGroup"])
}

// The Time To Live assigned to CNAME
func (r *CNameRecord) Ttl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ttl"])
}

// Use flag for: ttl
func (r *CNameRecord) UseTtl() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["useTtl"])
}

// The name of the DNS View in which the record resides
func (r *CNameRecord) View() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["view"])
}

// DNS Zone for the record
func (r *CNameRecord) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering CNameRecord resources.
type CNameRecordState struct {
	// Canonical name in FQDN format
	Canonical interface{}
	// Comment for the record; maximum 256 characters
	Comment interface{}
	// The time of the record creation in Epoch seconds format.
	CreationTime interface{}
	// The record creator.Valid values:DYNAMIC,STATIC,SYSTEM.Defaults to STATIC
	Creator interface{}
	// The GSS-TSIG principal that owns this record
	DdnsPrincipal interface{}
	// Determines if the DDNS updates for this record are allowed or not
	DdnsProtected interface{}
	// Determines if the record is disabled or not. False means that the record is enabled.
	Disable interface{}
	// The name for the CNAME record in punycode format. Values with leading or trailing white space are not valid for this
	// field. Cannot be written nor updated.
	DnsName interface{}
	// Determines if the reclamation is allowed for the record or not.
	ForbidReclamation interface{}
	// The name for a CNAME record in FQDN format
	Name interface{}
	// Determines if the record is reclaimable or not. Cannot be updated/written
	Reclaimable interface{}
	// The name of the shared record group in which the record resides. This field exists only on db_objects if this record is
	// a shared record. Cannot be updated/written
	SharedRecordGroup interface{}
	// The Time To Live assigned to CNAME
	Ttl interface{}
	// Use flag for: ttl
	UseTtl interface{}
	// The name of the DNS View in which the record resides
	View interface{}
	// DNS Zone for the record
	Zone interface{}
}

// The set of arguments for constructing a CNameRecord resource.
type CNameRecordArgs struct {
	// Canonical name in FQDN format
	Canonical interface{}
	// Comment for the record; maximum 256 characters
	Comment interface{}
	// The record creator.Valid values:DYNAMIC,STATIC,SYSTEM.Defaults to STATIC
	Creator interface{}
	// The GSS-TSIG principal that owns this record
	DdnsPrincipal interface{}
	// Determines if the DDNS updates for this record are allowed or not
	DdnsProtected interface{}
	// Determines if the record is disabled or not. False means that the record is enabled.
	Disable interface{}
	// The name for the CNAME record in punycode format. Values with leading or trailing white space are not valid for this
	// field. Cannot be written nor updated.
	DnsName interface{}
	// Determines if the reclamation is allowed for the record or not.
	ForbidReclamation interface{}
	// The name for a CNAME record in FQDN format
	Name interface{}
	// Determines if the record is reclaimable or not. Cannot be updated/written
	Reclaimable interface{}
	// The name of the shared record group in which the record resides. This field exists only on db_objects if this record is
	// a shared record. Cannot be updated/written
	SharedRecordGroup interface{}
	// The Time To Live assigned to CNAME
	Ttl interface{}
	// Use flag for: ttl
	UseTtl interface{}
	// The name of the DNS View in which the record resides
	View interface{}
}
