// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package infoblox

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type MXRecord struct {
	s *pulumi.ResourceState
}

// NewMXRecord registers a new resource with the given unique name, arguments, and options.
func NewMXRecord(ctx *pulumi.Context,
	name string, args *MXRecordArgs, opts ...pulumi.ResourceOpt) (*MXRecord, error) {
	if args == nil || args.MailExchanger == nil {
		return nil, errors.New("missing required argument 'MailExchanger'")
	}
	if args == nil || args.Preference == nil {
		return nil, errors.New("missing required argument 'Preference'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["comment"] = nil
		inputs["ddnsPrincipal"] = nil
		inputs["ddnsProtected"] = nil
		inputs["disable"] = nil
		inputs["mailExchanger"] = nil
		inputs["name"] = nil
		inputs["preference"] = nil
		inputs["ttl"] = nil
		inputs["useTtl"] = nil
		inputs["view"] = nil
	} else {
		inputs["comment"] = args.Comment
		inputs["ddnsPrincipal"] = args.DdnsPrincipal
		inputs["ddnsProtected"] = args.DdnsProtected
		inputs["disable"] = args.Disable
		inputs["mailExchanger"] = args.MailExchanger
		inputs["name"] = args.Name
		inputs["preference"] = args.Preference
		inputs["ttl"] = args.Ttl
		inputs["useTtl"] = args.UseTtl
		inputs["view"] = args.View
	}
	s, err := ctx.RegisterResource("infoblox:index/mXRecord:MXRecord", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MXRecord{s: s}, nil
}

// GetMXRecord gets an existing MXRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMXRecord(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MXRecordState, opts ...pulumi.ResourceOpt) (*MXRecord, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["comment"] = state.Comment
		inputs["ddnsPrincipal"] = state.DdnsPrincipal
		inputs["ddnsProtected"] = state.DdnsProtected
		inputs["disable"] = state.Disable
		inputs["mailExchanger"] = state.MailExchanger
		inputs["name"] = state.Name
		inputs["preference"] = state.Preference
		inputs["ttl"] = state.Ttl
		inputs["useTtl"] = state.UseTtl
		inputs["view"] = state.View
	}
	s, err := ctx.ReadResource("infoblox:index/mXRecord:MXRecord", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MXRecord{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *MXRecord) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *MXRecord) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A comment on the record
func (r *MXRecord) Comment() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["comment"])
}

// The GSS-TSIG principal that owns this record
func (r *MXRecord) DdnsPrincipal() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ddnsPrincipal"])
}

// Determines if the DDNS updates for this record are allowed or not
func (r *MXRecord) DdnsProtected() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["ddnsProtected"])
}

// Is the record disabled
func (r *MXRecord) Disable() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disable"])
}

// Mail exchanger name in FQDN format. This value can be in unicode format
func (r *MXRecord) MailExchanger() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["mailExchanger"])
}

// Name of the zone the MX record refers to
func (r *MXRecord) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Preference value, 0 to 65535 (inclusive) in 32-bit unsigned integer format
func (r *MXRecord) Preference() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["preference"])
}

// The Time To Live value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the
// record is valid
func (r *MXRecord) Ttl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ttl"])
}

// Use flag for: ttl
func (r *MXRecord) UseTtl() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["useTtl"])
}

// The name of the DNS view in which the record resides.
func (r *MXRecord) View() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["view"])
}

// Input properties used for looking up and filtering MXRecord resources.
type MXRecordState struct {
	// A comment on the record
	Comment interface{}
	// The GSS-TSIG principal that owns this record
	DdnsPrincipal interface{}
	// Determines if the DDNS updates for this record are allowed or not
	DdnsProtected interface{}
	// Is the record disabled
	Disable interface{}
	// Mail exchanger name in FQDN format. This value can be in unicode format
	MailExchanger interface{}
	// Name of the zone the MX record refers to
	Name interface{}
	// Preference value, 0 to 65535 (inclusive) in 32-bit unsigned integer format
	Preference interface{}
	// The Time To Live value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the
	// record is valid
	Ttl interface{}
	// Use flag for: ttl
	UseTtl interface{}
	// The name of the DNS view in which the record resides.
	View interface{}
}

// The set of arguments for constructing a MXRecord resource.
type MXRecordArgs struct {
	// A comment on the record
	Comment interface{}
	// The GSS-TSIG principal that owns this record
	DdnsPrincipal interface{}
	// Determines if the DDNS updates for this record are allowed or not
	DdnsProtected interface{}
	// Is the record disabled
	Disable interface{}
	// Mail exchanger name in FQDN format. This value can be in unicode format
	MailExchanger interface{}
	// Name of the zone the MX record refers to
	Name interface{}
	// Preference value, 0 to 65535 (inclusive) in 32-bit unsigned integer format
	Preference interface{}
	// The Time To Live value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the
	// record is valid
	Ttl interface{}
	// Use flag for: ttl
	UseTtl interface{}
	// The name of the DNS view in which the record resides.
	View interface{}
}
