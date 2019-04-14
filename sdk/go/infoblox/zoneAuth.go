// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package infoblox

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ZoneAuth struct {
	s *pulumi.ResourceState
}

// NewZoneAuth registers a new resource with the given unique name, arguments, and options.
func NewZoneAuth(ctx *pulumi.Context,
	name string, args *ZoneAuthArgs, opts ...pulumi.ResourceOpt) (*ZoneAuth, error) {
	if args == nil || args.Fqdn == nil {
		return nil, errors.New("missing required argument 'Fqdn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allowTransfers"] = nil
		inputs["allowUpdates"] = nil
		inputs["comment"] = nil
		inputs["copyXferToNotify"] = nil
		inputs["disable"] = nil
		inputs["dnsIntegrityEnable"] = nil
		inputs["dnsIntegrityMember"] = nil
		inputs["externalPrimaries"] = nil
		inputs["externalSecondaries"] = nil
		inputs["fqdn"] = nil
		inputs["gridPrimaries"] = nil
		inputs["gridPrimarySharedWithMsParentDelegation"] = nil
		inputs["gridSecondaries"] = nil
		inputs["locked"] = nil
		inputs["nsGroup"] = nil
		inputs["prefix"] = nil
		inputs["restartIfNeeded"] = nil
		inputs["scavengingSettings"] = nil
		inputs["soaDefaultTtl"] = nil
		inputs["soaExpire"] = nil
		inputs["soaNegativeTtl"] = nil
		inputs["soaRefresh"] = nil
		inputs["soaRetry"] = nil
		inputs["soaSerialNumber"] = nil
		inputs["useAllowTransfer"] = nil
		inputs["useCheckNamesPolicy"] = nil
		inputs["useCopyXferToNotify"] = nil
		inputs["useExternalPrimary"] = nil
		inputs["view"] = nil
		inputs["zoneFormat"] = nil
	} else {
		inputs["allowTransfers"] = args.AllowTransfers
		inputs["allowUpdates"] = args.AllowUpdates
		inputs["comment"] = args.Comment
		inputs["copyXferToNotify"] = args.CopyXferToNotify
		inputs["disable"] = args.Disable
		inputs["dnsIntegrityEnable"] = args.DnsIntegrityEnable
		inputs["dnsIntegrityMember"] = args.DnsIntegrityMember
		inputs["externalPrimaries"] = args.ExternalPrimaries
		inputs["externalSecondaries"] = args.ExternalSecondaries
		inputs["fqdn"] = args.Fqdn
		inputs["gridPrimaries"] = args.GridPrimaries
		inputs["gridPrimarySharedWithMsParentDelegation"] = args.GridPrimarySharedWithMsParentDelegation
		inputs["gridSecondaries"] = args.GridSecondaries
		inputs["locked"] = args.Locked
		inputs["nsGroup"] = args.NsGroup
		inputs["prefix"] = args.Prefix
		inputs["restartIfNeeded"] = args.RestartIfNeeded
		inputs["scavengingSettings"] = args.ScavengingSettings
		inputs["soaDefaultTtl"] = args.SoaDefaultTtl
		inputs["soaExpire"] = args.SoaExpire
		inputs["soaNegativeTtl"] = args.SoaNegativeTtl
		inputs["soaRefresh"] = args.SoaRefresh
		inputs["soaRetry"] = args.SoaRetry
		inputs["soaSerialNumber"] = args.SoaSerialNumber
		inputs["useAllowTransfer"] = args.UseAllowTransfer
		inputs["useCheckNamesPolicy"] = args.UseCheckNamesPolicy
		inputs["useCopyXferToNotify"] = args.UseCopyXferToNotify
		inputs["useExternalPrimary"] = args.UseExternalPrimary
		inputs["view"] = args.View
		inputs["zoneFormat"] = args.ZoneFormat
	}
	inputs["lockedBy"] = nil
	inputs["networkView"] = nil
	s, err := ctx.RegisterResource("infoblox:index/zoneAuth:ZoneAuth", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ZoneAuth{s: s}, nil
}

// GetZoneAuth gets an existing ZoneAuth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneAuth(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ZoneAuthState, opts ...pulumi.ResourceOpt) (*ZoneAuth, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allowTransfers"] = state.AllowTransfers
		inputs["allowUpdates"] = state.AllowUpdates
		inputs["comment"] = state.Comment
		inputs["copyXferToNotify"] = state.CopyXferToNotify
		inputs["disable"] = state.Disable
		inputs["dnsIntegrityEnable"] = state.DnsIntegrityEnable
		inputs["dnsIntegrityMember"] = state.DnsIntegrityMember
		inputs["externalPrimaries"] = state.ExternalPrimaries
		inputs["externalSecondaries"] = state.ExternalSecondaries
		inputs["fqdn"] = state.Fqdn
		inputs["gridPrimaries"] = state.GridPrimaries
		inputs["gridPrimarySharedWithMsParentDelegation"] = state.GridPrimarySharedWithMsParentDelegation
		inputs["gridSecondaries"] = state.GridSecondaries
		inputs["locked"] = state.Locked
		inputs["lockedBy"] = state.LockedBy
		inputs["networkView"] = state.NetworkView
		inputs["nsGroup"] = state.NsGroup
		inputs["prefix"] = state.Prefix
		inputs["restartIfNeeded"] = state.RestartIfNeeded
		inputs["scavengingSettings"] = state.ScavengingSettings
		inputs["soaDefaultTtl"] = state.SoaDefaultTtl
		inputs["soaExpire"] = state.SoaExpire
		inputs["soaNegativeTtl"] = state.SoaNegativeTtl
		inputs["soaRefresh"] = state.SoaRefresh
		inputs["soaRetry"] = state.SoaRetry
		inputs["soaSerialNumber"] = state.SoaSerialNumber
		inputs["useAllowTransfer"] = state.UseAllowTransfer
		inputs["useCheckNamesPolicy"] = state.UseCheckNamesPolicy
		inputs["useCopyXferToNotify"] = state.UseCopyXferToNotify
		inputs["useExternalPrimary"] = state.UseExternalPrimary
		inputs["view"] = state.View
		inputs["zoneFormat"] = state.ZoneFormat
	}
	s, err := ctx.ReadResource("infoblox:index/zoneAuth:ZoneAuth", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ZoneAuth{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ZoneAuth) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ZoneAuth) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Determines whether dynamic DNS updates are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks,
// and TSIG keys for the hosts.
func (r *ZoneAuth) AllowTransfers() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allowTransfers"])
}

// Determines whether dynamic DNS updates are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks,
// and TSIG keys for the hosts.
func (r *ZoneAuth) AllowUpdates() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allowUpdates"])
}

// Comment for the zone; maximum 256 characters
func (r *ZoneAuth) Comment() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["comment"])
}

// If this flag is set to True then copy allowed IPs from Allow Transfer to Also Notify.
func (r *ZoneAuth) CopyXferToNotify() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["copyXferToNotify"])
}

// Determines whether a zone is disabled or not
func (r *ZoneAuth) Disable() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disable"])
}

// If this is set to True, DNS integrity check is enabled for this zone
func (r *ZoneAuth) DnsIntegrityEnable() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["dnsIntegrityEnable"])
}

// The Grid member that performs DNS integrity checks for this zone
func (r *ZoneAuth) DnsIntegrityMember() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dnsIntegrityMember"])
}

// The primary preference list with Grid member names and/or External Server structs for this member.
func (r *ZoneAuth) ExternalPrimaries() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["externalPrimaries"])
}

// The primary preference list with Grid member names and/or External Server structs for this member.
func (r *ZoneAuth) ExternalSecondaries() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["externalSecondaries"])
}

// The name of this DNS zone. For a reverse zone, this is in “address/cidr” format
func (r *ZoneAuth) Fqdn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fqdn"])
}

// The grid primary servers for this zone.
func (r *ZoneAuth) GridPrimaries() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["gridPrimaries"])
}

// Determines if the server is duplicated with parent delegation.cannot be updated, nor written
func (r *ZoneAuth) GridPrimarySharedWithMsParentDelegation() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["gridPrimarySharedWithMsParentDelegation"])
}

// The grid primary servers for this zone.
func (r *ZoneAuth) GridSecondaries() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["gridSecondaries"])
}

// If you enable this flag, other administrators cannot make conflicting changes
func (r *ZoneAuth) Locked() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["locked"])
}

// The name of a superuser or the administrator who locked this zone (read-only)
func (r *ZoneAuth) LockedBy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lockedBy"])
}

// The name of the network view in which this zone resides (read-only)
func (r *ZoneAuth) NetworkView() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkView"])
}

// The name server group that serves DNS for this zone.
func (r *ZoneAuth) NsGroup() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["nsGroup"])
}

// The RFC2317 prefix value of this DNS zone
func (r *ZoneAuth) Prefix() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["prefix"])
}

// Restarts the member service. The default value is False. Not readable
func (r *ZoneAuth) RestartIfNeeded() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["restartIfNeeded"])
}

// The DNS scavenging settings object provides information about scavenging configuration e.g. conditions under which
// records can be scavenged, periodicity of scavenging operations.
func (r *ZoneAuth) ScavengingSettings() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["scavengingSettings"])
}

// The Time to Live (TTL) value of the SOA record of this zone
func (r *ZoneAuth) SoaDefaultTtl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["soaDefaultTtl"])
}

// This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about the
// zone because the zone data is too old to be useful. The default is one week.
func (r *ZoneAuth) SoaExpire() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["soaExpire"])
}

// The negative Time to Live (TTL)
func (r *ZoneAuth) SoaNegativeTtl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["soaNegativeTtl"])
}

// This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that
// its data is current, and retrieve fresh data if it is not
func (r *ZoneAuth) SoaRefresh() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["soaRefresh"])
}

// This indicates how long a secondary server must wait before attempting to recontact the primary server after a
// connection failure between the two servers occurs
func (r *ZoneAuth) SoaRetry() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["soaRetry"])
}

// The serial number in the SOA record incrementally changes every time the record is modified. The SOA serial number to be
// used in conjunction with set_soa_serial_number (read-only)
func (r *ZoneAuth) SoaSerialNumber() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["soaSerialNumber"])
}

// allow_transfer
func (r *ZoneAuth) UseAllowTransfer() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["useAllowTransfer"])
}

// Apply policy to dynamic updates and inbound zone transfers (This value applies only if the host name restriction policy
// is set to “Strict Hostname Checking”.)
func (r *ZoneAuth) UseCheckNamesPolicy() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["useCheckNamesPolicy"])
}

// Use flag for: copy_xfer_to_notify.
func (r *ZoneAuth) UseCopyXferToNotify() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["useCopyXferToNotify"])
}

// This flag controls whether the zone is using an external primary.
func (r *ZoneAuth) UseExternalPrimary() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["useExternalPrimary"])
}

// The name of the DNS view in which the zone resides
func (r *ZoneAuth) View() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["view"])
}

// Determines the format of this zone - API default FORWARD
func (r *ZoneAuth) ZoneFormat() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zoneFormat"])
}

// Input properties used for looking up and filtering ZoneAuth resources.
type ZoneAuthState struct {
	// Determines whether dynamic DNS updates are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks,
	// and TSIG keys for the hosts.
	AllowTransfers interface{}
	// Determines whether dynamic DNS updates are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks,
	// and TSIG keys for the hosts.
	AllowUpdates interface{}
	// Comment for the zone; maximum 256 characters
	Comment interface{}
	// If this flag is set to True then copy allowed IPs from Allow Transfer to Also Notify.
	CopyXferToNotify interface{}
	// Determines whether a zone is disabled or not
	Disable interface{}
	// If this is set to True, DNS integrity check is enabled for this zone
	DnsIntegrityEnable interface{}
	// The Grid member that performs DNS integrity checks for this zone
	DnsIntegrityMember interface{}
	// The primary preference list with Grid member names and/or External Server structs for this member.
	ExternalPrimaries interface{}
	// The primary preference list with Grid member names and/or External Server structs for this member.
	ExternalSecondaries interface{}
	// The name of this DNS zone. For a reverse zone, this is in “address/cidr” format
	Fqdn interface{}
	// The grid primary servers for this zone.
	GridPrimaries interface{}
	// Determines if the server is duplicated with parent delegation.cannot be updated, nor written
	GridPrimarySharedWithMsParentDelegation interface{}
	// The grid primary servers for this zone.
	GridSecondaries interface{}
	// If you enable this flag, other administrators cannot make conflicting changes
	Locked interface{}
	// The name of a superuser or the administrator who locked this zone (read-only)
	LockedBy interface{}
	// The name of the network view in which this zone resides (read-only)
	NetworkView interface{}
	// The name server group that serves DNS for this zone.
	NsGroup interface{}
	// The RFC2317 prefix value of this DNS zone
	Prefix interface{}
	// Restarts the member service. The default value is False. Not readable
	RestartIfNeeded interface{}
	// The DNS scavenging settings object provides information about scavenging configuration e.g. conditions under which
	// records can be scavenged, periodicity of scavenging operations.
	ScavengingSettings interface{}
	// The Time to Live (TTL) value of the SOA record of this zone
	SoaDefaultTtl interface{}
	// This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about
	// the zone because the zone data is too old to be useful. The default is one week.
	SoaExpire interface{}
	// The negative Time to Live (TTL)
	SoaNegativeTtl interface{}
	// This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that
	// its data is current, and retrieve fresh data if it is not
	SoaRefresh interface{}
	// This indicates how long a secondary server must wait before attempting to recontact the primary server after a
	// connection failure between the two servers occurs
	SoaRetry interface{}
	// The serial number in the SOA record incrementally changes every time the record is modified. The SOA serial number to
	// be used in conjunction with set_soa_serial_number (read-only)
	SoaSerialNumber interface{}
	// allow_transfer
	UseAllowTransfer interface{}
	// Apply policy to dynamic updates and inbound zone transfers (This value applies only if the host name restriction policy
	// is set to “Strict Hostname Checking”.)
	UseCheckNamesPolicy interface{}
	// Use flag for: copy_xfer_to_notify.
	UseCopyXferToNotify interface{}
	// This flag controls whether the zone is using an external primary.
	UseExternalPrimary interface{}
	// The name of the DNS view in which the zone resides
	View interface{}
	// Determines the format of this zone - API default FORWARD
	ZoneFormat interface{}
}

// The set of arguments for constructing a ZoneAuth resource.
type ZoneAuthArgs struct {
	// Determines whether dynamic DNS updates are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks,
	// and TSIG keys for the hosts.
	AllowTransfers interface{}
	// Determines whether dynamic DNS updates are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks,
	// and TSIG keys for the hosts.
	AllowUpdates interface{}
	// Comment for the zone; maximum 256 characters
	Comment interface{}
	// If this flag is set to True then copy allowed IPs from Allow Transfer to Also Notify.
	CopyXferToNotify interface{}
	// Determines whether a zone is disabled or not
	Disable interface{}
	// If this is set to True, DNS integrity check is enabled for this zone
	DnsIntegrityEnable interface{}
	// The Grid member that performs DNS integrity checks for this zone
	DnsIntegrityMember interface{}
	// The primary preference list with Grid member names and/or External Server structs for this member.
	ExternalPrimaries interface{}
	// The primary preference list with Grid member names and/or External Server structs for this member.
	ExternalSecondaries interface{}
	// The name of this DNS zone. For a reverse zone, this is in “address/cidr” format
	Fqdn interface{}
	// The grid primary servers for this zone.
	GridPrimaries interface{}
	// Determines if the server is duplicated with parent delegation.cannot be updated, nor written
	GridPrimarySharedWithMsParentDelegation interface{}
	// The grid primary servers for this zone.
	GridSecondaries interface{}
	// If you enable this flag, other administrators cannot make conflicting changes
	Locked interface{}
	// The name server group that serves DNS for this zone.
	NsGroup interface{}
	// The RFC2317 prefix value of this DNS zone
	Prefix interface{}
	// Restarts the member service. The default value is False. Not readable
	RestartIfNeeded interface{}
	// The DNS scavenging settings object provides information about scavenging configuration e.g. conditions under which
	// records can be scavenged, periodicity of scavenging operations.
	ScavengingSettings interface{}
	// The Time to Live (TTL) value of the SOA record of this zone
	SoaDefaultTtl interface{}
	// This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about
	// the zone because the zone data is too old to be useful. The default is one week.
	SoaExpire interface{}
	// The negative Time to Live (TTL)
	SoaNegativeTtl interface{}
	// This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that
	// its data is current, and retrieve fresh data if it is not
	SoaRefresh interface{}
	// This indicates how long a secondary server must wait before attempting to recontact the primary server after a
	// connection failure between the two servers occurs
	SoaRetry interface{}
	// The serial number in the SOA record incrementally changes every time the record is modified. The SOA serial number to
	// be used in conjunction with set_soa_serial_number (read-only)
	SoaSerialNumber interface{}
	// allow_transfer
	UseAllowTransfer interface{}
	// Apply policy to dynamic updates and inbound zone transfers (This value applies only if the host name restriction policy
	// is set to “Strict Hostname Checking”.)
	UseCheckNamesPolicy interface{}
	// Use flag for: copy_xfer_to_notify.
	UseCopyXferToNotify interface{}
	// This flag controls whether the zone is using an external primary.
	UseExternalPrimary interface{}
	// The name of the DNS view in which the zone resides
	View interface{}
	// Determines the format of this zone - API default FORWARD
	ZoneFormat interface{}
}
