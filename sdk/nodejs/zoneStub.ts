// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ZoneStub extends pulumi.CustomResource {
    /**
     * Get an existing ZoneStub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneStubState, opts?: pulumi.CustomResourceOptions): ZoneStub {
        return new ZoneStub(name, <any>state, { ...opts, id: id });
    }

    /**
     * Comment for the zone; maximum 256 characters
     */
    public readonly comment: pulumi.Output<string | undefined>;
    /**
     * Is the zone disabled
     */
    public readonly disable: pulumi.Output<boolean | undefined>;
    /**
     * Is forward disabled for this zone
     */
    public readonly disableForwarding: pulumi.Output<boolean | undefined>;
    /**
     * Name of the external name server group
     */
    public readonly externalNsgroup: pulumi.Output<string | undefined>;
    /**
     * Fqdn for the zone
     */
    public readonly fqdn: pulumi.Output<string>;
    /**
     * Is the record locked to prevent changes
     */
    public readonly locked: pulumi.Output<boolean | undefined>;
    /**
     * IPv4 Netmask or IPv6 prefix for this zone.
     */
    public readonly maskPrefix: pulumi.Output<string | undefined>;
    /**
     * Name of the name server group
     */
    public readonly nsgroup: pulumi.Output<string | undefined>;
    /**
     * IPv4 Netmask or IPv6 prefix for this zone.
     */
    public readonly prefix: pulumi.Output<string | undefined>;
    /**
     * The primary preference list with Grid member names and/or External Server structs for this member.
     */
    public readonly stubFroms: pulumi.Output<{ address: string, name: string, sharedWithMsParentDelegation: boolean, stealth?: boolean, tsigKey?: string, tsigKeyAlg?: string, tsigKeyName?: string, useTsigKeyName?: boolean }[] | undefined>;
    /**
     * The grid primary servers for this zone.
     */
    public readonly stubMembers: pulumi.Output<{ enablePreferredPrimaries?: boolean, gridReplicate?: boolean, lead?: boolean, name: string, preferredPrimaries?: { address: string, name: string, sharedWithMsParentDelegation: boolean, stealth?: boolean, tsigKey?: string, tsigKeyAlg?: string, tsigKeyName?: string, useTsigKeyName?: boolean }[], stealth?: boolean }[] | undefined>;
    /**
     * The name of the DNS view in which the zone resides
     */
    public readonly view: pulumi.Output<string>;
    /**
     * Determines the format of this zone - API default FORWARD
     */
    public readonly zoneFormat: pulumi.Output<string>;

    /**
     * Create a ZoneStub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneStubArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneStubArgs | ZoneStubState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ZoneStubState = argsOrState as ZoneStubState | undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["disable"] = state ? state.disable : undefined;
            inputs["disableForwarding"] = state ? state.disableForwarding : undefined;
            inputs["externalNsgroup"] = state ? state.externalNsgroup : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["locked"] = state ? state.locked : undefined;
            inputs["maskPrefix"] = state ? state.maskPrefix : undefined;
            inputs["nsgroup"] = state ? state.nsgroup : undefined;
            inputs["prefix"] = state ? state.prefix : undefined;
            inputs["stubFroms"] = state ? state.stubFroms : undefined;
            inputs["stubMembers"] = state ? state.stubMembers : undefined;
            inputs["view"] = state ? state.view : undefined;
            inputs["zoneFormat"] = state ? state.zoneFormat : undefined;
        } else {
            const args = argsOrState as ZoneStubArgs | undefined;
            if (!args || args.fqdn === undefined) {
                throw new Error("Missing required property 'fqdn'");
            }
            inputs["comment"] = args ? args.comment : undefined;
            inputs["disable"] = args ? args.disable : undefined;
            inputs["disableForwarding"] = args ? args.disableForwarding : undefined;
            inputs["externalNsgroup"] = args ? args.externalNsgroup : undefined;
            inputs["fqdn"] = args ? args.fqdn : undefined;
            inputs["locked"] = args ? args.locked : undefined;
            inputs["maskPrefix"] = args ? args.maskPrefix : undefined;
            inputs["nsgroup"] = args ? args.nsgroup : undefined;
            inputs["prefix"] = args ? args.prefix : undefined;
            inputs["stubFroms"] = args ? args.stubFroms : undefined;
            inputs["stubMembers"] = args ? args.stubMembers : undefined;
            inputs["view"] = args ? args.view : undefined;
            inputs["zoneFormat"] = args ? args.zoneFormat : undefined;
        }
        super("infoblox:index/zoneStub:ZoneStub", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ZoneStub resources.
 */
export interface ZoneStubState {
    /**
     * Comment for the zone; maximum 256 characters
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * Is the zone disabled
     */
    readonly disable?: pulumi.Input<boolean>;
    /**
     * Is forward disabled for this zone
     */
    readonly disableForwarding?: pulumi.Input<boolean>;
    /**
     * Name of the external name server group
     */
    readonly externalNsgroup?: pulumi.Input<string>;
    /**
     * Fqdn for the zone
     */
    readonly fqdn?: pulumi.Input<string>;
    /**
     * Is the record locked to prevent changes
     */
    readonly locked?: pulumi.Input<boolean>;
    /**
     * IPv4 Netmask or IPv6 prefix for this zone.
     */
    readonly maskPrefix?: pulumi.Input<string>;
    /**
     * Name of the name server group
     */
    readonly nsgroup?: pulumi.Input<string>;
    /**
     * IPv4 Netmask or IPv6 prefix for this zone.
     */
    readonly prefix?: pulumi.Input<string>;
    /**
     * The primary preference list with Grid member names and/or External Server structs for this member.
     */
    readonly stubFroms?: pulumi.Input<pulumi.Input<{ address: pulumi.Input<string>, name: pulumi.Input<string>, sharedWithMsParentDelegation?: pulumi.Input<boolean>, stealth?: pulumi.Input<boolean>, tsigKey?: pulumi.Input<string>, tsigKeyAlg?: pulumi.Input<string>, tsigKeyName?: pulumi.Input<string>, useTsigKeyName?: pulumi.Input<boolean> }>[]>;
    /**
     * The grid primary servers for this zone.
     */
    readonly stubMembers?: pulumi.Input<pulumi.Input<{ enablePreferredPrimaries?: pulumi.Input<boolean>, gridReplicate?: pulumi.Input<boolean>, lead?: pulumi.Input<boolean>, name: pulumi.Input<string>, preferredPrimaries?: pulumi.Input<pulumi.Input<{ address: pulumi.Input<string>, name: pulumi.Input<string>, sharedWithMsParentDelegation?: pulumi.Input<boolean>, stealth?: pulumi.Input<boolean>, tsigKey?: pulumi.Input<string>, tsigKeyAlg?: pulumi.Input<string>, tsigKeyName?: pulumi.Input<string>, useTsigKeyName?: pulumi.Input<boolean> }>[]>, stealth?: pulumi.Input<boolean> }>[]>;
    /**
     * The name of the DNS view in which the zone resides
     */
    readonly view?: pulumi.Input<string>;
    /**
     * Determines the format of this zone - API default FORWARD
     */
    readonly zoneFormat?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ZoneStub resource.
 */
export interface ZoneStubArgs {
    /**
     * Comment for the zone; maximum 256 characters
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * Is the zone disabled
     */
    readonly disable?: pulumi.Input<boolean>;
    /**
     * Is forward disabled for this zone
     */
    readonly disableForwarding?: pulumi.Input<boolean>;
    /**
     * Name of the external name server group
     */
    readonly externalNsgroup?: pulumi.Input<string>;
    /**
     * Fqdn for the zone
     */
    readonly fqdn: pulumi.Input<string>;
    /**
     * Is the record locked to prevent changes
     */
    readonly locked?: pulumi.Input<boolean>;
    /**
     * IPv4 Netmask or IPv6 prefix for this zone.
     */
    readonly maskPrefix?: pulumi.Input<string>;
    /**
     * Name of the name server group
     */
    readonly nsgroup?: pulumi.Input<string>;
    /**
     * IPv4 Netmask or IPv6 prefix for this zone.
     */
    readonly prefix?: pulumi.Input<string>;
    /**
     * The primary preference list with Grid member names and/or External Server structs for this member.
     */
    readonly stubFroms?: pulumi.Input<pulumi.Input<{ address: pulumi.Input<string>, name: pulumi.Input<string>, sharedWithMsParentDelegation?: pulumi.Input<boolean>, stealth?: pulumi.Input<boolean>, tsigKey?: pulumi.Input<string>, tsigKeyAlg?: pulumi.Input<string>, tsigKeyName?: pulumi.Input<string>, useTsigKeyName?: pulumi.Input<boolean> }>[]>;
    /**
     * The grid primary servers for this zone.
     */
    readonly stubMembers?: pulumi.Input<pulumi.Input<{ enablePreferredPrimaries?: pulumi.Input<boolean>, gridReplicate?: pulumi.Input<boolean>, lead?: pulumi.Input<boolean>, name: pulumi.Input<string>, preferredPrimaries?: pulumi.Input<pulumi.Input<{ address: pulumi.Input<string>, name: pulumi.Input<string>, sharedWithMsParentDelegation?: pulumi.Input<boolean>, stealth?: pulumi.Input<boolean>, tsigKey?: pulumi.Input<string>, tsigKeyAlg?: pulumi.Input<string>, tsigKeyName?: pulumi.Input<string>, useTsigKeyName?: pulumi.Input<boolean> }>[]>, stealth?: pulumi.Input<boolean> }>[]>;
    /**
     * The name of the DNS view in which the zone resides
     */
    readonly view?: pulumi.Input<string>;
    /**
     * Determines the format of this zone - API default FORWARD
     */
    readonly zoneFormat?: pulumi.Input<string>;
}
