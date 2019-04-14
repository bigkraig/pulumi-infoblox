# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class ZoneAuth(pulumi.CustomResource):
    allow_transfers: pulumi.Output[list]
    allow_updates: pulumi.Output[list]
    comment: pulumi.Output[str]
    copy_xfer_to_notify: pulumi.Output[bool]
    disable: pulumi.Output[bool]
    dns_integrity_enable: pulumi.Output[bool]
    dns_integrity_member: pulumi.Output[str]
    external_primaries: pulumi.Output[list]
    external_secondaries: pulumi.Output[list]
    fqdn: pulumi.Output[str]
    grid_primaries: pulumi.Output[list]
    grid_primary_shared_with_ms_parent_delegation: pulumi.Output[bool]
    grid_secondaries: pulumi.Output[list]
    locked: pulumi.Output[bool]
    locked_by: pulumi.Output[str]
    network_view: pulumi.Output[str]
    ns_group: pulumi.Output[str]
    prefix: pulumi.Output[str]
    restart_if_needed: pulumi.Output[bool]
    scavenging_settings: pulumi.Output[dict]
    soa_default_ttl: pulumi.Output[float]
    soa_expire: pulumi.Output[float]
    soa_negative_ttl: pulumi.Output[float]
    soa_refresh: pulumi.Output[float]
    soa_retry: pulumi.Output[float]
    soa_serial_number: pulumi.Output[float]
    use_allow_transfer: pulumi.Output[bool]
    use_check_names_policy: pulumi.Output[bool]
    use_copy_xfer_to_notify: pulumi.Output[bool]
    use_external_primary: pulumi.Output[bool]
    view: pulumi.Output[str]
    zone_format: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, allow_transfers=None, allow_updates=None, comment=None, copy_xfer_to_notify=None, disable=None, dns_integrity_enable=None, dns_integrity_member=None, external_primaries=None, external_secondaries=None, fqdn=None, grid_primaries=None, grid_primary_shared_with_ms_parent_delegation=None, grid_secondaries=None, locked=None, ns_group=None, prefix=None, restart_if_needed=None, scavenging_settings=None, soa_default_ttl=None, soa_expire=None, soa_negative_ttl=None, soa_refresh=None, soa_retry=None, soa_serial_number=None, use_allow_transfer=None, use_check_names_policy=None, use_copy_xfer_to_notify=None, use_external_primary=None, view=None, zone_format=None, __name__=None, __opts__=None):
        """
        Create a ZoneAuth resource with the given unique name, props, and options.
        
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

        __props__['allow_transfers'] = allow_transfers

        __props__['allow_updates'] = allow_updates

        __props__['comment'] = comment

        __props__['copy_xfer_to_notify'] = copy_xfer_to_notify

        __props__['disable'] = disable

        __props__['dns_integrity_enable'] = dns_integrity_enable

        __props__['dns_integrity_member'] = dns_integrity_member

        __props__['external_primaries'] = external_primaries

        __props__['external_secondaries'] = external_secondaries

        if fqdn is None:
            raise TypeError("Missing required property 'fqdn'")
        __props__['fqdn'] = fqdn

        __props__['grid_primaries'] = grid_primaries

        __props__['grid_primary_shared_with_ms_parent_delegation'] = grid_primary_shared_with_ms_parent_delegation

        __props__['grid_secondaries'] = grid_secondaries

        __props__['locked'] = locked

        __props__['ns_group'] = ns_group

        __props__['prefix'] = prefix

        __props__['restart_if_needed'] = restart_if_needed

        __props__['scavenging_settings'] = scavenging_settings

        __props__['soa_default_ttl'] = soa_default_ttl

        __props__['soa_expire'] = soa_expire

        __props__['soa_negative_ttl'] = soa_negative_ttl

        __props__['soa_refresh'] = soa_refresh

        __props__['soa_retry'] = soa_retry

        __props__['soa_serial_number'] = soa_serial_number

        __props__['use_allow_transfer'] = use_allow_transfer

        __props__['use_check_names_policy'] = use_check_names_policy

        __props__['use_copy_xfer_to_notify'] = use_copy_xfer_to_notify

        __props__['use_external_primary'] = use_external_primary

        __props__['view'] = view

        __props__['zone_format'] = zone_format

        __props__['locked_by'] = None
        __props__['network_view'] = None

        super(ZoneAuth, __self__).__init__(
            'infoblox:index/zoneAuth:ZoneAuth',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

