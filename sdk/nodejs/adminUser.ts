// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AdminUser extends pulumi.CustomResource {
    /**
     * Get an existing AdminUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdminUserState, opts?: pulumi.CustomResourceOptions): AdminUser {
        return new AdminUser(name, <any>state, { ...opts, id: id });
    }

    /**
     * The admin_groups the user belongs to , there can be only 1
     */
    public readonly adminGroups: pulumi.Output<string>;
    /**
     * a comment on the user
     */
    public readonly comment: pulumi.Output<string | undefined>;
    /**
     * Should the user be disabled
     */
    public readonly disable: pulumi.Output<boolean | undefined>;
    /**
     * Email address for the user
     */
    public readonly email: pulumi.Output<string | undefined>;
    /**
     * Name for the user
     */
    public readonly name: pulumi.Output<string>;
    public readonly password: pulumi.Output<string>;

    /**
     * Create a AdminUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AdminUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdminUserArgs | AdminUserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AdminUserState = argsOrState as AdminUserState | undefined;
            inputs["adminGroups"] = state ? state.adminGroups : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["disable"] = state ? state.disable : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
        } else {
            const args = argsOrState as AdminUserArgs | undefined;
            if (!args || args.adminGroups === undefined) {
                throw new Error("Missing required property 'adminGroups'");
            }
            if (!args || args.password === undefined) {
                throw new Error("Missing required property 'password'");
            }
            inputs["adminGroups"] = args ? args.adminGroups : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["disable"] = args ? args.disable : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
        }
        super("infoblox:index/adminUser:AdminUser", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AdminUser resources.
 */
export interface AdminUserState {
    /**
     * The admin_groups the user belongs to , there can be only 1
     */
    readonly adminGroups?: pulumi.Input<string>;
    /**
     * a comment on the user
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * Should the user be disabled
     */
    readonly disable?: pulumi.Input<boolean>;
    /**
     * Email address for the user
     */
    readonly email?: pulumi.Input<string>;
    /**
     * Name for the user
     */
    readonly name?: pulumi.Input<string>;
    readonly password?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AdminUser resource.
 */
export interface AdminUserArgs {
    /**
     * The admin_groups the user belongs to , there can be only 1
     */
    readonly adminGroups: pulumi.Input<string>;
    /**
     * a comment on the user
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * Should the user be disabled
     */
    readonly disable?: pulumi.Input<boolean>;
    /**
     * Email address for the user
     */
    readonly email?: pulumi.Input<string>;
    /**
     * Name for the user
     */
    readonly name?: pulumi.Input<string>;
    readonly password: pulumi.Input<string>;
}