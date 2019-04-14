# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Network(pulumi.CustomResource):
    authority: pulumi.Output[bool]
    auto_create_reversezone: pulumi.Output[bool]
    comment: pulumi.Output[str]
    disable: pulumi.Output[bool]
    discovery_member: pulumi.Output[str]
    enable_ddns: pulumi.Output[bool]
    enable_dhcp_thresholds: pulumi.Output[bool]
    enable_discovery: pulumi.Output[bool]
    high_water_mark: pulumi.Output[float]
    high_water_mark_reset: pulumi.Output[float]
    ipv4addr: pulumi.Output[str]
    lease_scavenge_time: pulumi.Output[float]
    low_water_mark: pulumi.Output[float]
    low_water_mark_reset: pulumi.Output[float]
    members: pulumi.Output[list]
    netmask: pulumi.Output[float]
    network: pulumi.Output[str]
    network_container: pulumi.Output[str]
    network_view: pulumi.Output[str]
    options: pulumi.Output[list]
    recycle_leases: pulumi.Output[bool]
    restart_if_needed: pulumi.Output[bool]
    update_dns_on_lease_renewal: pulumi.Output[bool]
    use_authority: pulumi.Output[bool]
    use_enable_ddns: pulumi.Output[bool]
    use_enable_dhcp_thresholds: pulumi.Output[bool]
    use_enable_discovery: pulumi.Output[bool]
    use_options: pulumi.Output[bool]
    use_recycle_leases: pulumi.Output[bool]
    def __init__(__self__, resource_name, opts=None, authority=None, auto_create_reversezone=None, comment=None, disable=None, discovery_member=None, enable_ddns=None, enable_dhcp_thresholds=None, enable_discovery=None, high_water_mark=None, high_water_mark_reset=None, ipv4addr=None, lease_scavenge_time=None, low_water_mark=None, low_water_mark_reset=None, members=None, netmask=None, network=None, network_container=None, network_view=None, options=None, recycle_leases=None, restart_if_needed=None, update_dns_on_lease_renewal=None, use_authority=None, use_enable_ddns=None, use_enable_dhcp_thresholds=None, use_enable_discovery=None, use_options=None, use_recycle_leases=None, __name__=None, __opts__=None):
        """
        Create a Network resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['authority'] = authority

        __props__['auto_create_reversezone'] = auto_create_reversezone

        __props__['comment'] = comment

        __props__['disable'] = disable

        __props__['discovery_member'] = discovery_member

        __props__['enable_ddns'] = enable_ddns

        __props__['enable_dhcp_thresholds'] = enable_dhcp_thresholds

        __props__['enable_discovery'] = enable_discovery

        __props__['high_water_mark'] = high_water_mark

        __props__['high_water_mark_reset'] = high_water_mark_reset

        __props__['ipv4addr'] = ipv4addr

        __props__['lease_scavenge_time'] = lease_scavenge_time

        __props__['low_water_mark'] = low_water_mark

        __props__['low_water_mark_reset'] = low_water_mark_reset

        __props__['members'] = members

        __props__['netmask'] = netmask

        if network is None:
            raise TypeError("Missing required property 'network'")
        __props__['network'] = network

        __props__['network_container'] = network_container

        __props__['network_view'] = network_view

        __props__['options'] = options

        __props__['recycle_leases'] = recycle_leases

        __props__['restart_if_needed'] = restart_if_needed

        __props__['update_dns_on_lease_renewal'] = update_dns_on_lease_renewal

        __props__['use_authority'] = use_authority

        __props__['use_enable_ddns'] = use_enable_ddns

        __props__['use_enable_dhcp_thresholds'] = use_enable_dhcp_thresholds

        __props__['use_enable_discovery'] = use_enable_discovery

        __props__['use_options'] = use_options

        __props__['use_recycle_leases'] = use_recycle_leases

        super(Network, __self__).__init__(
            'infoblox:index/network:Network',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

