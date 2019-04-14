// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class CNameRecord extends pulumi.CustomResource {
    /**
     * Get an existing CNameRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CNameRecordState, opts?: pulumi.CustomResourceOptions): CNameRecord {
        return new CNameRecord(name, <any>state, { ...opts, id: id });
    }

    /**
     * Canonical name in FQDN format
     */
    public readonly canonical: pulumi.Output<string>;
    /**
     * Comment for the record; maximum 256 characters
     */
    public readonly comment: pulumi.Output<string | undefined>;
    /**
     * The time of the record creation in Epoch seconds format.
     */
    public /*out*/ readonly creationTime: pulumi.Output<number>;
    /**
     * The record creator.Valid values:DYNAMIC,STATIC,SYSTEM.Defaults to STATIC
     */
    public readonly creator: pulumi.Output<string>;
    /**
     * The GSS-TSIG principal that owns this record
     */
    public readonly ddnsPrincipal: pulumi.Output<string | undefined>;
    /**
     * Determines if the DDNS updates for this record are allowed or not
     */
    public readonly ddnsProtected: pulumi.Output<boolean | undefined>;
    /**
     * Determines if the record is disabled or not. False means that the record is enabled.
     */
    public readonly disable: pulumi.Output<boolean | undefined>;
    /**
     * The name for the CNAME record in punycode format. Values with leading or trailing white space are not valid for this
     * field. Cannot be written nor updated.
     */
    public readonly dnsName: pulumi.Output<string>;
    /**
     * Determines if the reclamation is allowed for the record or not.
     */
    public readonly forbidReclamation: pulumi.Output<boolean | undefined>;
    /**
     * The name for a CNAME record in FQDN format
     */
    public readonly name: pulumi.Output<string>;
    /**
     * Determines if the record is reclaimable or not. Cannot be updated/written
     */
    public readonly reclaimable: pulumi.Output<boolean>;
    /**
     * The name of the shared record group in which the record resides. This field exists only on db_objects if this record
     * is a shared record. Cannot be updated/written
     */
    public readonly sharedRecordGroup: pulumi.Output<string>;
    /**
     * The Time To Live assigned to CNAME
     */
    public readonly ttl: pulumi.Output<number>;
    /**
     * Use flag for: ttl
     */
    public readonly useTtl: pulumi.Output<boolean>;
    /**
     * The name of the DNS View in which the record resides
     */
    public readonly view: pulumi.Output<string | undefined>;
    /**
     * DNS Zone for the record
     */
    public /*out*/ readonly zone: pulumi.Output<string>;

    /**
     * Create a CNameRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CNameRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CNameRecordArgs | CNameRecordState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: CNameRecordState = argsOrState as CNameRecordState | undefined;
            inputs["canonical"] = state ? state.canonical : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["creationTime"] = state ? state.creationTime : undefined;
            inputs["creator"] = state ? state.creator : undefined;
            inputs["ddnsPrincipal"] = state ? state.ddnsPrincipal : undefined;
            inputs["ddnsProtected"] = state ? state.ddnsProtected : undefined;
            inputs["disable"] = state ? state.disable : undefined;
            inputs["dnsName"] = state ? state.dnsName : undefined;
            inputs["forbidReclamation"] = state ? state.forbidReclamation : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["reclaimable"] = state ? state.reclaimable : undefined;
            inputs["sharedRecordGroup"] = state ? state.sharedRecordGroup : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["useTtl"] = state ? state.useTtl : undefined;
            inputs["view"] = state ? state.view : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as CNameRecordArgs | undefined;
            if (!args || args.canonical === undefined) {
                throw new Error("Missing required property 'canonical'");
            }
            inputs["canonical"] = args ? args.canonical : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["creator"] = args ? args.creator : undefined;
            inputs["ddnsPrincipal"] = args ? args.ddnsPrincipal : undefined;
            inputs["ddnsProtected"] = args ? args.ddnsProtected : undefined;
            inputs["disable"] = args ? args.disable : undefined;
            inputs["dnsName"] = args ? args.dnsName : undefined;
            inputs["forbidReclamation"] = args ? args.forbidReclamation : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["reclaimable"] = args ? args.reclaimable : undefined;
            inputs["sharedRecordGroup"] = args ? args.sharedRecordGroup : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["useTtl"] = args ? args.useTtl : undefined;
            inputs["view"] = args ? args.view : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["zone"] = undefined /*out*/;
        }
        super("infoblox:index/cNameRecord:CNameRecord", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CNameRecord resources.
 */
export interface CNameRecordState {
    /**
     * Canonical name in FQDN format
     */
    readonly canonical?: pulumi.Input<string>;
    /**
     * Comment for the record; maximum 256 characters
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * The time of the record creation in Epoch seconds format.
     */
    readonly creationTime?: pulumi.Input<number>;
    /**
     * The record creator.Valid values:DYNAMIC,STATIC,SYSTEM.Defaults to STATIC
     */
    readonly creator?: pulumi.Input<string>;
    /**
     * The GSS-TSIG principal that owns this record
     */
    readonly ddnsPrincipal?: pulumi.Input<string>;
    /**
     * Determines if the DDNS updates for this record are allowed or not
     */
    readonly ddnsProtected?: pulumi.Input<boolean>;
    /**
     * Determines if the record is disabled or not. False means that the record is enabled.
     */
    readonly disable?: pulumi.Input<boolean>;
    /**
     * The name for the CNAME record in punycode format. Values with leading or trailing white space are not valid for this
     * field. Cannot be written nor updated.
     */
    readonly dnsName?: pulumi.Input<string>;
    /**
     * Determines if the reclamation is allowed for the record or not.
     */
    readonly forbidReclamation?: pulumi.Input<boolean>;
    /**
     * The name for a CNAME record in FQDN format
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Determines if the record is reclaimable or not. Cannot be updated/written
     */
    readonly reclaimable?: pulumi.Input<boolean>;
    /**
     * The name of the shared record group in which the record resides. This field exists only on db_objects if this record
     * is a shared record. Cannot be updated/written
     */
    readonly sharedRecordGroup?: pulumi.Input<string>;
    /**
     * The Time To Live assigned to CNAME
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * Use flag for: ttl
     */
    readonly useTtl?: pulumi.Input<boolean>;
    /**
     * The name of the DNS View in which the record resides
     */
    readonly view?: pulumi.Input<string>;
    /**
     * DNS Zone for the record
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CNameRecord resource.
 */
export interface CNameRecordArgs {
    /**
     * Canonical name in FQDN format
     */
    readonly canonical: pulumi.Input<string>;
    /**
     * Comment for the record; maximum 256 characters
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * The record creator.Valid values:DYNAMIC,STATIC,SYSTEM.Defaults to STATIC
     */
    readonly creator?: pulumi.Input<string>;
    /**
     * The GSS-TSIG principal that owns this record
     */
    readonly ddnsPrincipal?: pulumi.Input<string>;
    /**
     * Determines if the DDNS updates for this record are allowed or not
     */
    readonly ddnsProtected?: pulumi.Input<boolean>;
    /**
     * Determines if the record is disabled or not. False means that the record is enabled.
     */
    readonly disable?: pulumi.Input<boolean>;
    /**
     * The name for the CNAME record in punycode format. Values with leading or trailing white space are not valid for this
     * field. Cannot be written nor updated.
     */
    readonly dnsName?: pulumi.Input<string>;
    /**
     * Determines if the reclamation is allowed for the record or not.
     */
    readonly forbidReclamation?: pulumi.Input<boolean>;
    /**
     * The name for a CNAME record in FQDN format
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Determines if the record is reclaimable or not. Cannot be updated/written
     */
    readonly reclaimable?: pulumi.Input<boolean>;
    /**
     * The name of the shared record group in which the record resides. This field exists only on db_objects if this record
     * is a shared record. Cannot be updated/written
     */
    readonly sharedRecordGroup?: pulumi.Input<string>;
    /**
     * The Time To Live assigned to CNAME
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * Use flag for: ttl
     */
    readonly useTtl?: pulumi.Input<boolean>;
    /**
     * The name of the DNS View in which the record resides
     */
    readonly view?: pulumi.Input<string>;
}