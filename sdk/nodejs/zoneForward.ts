// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ZoneForward extends pulumi.CustomResource {
    /**
     * Get an existing ZoneForward resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneForwardState, opts?: pulumi.CustomResourceOptions): ZoneForward {
        return new ZoneForward(name, <any>state, { ...opts, id: id });
    }

    /**
     * The IPv4 Address or IPv6 Address of the server that is serving this zone. Not writable
     */
    public /*out*/ readonly address: pulumi.Output<string>;
    /**
     * Comment for the zone; maximum 256 characters
     */
    public readonly comment: pulumi.Output<string | undefined>;
    /**
     * Determines whether a zone is disabled or not
     */
    public readonly disable: pulumi.Output<boolean | undefined>;
    /**
     * The displayed name of the DNS zone.Not writable
     */
    public readonly displayDomain: pulumi.Output<string>;
    /**
     * The name of this DNS zone in punycode format.For a reverse zone, this is in “address/cidr” format.For other
     * zones, this is in FQDN format in punycode format.Cannot be updated
     */
    public readonly dnsFqdn: pulumi.Output<string>;
    /**
     * The primary preference list with Grid member names and/or External Server structs for this member.
     */
    public readonly forwardTos: pulumi.Output<{ address: string, name: string, sharedWithMsParentDelegation: boolean, stealth?: boolean, tsigKey?: string, tsigKeyAlg?: string, tsigKeyName?: string, useTsigKeyName?: boolean }[]>;
    /**
     * Determines if the appliance sends queries to forwarders only and not to other internal or Internet root servers
     */
    public readonly forwardersOnly: pulumi.Output<boolean | undefined>;
    /**
     * The primary preference list with Grid member names and/or External Server structs for this member.
     */
    public readonly forwardingServers: pulumi.Output<{ forwardTos?: { address: string, name: string, sharedWithMsParentDelegation: boolean, stealth?: boolean, tsigKey?: string, tsigKeyAlg?: string, tsigKeyName?: string, useTsigKeyName?: boolean }[], forwardersOnly?: boolean, name: string, useOverrideForwarders?: boolean }[] | undefined>;
    /**
     * The name of this DNS zone. For a reverse zone, this is in “address/cidr” format
     */
    public readonly fqdn: pulumi.Output<string | undefined>;
    /**
     * If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes
     * only. The zone will continue to serve DNS data even when it is locked.The default value is False.
     */
    public readonly locked: pulumi.Output<boolean | undefined>;
    /**
     * The name of a superuser or the administrator who locked this zone.Not writable
     */
    public readonly lockedBy: pulumi.Output<string>;
    /**
     * IPv4 Netmask or IPv6 prefix for this zone.Not Writable
     */
    public readonly maskPrefix: pulumi.Output<string>;
    /**
     * A forwarding member name server group. Values with leading or trailing white space are not valid for this field. The
     * default value is undefined.
     */
    public readonly nsGroup: pulumi.Output<string | undefined>;
    /**
     * The parent zone of this zone. Note that when searching for reverse zones, the “in-addr.arpa” notation should be
     * used. Not writable.
     */
    public readonly parent: pulumi.Output<string>;
    /**
     * The RFC2317 prefix value of this DNS zone.
     */
    public readonly prefix: pulumi.Output<string | undefined>;
    /**
     * This is true if the zone is associated with a shared record group. Not writable
     */
    public readonly usingSrgAssociations: pulumi.Output<boolean>;
    /**
     * The name of the DNS view in which the zone resides
     */
    public readonly view: pulumi.Output<string>;
    /**
     * Determines the format of this zone - API default FORWARD. Cannot be updated.
     */
    public readonly zoneFormat: pulumi.Output<string>;

    /**
     * Create a ZoneForward resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneForwardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneForwardArgs | ZoneForwardState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ZoneForwardState = argsOrState as ZoneForwardState | undefined;
            inputs["address"] = state ? state.address : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["disable"] = state ? state.disable : undefined;
            inputs["displayDomain"] = state ? state.displayDomain : undefined;
            inputs["dnsFqdn"] = state ? state.dnsFqdn : undefined;
            inputs["forwardTos"] = state ? state.forwardTos : undefined;
            inputs["forwardersOnly"] = state ? state.forwardersOnly : undefined;
            inputs["forwardingServers"] = state ? state.forwardingServers : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["locked"] = state ? state.locked : undefined;
            inputs["lockedBy"] = state ? state.lockedBy : undefined;
            inputs["maskPrefix"] = state ? state.maskPrefix : undefined;
            inputs["nsGroup"] = state ? state.nsGroup : undefined;
            inputs["parent"] = state ? state.parent : undefined;
            inputs["prefix"] = state ? state.prefix : undefined;
            inputs["usingSrgAssociations"] = state ? state.usingSrgAssociations : undefined;
            inputs["view"] = state ? state.view : undefined;
            inputs["zoneFormat"] = state ? state.zoneFormat : undefined;
        } else {
            const args = argsOrState as ZoneForwardArgs | undefined;
            if (!args || args.forwardTos === undefined) {
                throw new Error("Missing required property 'forwardTos'");
            }
            inputs["comment"] = args ? args.comment : undefined;
            inputs["disable"] = args ? args.disable : undefined;
            inputs["displayDomain"] = args ? args.displayDomain : undefined;
            inputs["dnsFqdn"] = args ? args.dnsFqdn : undefined;
            inputs["forwardTos"] = args ? args.forwardTos : undefined;
            inputs["forwardersOnly"] = args ? args.forwardersOnly : undefined;
            inputs["forwardingServers"] = args ? args.forwardingServers : undefined;
            inputs["fqdn"] = args ? args.fqdn : undefined;
            inputs["locked"] = args ? args.locked : undefined;
            inputs["lockedBy"] = args ? args.lockedBy : undefined;
            inputs["maskPrefix"] = args ? args.maskPrefix : undefined;
            inputs["nsGroup"] = args ? args.nsGroup : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["prefix"] = args ? args.prefix : undefined;
            inputs["usingSrgAssociations"] = args ? args.usingSrgAssociations : undefined;
            inputs["view"] = args ? args.view : undefined;
            inputs["zoneFormat"] = args ? args.zoneFormat : undefined;
            inputs["address"] = undefined /*out*/;
        }
        super("infoblox:index/zoneForward:ZoneForward", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ZoneForward resources.
 */
export interface ZoneForwardState {
    /**
     * The IPv4 Address or IPv6 Address of the server that is serving this zone. Not writable
     */
    readonly address?: pulumi.Input<string>;
    /**
     * Comment for the zone; maximum 256 characters
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * Determines whether a zone is disabled or not
     */
    readonly disable?: pulumi.Input<boolean>;
    /**
     * The displayed name of the DNS zone.Not writable
     */
    readonly displayDomain?: pulumi.Input<string>;
    /**
     * The name of this DNS zone in punycode format.For a reverse zone, this is in “address/cidr” format.For other
     * zones, this is in FQDN format in punycode format.Cannot be updated
     */
    readonly dnsFqdn?: pulumi.Input<string>;
    /**
     * The primary preference list with Grid member names and/or External Server structs for this member.
     */
    readonly forwardTos?: pulumi.Input<pulumi.Input<{ address: pulumi.Input<string>, name: pulumi.Input<string>, sharedWithMsParentDelegation?: pulumi.Input<boolean>, stealth?: pulumi.Input<boolean>, tsigKey?: pulumi.Input<string>, tsigKeyAlg?: pulumi.Input<string>, tsigKeyName?: pulumi.Input<string>, useTsigKeyName?: pulumi.Input<boolean> }>[]>;
    /**
     * Determines if the appliance sends queries to forwarders only and not to other internal or Internet root servers
     */
    readonly forwardersOnly?: pulumi.Input<boolean>;
    /**
     * The primary preference list with Grid member names and/or External Server structs for this member.
     */
    readonly forwardingServers?: pulumi.Input<pulumi.Input<{ forwardTos?: pulumi.Input<pulumi.Input<{ address: pulumi.Input<string>, name: pulumi.Input<string>, sharedWithMsParentDelegation?: pulumi.Input<boolean>, stealth?: pulumi.Input<boolean>, tsigKey?: pulumi.Input<string>, tsigKeyAlg?: pulumi.Input<string>, tsigKeyName?: pulumi.Input<string>, useTsigKeyName?: pulumi.Input<boolean> }>[]>, forwardersOnly?: pulumi.Input<boolean>, name: pulumi.Input<string>, useOverrideForwarders?: pulumi.Input<boolean> }>[]>;
    /**
     * The name of this DNS zone. For a reverse zone, this is in “address/cidr” format
     */
    readonly fqdn?: pulumi.Input<string>;
    /**
     * If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes
     * only. The zone will continue to serve DNS data even when it is locked.The default value is False.
     */
    readonly locked?: pulumi.Input<boolean>;
    /**
     * The name of a superuser or the administrator who locked this zone.Not writable
     */
    readonly lockedBy?: pulumi.Input<string>;
    /**
     * IPv4 Netmask or IPv6 prefix for this zone.Not Writable
     */
    readonly maskPrefix?: pulumi.Input<string>;
    /**
     * A forwarding member name server group. Values with leading or trailing white space are not valid for this field. The
     * default value is undefined.
     */
    readonly nsGroup?: pulumi.Input<string>;
    /**
     * The parent zone of this zone. Note that when searching for reverse zones, the “in-addr.arpa” notation should be
     * used. Not writable.
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * The RFC2317 prefix value of this DNS zone.
     */
    readonly prefix?: pulumi.Input<string>;
    /**
     * This is true if the zone is associated with a shared record group. Not writable
     */
    readonly usingSrgAssociations?: pulumi.Input<boolean>;
    /**
     * The name of the DNS view in which the zone resides
     */
    readonly view?: pulumi.Input<string>;
    /**
     * Determines the format of this zone - API default FORWARD. Cannot be updated.
     */
    readonly zoneFormat?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ZoneForward resource.
 */
export interface ZoneForwardArgs {
    /**
     * Comment for the zone; maximum 256 characters
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * Determines whether a zone is disabled or not
     */
    readonly disable?: pulumi.Input<boolean>;
    /**
     * The displayed name of the DNS zone.Not writable
     */
    readonly displayDomain?: pulumi.Input<string>;
    /**
     * The name of this DNS zone in punycode format.For a reverse zone, this is in “address/cidr” format.For other
     * zones, this is in FQDN format in punycode format.Cannot be updated
     */
    readonly dnsFqdn?: pulumi.Input<string>;
    /**
     * The primary preference list with Grid member names and/or External Server structs for this member.
     */
    readonly forwardTos: pulumi.Input<pulumi.Input<{ address: pulumi.Input<string>, name: pulumi.Input<string>, sharedWithMsParentDelegation?: pulumi.Input<boolean>, stealth?: pulumi.Input<boolean>, tsigKey?: pulumi.Input<string>, tsigKeyAlg?: pulumi.Input<string>, tsigKeyName?: pulumi.Input<string>, useTsigKeyName?: pulumi.Input<boolean> }>[]>;
    /**
     * Determines if the appliance sends queries to forwarders only and not to other internal or Internet root servers
     */
    readonly forwardersOnly?: pulumi.Input<boolean>;
    /**
     * The primary preference list with Grid member names and/or External Server structs for this member.
     */
    readonly forwardingServers?: pulumi.Input<pulumi.Input<{ forwardTos?: pulumi.Input<pulumi.Input<{ address: pulumi.Input<string>, name: pulumi.Input<string>, sharedWithMsParentDelegation?: pulumi.Input<boolean>, stealth?: pulumi.Input<boolean>, tsigKey?: pulumi.Input<string>, tsigKeyAlg?: pulumi.Input<string>, tsigKeyName?: pulumi.Input<string>, useTsigKeyName?: pulumi.Input<boolean> }>[]>, forwardersOnly?: pulumi.Input<boolean>, name: pulumi.Input<string>, useOverrideForwarders?: pulumi.Input<boolean> }>[]>;
    /**
     * The name of this DNS zone. For a reverse zone, this is in “address/cidr” format
     */
    readonly fqdn?: pulumi.Input<string>;
    /**
     * If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes
     * only. The zone will continue to serve DNS data even when it is locked.The default value is False.
     */
    readonly locked?: pulumi.Input<boolean>;
    /**
     * The name of a superuser or the administrator who locked this zone.Not writable
     */
    readonly lockedBy?: pulumi.Input<string>;
    /**
     * IPv4 Netmask or IPv6 prefix for this zone.Not Writable
     */
    readonly maskPrefix?: pulumi.Input<string>;
    /**
     * A forwarding member name server group. Values with leading or trailing white space are not valid for this field. The
     * default value is undefined.
     */
    readonly nsGroup?: pulumi.Input<string>;
    /**
     * The parent zone of this zone. Note that when searching for reverse zones, the “in-addr.arpa” notation should be
     * used. Not writable.
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * The RFC2317 prefix value of this DNS zone.
     */
    readonly prefix?: pulumi.Input<string>;
    /**
     * This is true if the zone is associated with a shared record group. Not writable
     */
    readonly usingSrgAssociations?: pulumi.Input<boolean>;
    /**
     * The name of the DNS view in which the zone resides
     */
    readonly view?: pulumi.Input<string>;
    /**
     * Determines the format of this zone - API default FORWARD. Cannot be updated.
     */
    readonly zoneFormat?: pulumi.Input<string>;
}
