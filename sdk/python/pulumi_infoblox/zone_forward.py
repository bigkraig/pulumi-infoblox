# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class ZoneForward(pulumi.CustomResource):
    address: pulumi.Output[str]
    comment: pulumi.Output[str]
    disable: pulumi.Output[bool]
    display_domain: pulumi.Output[str]
    dns_fqdn: pulumi.Output[str]
    forward_tos: pulumi.Output[list]
    forwarders_only: pulumi.Output[bool]
    forwarding_servers: pulumi.Output[list]
    fqdn: pulumi.Output[str]
    locked: pulumi.Output[bool]
    locked_by: pulumi.Output[str]
    mask_prefix: pulumi.Output[str]
    ns_group: pulumi.Output[str]
    parent: pulumi.Output[str]
    prefix: pulumi.Output[str]
    using_srg_associations: pulumi.Output[bool]
    view: pulumi.Output[str]
    zone_format: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, comment=None, disable=None, display_domain=None, dns_fqdn=None, forward_tos=None, forwarders_only=None, forwarding_servers=None, fqdn=None, locked=None, locked_by=None, mask_prefix=None, ns_group=None, parent=None, prefix=None, using_srg_associations=None, view=None, zone_format=None, __name__=None, __opts__=None):
        """
        Create a ZoneForward resource with the given unique name, props, and options.
        
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

        __props__['comment'] = comment

        __props__['disable'] = disable

        __props__['display_domain'] = display_domain

        __props__['dns_fqdn'] = dns_fqdn

        if forward_tos is None:
            raise TypeError("Missing required property 'forward_tos'")
        __props__['forward_tos'] = forward_tos

        __props__['forwarders_only'] = forwarders_only

        __props__['forwarding_servers'] = forwarding_servers

        __props__['fqdn'] = fqdn

        __props__['locked'] = locked

        __props__['locked_by'] = locked_by

        __props__['mask_prefix'] = mask_prefix

        __props__['ns_group'] = ns_group

        __props__['parent'] = parent

        __props__['prefix'] = prefix

        __props__['using_srg_associations'] = using_srg_associations

        __props__['view'] = view

        __props__['zone_format'] = zone_format

        __props__['address'] = None

        super(ZoneForward, __self__).__init__(
            'infoblox:index/zoneForward:ZoneForward',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
