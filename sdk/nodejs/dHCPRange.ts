// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class DHCPRange extends pulumi.CustomResource {
    /**
     * Get an existing DHCPRange resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DHCPRangeState, opts?: pulumi.CustomResourceOptions): DHCPRange {
        return new DHCPRange(name, <any>state, { ...opts, id: id });
    }

    public readonly comment: pulumi.Output<string | undefined>;
    /**
     * The IPv4 Address end address of the range.
     */
    public readonly endAddr: pulumi.Output<string>;
    /**
     * Infoblox DHCP member that serves this range
     */
    public readonly members: pulumi.Output<{ ipv4addr?: string, name?: string }[] | undefined>;
    /**
     * This field contains the name of the Microsoft scope.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The network to which this range belongs, in IPv4 Address/CIDR format.
     */
    public readonly network: pulumi.Output<string | undefined>;
    /**
     * The name of the network view in which this range resides.
     */
    public readonly networkView: pulumi.Output<string | undefined>;
    /**
     * Restarts any services if required by this change. Default: true.
     */
    public readonly restartIfNeeded: pulumi.Output<boolean | undefined>;
    /**
     * Must be set to 'MEMBER' if member is specified
     */
    public readonly serverAssociationType: pulumi.Output<string | undefined>;
    /**
     * The IPv4 Address starting address of the range.
     */
    public readonly startAddr: pulumi.Output<string>;

    /**
     * Create a DHCPRange resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DHCPRangeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DHCPRangeArgs | DHCPRangeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DHCPRangeState = argsOrState as DHCPRangeState | undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["endAddr"] = state ? state.endAddr : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["networkView"] = state ? state.networkView : undefined;
            inputs["restartIfNeeded"] = state ? state.restartIfNeeded : undefined;
            inputs["serverAssociationType"] = state ? state.serverAssociationType : undefined;
            inputs["startAddr"] = state ? state.startAddr : undefined;
        } else {
            const args = argsOrState as DHCPRangeArgs | undefined;
            if (!args || args.endAddr === undefined) {
                throw new Error("Missing required property 'endAddr'");
            }
            if (!args || args.startAddr === undefined) {
                throw new Error("Missing required property 'startAddr'");
            }
            inputs["comment"] = args ? args.comment : undefined;
            inputs["endAddr"] = args ? args.endAddr : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["networkView"] = args ? args.networkView : undefined;
            inputs["restartIfNeeded"] = args ? args.restartIfNeeded : undefined;
            inputs["serverAssociationType"] = args ? args.serverAssociationType : undefined;
            inputs["startAddr"] = args ? args.startAddr : undefined;
        }
        super("infoblox:index/dHCPRange:DHCPRange", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DHCPRange resources.
 */
export interface DHCPRangeState {
    readonly comment?: pulumi.Input<string>;
    /**
     * The IPv4 Address end address of the range.
     */
    readonly endAddr?: pulumi.Input<string>;
    /**
     * Infoblox DHCP member that serves this range
     */
    readonly members?: pulumi.Input<pulumi.Input<{ ipv4addr?: pulumi.Input<string>, name?: pulumi.Input<string> }>[]>;
    /**
     * This field contains the name of the Microsoft scope.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network to which this range belongs, in IPv4 Address/CIDR format.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * The name of the network view in which this range resides.
     */
    readonly networkView?: pulumi.Input<string>;
    /**
     * Restarts any services if required by this change. Default: true.
     */
    readonly restartIfNeeded?: pulumi.Input<boolean>;
    /**
     * Must be set to 'MEMBER' if member is specified
     */
    readonly serverAssociationType?: pulumi.Input<string>;
    /**
     * The IPv4 Address starting address of the range.
     */
    readonly startAddr?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DHCPRange resource.
 */
export interface DHCPRangeArgs {
    readonly comment?: pulumi.Input<string>;
    /**
     * The IPv4 Address end address of the range.
     */
    readonly endAddr: pulumi.Input<string>;
    /**
     * Infoblox DHCP member that serves this range
     */
    readonly members?: pulumi.Input<pulumi.Input<{ ipv4addr?: pulumi.Input<string>, name?: pulumi.Input<string> }>[]>;
    /**
     * This field contains the name of the Microsoft scope.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network to which this range belongs, in IPv4 Address/CIDR format.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * The name of the network view in which this range resides.
     */
    readonly networkView?: pulumi.Input<string>;
    /**
     * Restarts any services if required by this change. Default: true.
     */
    readonly restartIfNeeded?: pulumi.Input<boolean>;
    /**
     * Must be set to 'MEMBER' if member is specified
     */
    readonly serverAssociationType?: pulumi.Input<string>;
    /**
     * The IPv4 Address starting address of the range.
     */
    readonly startAddr: pulumi.Input<string>;
}
