# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class arecord(pulumi.CustomResource):
    comment: pulumi.Output[str]
    creation_time: pulumi.Output[float]
    creator: pulumi.Output[str]
    ddns_principal: pulumi.Output[str]
    ddns_protected: pulumi.Output[bool]
    disable: pulumi.Output[bool]
    dns_name: pulumi.Output[str]
    forbid_reclamation: pulumi.Output[bool]
    ipv4addr: pulumi.Output[str]
    name: pulumi.Output[str]
    reclaimable: pulumi.Output[bool]
    shared_record_group: pulumi.Output[str]
    ttl: pulumi.Output[float]
    use_ttl: pulumi.Output[bool]
    view: pulumi.Output[str]
    zone: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, comment=None, creator=None, ddns_principal=None, ddns_protected=None, disable=None, dns_name=None, forbid_reclamation=None, ipv4addr=None, name=None, reclaimable=None, shared_record_group=None, ttl=None, use_ttl=None, view=None, __name__=None, __opts__=None):
        """
        Create a arecord resource with the given unique name, props, and options.
        
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

        __props__['creator'] = creator

        __props__['ddns_principal'] = ddns_principal

        __props__['ddns_protected'] = ddns_protected

        __props__['disable'] = disable

        __props__['dns_name'] = dns_name

        __props__['forbid_reclamation'] = forbid_reclamation

        if ipv4addr is None:
            raise TypeError("Missing required property 'ipv4addr'")
        __props__['ipv4addr'] = ipv4addr

        __props__['name'] = name

        __props__['reclaimable'] = reclaimable

        __props__['shared_record_group'] = shared_record_group

        __props__['ttl'] = ttl

        __props__['use_ttl'] = use_ttl

        __props__['view'] = view

        __props__['creation_time'] = None
        __props__['zone'] = None

        super(arecord, __self__).__init__(
            'infoblox:infoblox/arecord:arecord',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
