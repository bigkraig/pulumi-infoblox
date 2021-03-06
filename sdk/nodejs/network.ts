// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Network extends pulumi.CustomResource {
    /**
     * Get an existing Network resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkState, opts?: pulumi.CustomResourceOptions): Network {
        return new Network(name, <any>state, { ...opts, id: id });
    }

    /**
     * Authority for the DHCP network. Associated with the field use_authority
     */
    public readonly authority: pulumi.Output<boolean | undefined>;
    /**
     * This flag controls whether reverse zones are automatically created when the network is added. Cannot be updated, nor
     * is readable
     */
    public readonly autoCreateReversezone: pulumi.Output<boolean | undefined>;
    /**
     * Comment for the network, maximum 256 characters.
     */
    public readonly comment: pulumi.Output<string | undefined>;
    /**
     * Determines whether a network is disabled or not. When this is set to False, the network is enabled.
     */
    public readonly disable: pulumi.Output<boolean | undefined>;
    /**
     * The member that will run discovery for this network.
     */
    public readonly discoveryMember: pulumi.Output<string | undefined>;
    /**
     * The dynamic DNS updates flag of a DHCP network object. If set to True, the DHCP server sends DDNS updates to DNS
     * servers in the same Grid, and to external DNS servers.
     */
    public readonly enableDdns: pulumi.Output<boolean | undefined>;
    /**
     * Determines if DHCP thresholds are enabled for the network.
     */
    public readonly enableDhcpThresholds: pulumi.Output<boolean | undefined>;
    /**
     * Determines whether a discovery is enabled or not for this network. When this is set to False, the network discovery
     * is disabled.
     */
    public readonly enableDiscovery: pulumi.Output<boolean | undefined>;
    /**
     * The percentage of DHCP network usage threshold above which network usage is not expected and may warrant your
     * attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning
     * (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.
     */
    public readonly highWaterMark: pulumi.Output<number>;
    /**
     * The percentage of DHCP network usage below which the corresponding SNMP trap is reset. A number that specifies the
     * percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the
     * high watermark value.
     */
    public readonly highWaterMarkReset: pulumi.Output<number>;
    /**
     * The IPv4 Address of the network.
     */
    public readonly ipv4addr: pulumi.Output<string>;
    /**
     * An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database
     * before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive
     * value must be greater than 86400 seconds (1 day).
     */
    public readonly leaseScavengeTime: pulumi.Output<number>;
    /**
     * The percentage of DHCP network usage below which the Infoblox appliance generates a syslog message and sends a
     * warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.
     */
    public readonly lowWaterMark: pulumi.Output<number>;
    /**
     * The percentage of DHCP network usage threshold below which network usage is not expected and may warrant your
     * attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning
     * (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low
     * watermark reset value must be higher than the low watermark value.
     */
    public readonly lowWaterMarkReset: pulumi.Output<number>;
    /**
     * DHCP Member which is going to serve this network.
     */
    public readonly members: pulumi.Output<{ ipv4addr?: string, ipv6addr?: string, name?: string }[] | undefined>;
    /**
     * Number of bits in the network mask example: 8,16,24 etc
     */
    public readonly netmask: pulumi.Output<number>;
    /**
     * The network address in IPv4 Address/CIDR format.
     */
    public readonly network: pulumi.Output<string>;
    /**
     * The network container to which this network belongs (if any). Cannot be updated nor written
     */
    public readonly networkContainer: pulumi.Output<string>;
    /**
     * The name of the network view in which this network resides.
     */
    public readonly networkView: pulumi.Output<string>;
    /**
     * DHCP Related] Options such as DNS servers, gateway, ntp, etc
     */
    public readonly options: pulumi.Output<{ name?: string, num?: number, useOption?: boolean, value?: string, vendorClass?: string }[]>;
    /**
     * If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the
     * leases are permanently deleted.
     */
    public readonly recycleLeases: pulumi.Output<boolean | undefined>;
    /**
     * Restarts the member service. Not readable
     */
    public readonly restartIfNeeded: pulumi.Output<boolean | undefined>;
    /**
     * This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.
     */
    public readonly updateDnsOnLeaseRenewal: pulumi.Output<boolean | undefined>;
    public readonly useAuthority: pulumi.Output<boolean | undefined>;
    public readonly useEnableDdns: pulumi.Output<boolean | undefined>;
    public readonly useEnableDhcpThresholds: pulumi.Output<boolean | undefined>;
    public readonly useEnableDiscovery: pulumi.Output<boolean | undefined>;
    public readonly useOptions: pulumi.Output<boolean | undefined>;
    public readonly useRecycleLeases: pulumi.Output<boolean | undefined>;

    /**
     * Create a Network resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkArgs | NetworkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: NetworkState = argsOrState as NetworkState | undefined;
            inputs["authority"] = state ? state.authority : undefined;
            inputs["autoCreateReversezone"] = state ? state.autoCreateReversezone : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["disable"] = state ? state.disable : undefined;
            inputs["discoveryMember"] = state ? state.discoveryMember : undefined;
            inputs["enableDdns"] = state ? state.enableDdns : undefined;
            inputs["enableDhcpThresholds"] = state ? state.enableDhcpThresholds : undefined;
            inputs["enableDiscovery"] = state ? state.enableDiscovery : undefined;
            inputs["highWaterMark"] = state ? state.highWaterMark : undefined;
            inputs["highWaterMarkReset"] = state ? state.highWaterMarkReset : undefined;
            inputs["ipv4addr"] = state ? state.ipv4addr : undefined;
            inputs["leaseScavengeTime"] = state ? state.leaseScavengeTime : undefined;
            inputs["lowWaterMark"] = state ? state.lowWaterMark : undefined;
            inputs["lowWaterMarkReset"] = state ? state.lowWaterMarkReset : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["netmask"] = state ? state.netmask : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["networkContainer"] = state ? state.networkContainer : undefined;
            inputs["networkView"] = state ? state.networkView : undefined;
            inputs["options"] = state ? state.options : undefined;
            inputs["recycleLeases"] = state ? state.recycleLeases : undefined;
            inputs["restartIfNeeded"] = state ? state.restartIfNeeded : undefined;
            inputs["updateDnsOnLeaseRenewal"] = state ? state.updateDnsOnLeaseRenewal : undefined;
            inputs["useAuthority"] = state ? state.useAuthority : undefined;
            inputs["useEnableDdns"] = state ? state.useEnableDdns : undefined;
            inputs["useEnableDhcpThresholds"] = state ? state.useEnableDhcpThresholds : undefined;
            inputs["useEnableDiscovery"] = state ? state.useEnableDiscovery : undefined;
            inputs["useOptions"] = state ? state.useOptions : undefined;
            inputs["useRecycleLeases"] = state ? state.useRecycleLeases : undefined;
        } else {
            const args = argsOrState as NetworkArgs | undefined;
            if (!args || args.network === undefined) {
                throw new Error("Missing required property 'network'");
            }
            inputs["authority"] = args ? args.authority : undefined;
            inputs["autoCreateReversezone"] = args ? args.autoCreateReversezone : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["disable"] = args ? args.disable : undefined;
            inputs["discoveryMember"] = args ? args.discoveryMember : undefined;
            inputs["enableDdns"] = args ? args.enableDdns : undefined;
            inputs["enableDhcpThresholds"] = args ? args.enableDhcpThresholds : undefined;
            inputs["enableDiscovery"] = args ? args.enableDiscovery : undefined;
            inputs["highWaterMark"] = args ? args.highWaterMark : undefined;
            inputs["highWaterMarkReset"] = args ? args.highWaterMarkReset : undefined;
            inputs["ipv4addr"] = args ? args.ipv4addr : undefined;
            inputs["leaseScavengeTime"] = args ? args.leaseScavengeTime : undefined;
            inputs["lowWaterMark"] = args ? args.lowWaterMark : undefined;
            inputs["lowWaterMarkReset"] = args ? args.lowWaterMarkReset : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["netmask"] = args ? args.netmask : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["networkContainer"] = args ? args.networkContainer : undefined;
            inputs["networkView"] = args ? args.networkView : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["recycleLeases"] = args ? args.recycleLeases : undefined;
            inputs["restartIfNeeded"] = args ? args.restartIfNeeded : undefined;
            inputs["updateDnsOnLeaseRenewal"] = args ? args.updateDnsOnLeaseRenewal : undefined;
            inputs["useAuthority"] = args ? args.useAuthority : undefined;
            inputs["useEnableDdns"] = args ? args.useEnableDdns : undefined;
            inputs["useEnableDhcpThresholds"] = args ? args.useEnableDhcpThresholds : undefined;
            inputs["useEnableDiscovery"] = args ? args.useEnableDiscovery : undefined;
            inputs["useOptions"] = args ? args.useOptions : undefined;
            inputs["useRecycleLeases"] = args ? args.useRecycleLeases : undefined;
        }
        super("infoblox:index/network:Network", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Network resources.
 */
export interface NetworkState {
    /**
     * Authority for the DHCP network. Associated with the field use_authority
     */
    readonly authority?: pulumi.Input<boolean>;
    /**
     * This flag controls whether reverse zones are automatically created when the network is added. Cannot be updated, nor
     * is readable
     */
    readonly autoCreateReversezone?: pulumi.Input<boolean>;
    /**
     * Comment for the network, maximum 256 characters.
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * Determines whether a network is disabled or not. When this is set to False, the network is enabled.
     */
    readonly disable?: pulumi.Input<boolean>;
    /**
     * The member that will run discovery for this network.
     */
    readonly discoveryMember?: pulumi.Input<string>;
    /**
     * The dynamic DNS updates flag of a DHCP network object. If set to True, the DHCP server sends DDNS updates to DNS
     * servers in the same Grid, and to external DNS servers.
     */
    readonly enableDdns?: pulumi.Input<boolean>;
    /**
     * Determines if DHCP thresholds are enabled for the network.
     */
    readonly enableDhcpThresholds?: pulumi.Input<boolean>;
    /**
     * Determines whether a discovery is enabled or not for this network. When this is set to False, the network discovery
     * is disabled.
     */
    readonly enableDiscovery?: pulumi.Input<boolean>;
    /**
     * The percentage of DHCP network usage threshold above which network usage is not expected and may warrant your
     * attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning
     * (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.
     */
    readonly highWaterMark?: pulumi.Input<number>;
    /**
     * The percentage of DHCP network usage below which the corresponding SNMP trap is reset. A number that specifies the
     * percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the
     * high watermark value.
     */
    readonly highWaterMarkReset?: pulumi.Input<number>;
    /**
     * The IPv4 Address of the network.
     */
    readonly ipv4addr?: pulumi.Input<string>;
    /**
     * An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database
     * before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive
     * value must be greater than 86400 seconds (1 day).
     */
    readonly leaseScavengeTime?: pulumi.Input<number>;
    /**
     * The percentage of DHCP network usage below which the Infoblox appliance generates a syslog message and sends a
     * warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.
     */
    readonly lowWaterMark?: pulumi.Input<number>;
    /**
     * The percentage of DHCP network usage threshold below which network usage is not expected and may warrant your
     * attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning
     * (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low
     * watermark reset value must be higher than the low watermark value.
     */
    readonly lowWaterMarkReset?: pulumi.Input<number>;
    /**
     * DHCP Member which is going to serve this network.
     */
    readonly members?: pulumi.Input<pulumi.Input<{ ipv4addr?: pulumi.Input<string>, ipv6addr?: pulumi.Input<string>, name?: pulumi.Input<string> }>[]>;
    /**
     * Number of bits in the network mask example: 8,16,24 etc
     */
    readonly netmask?: pulumi.Input<number>;
    /**
     * The network address in IPv4 Address/CIDR format.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * The network container to which this network belongs (if any). Cannot be updated nor written
     */
    readonly networkContainer?: pulumi.Input<string>;
    /**
     * The name of the network view in which this network resides.
     */
    readonly networkView?: pulumi.Input<string>;
    /**
     * DHCP Related] Options such as DNS servers, gateway, ntp, etc
     */
    readonly options?: pulumi.Input<pulumi.Input<{ name?: pulumi.Input<string>, num?: pulumi.Input<number>, useOption?: pulumi.Input<boolean>, value?: pulumi.Input<string>, vendorClass?: pulumi.Input<string> }>[]>;
    /**
     * If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the
     * leases are permanently deleted.
     */
    readonly recycleLeases?: pulumi.Input<boolean>;
    /**
     * Restarts the member service. Not readable
     */
    readonly restartIfNeeded?: pulumi.Input<boolean>;
    /**
     * This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.
     */
    readonly updateDnsOnLeaseRenewal?: pulumi.Input<boolean>;
    readonly useAuthority?: pulumi.Input<boolean>;
    readonly useEnableDdns?: pulumi.Input<boolean>;
    readonly useEnableDhcpThresholds?: pulumi.Input<boolean>;
    readonly useEnableDiscovery?: pulumi.Input<boolean>;
    readonly useOptions?: pulumi.Input<boolean>;
    readonly useRecycleLeases?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Network resource.
 */
export interface NetworkArgs {
    /**
     * Authority for the DHCP network. Associated with the field use_authority
     */
    readonly authority?: pulumi.Input<boolean>;
    /**
     * This flag controls whether reverse zones are automatically created when the network is added. Cannot be updated, nor
     * is readable
     */
    readonly autoCreateReversezone?: pulumi.Input<boolean>;
    /**
     * Comment for the network, maximum 256 characters.
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * Determines whether a network is disabled or not. When this is set to False, the network is enabled.
     */
    readonly disable?: pulumi.Input<boolean>;
    /**
     * The member that will run discovery for this network.
     */
    readonly discoveryMember?: pulumi.Input<string>;
    /**
     * The dynamic DNS updates flag of a DHCP network object. If set to True, the DHCP server sends DDNS updates to DNS
     * servers in the same Grid, and to external DNS servers.
     */
    readonly enableDdns?: pulumi.Input<boolean>;
    /**
     * Determines if DHCP thresholds are enabled for the network.
     */
    readonly enableDhcpThresholds?: pulumi.Input<boolean>;
    /**
     * Determines whether a discovery is enabled or not for this network. When this is set to False, the network discovery
     * is disabled.
     */
    readonly enableDiscovery?: pulumi.Input<boolean>;
    /**
     * The percentage of DHCP network usage threshold above which network usage is not expected and may warrant your
     * attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning
     * (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.
     */
    readonly highWaterMark?: pulumi.Input<number>;
    /**
     * The percentage of DHCP network usage below which the corresponding SNMP trap is reset. A number that specifies the
     * percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the
     * high watermark value.
     */
    readonly highWaterMarkReset?: pulumi.Input<number>;
    /**
     * The IPv4 Address of the network.
     */
    readonly ipv4addr?: pulumi.Input<string>;
    /**
     * An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database
     * before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive
     * value must be greater than 86400 seconds (1 day).
     */
    readonly leaseScavengeTime?: pulumi.Input<number>;
    /**
     * The percentage of DHCP network usage below which the Infoblox appliance generates a syslog message and sends a
     * warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.
     */
    readonly lowWaterMark?: pulumi.Input<number>;
    /**
     * The percentage of DHCP network usage threshold below which network usage is not expected and may warrant your
     * attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning
     * (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low
     * watermark reset value must be higher than the low watermark value.
     */
    readonly lowWaterMarkReset?: pulumi.Input<number>;
    /**
     * DHCP Member which is going to serve this network.
     */
    readonly members?: pulumi.Input<pulumi.Input<{ ipv4addr?: pulumi.Input<string>, ipv6addr?: pulumi.Input<string>, name?: pulumi.Input<string> }>[]>;
    /**
     * Number of bits in the network mask example: 8,16,24 etc
     */
    readonly netmask?: pulumi.Input<number>;
    /**
     * The network address in IPv4 Address/CIDR format.
     */
    readonly network: pulumi.Input<string>;
    /**
     * The network container to which this network belongs (if any). Cannot be updated nor written
     */
    readonly networkContainer?: pulumi.Input<string>;
    /**
     * The name of the network view in which this network resides.
     */
    readonly networkView?: pulumi.Input<string>;
    /**
     * DHCP Related] Options such as DNS servers, gateway, ntp, etc
     */
    readonly options?: pulumi.Input<pulumi.Input<{ name?: pulumi.Input<string>, num?: pulumi.Input<number>, useOption?: pulumi.Input<boolean>, value?: pulumi.Input<string>, vendorClass?: pulumi.Input<string> }>[]>;
    /**
     * If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the
     * leases are permanently deleted.
     */
    readonly recycleLeases?: pulumi.Input<boolean>;
    /**
     * Restarts the member service. Not readable
     */
    readonly restartIfNeeded?: pulumi.Input<boolean>;
    /**
     * This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.
     */
    readonly updateDnsOnLeaseRenewal?: pulumi.Input<boolean>;
    readonly useAuthority?: pulumi.Input<boolean>;
    readonly useEnableDdns?: pulumi.Input<boolean>;
    readonly useEnableDhcpThresholds?: pulumi.Input<boolean>;
    readonly useEnableDiscovery?: pulumi.Input<boolean>;
    readonly useOptions?: pulumi.Input<boolean>;
    readonly useRecycleLeases?: pulumi.Input<boolean>;
}
