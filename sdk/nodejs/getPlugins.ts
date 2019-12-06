// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getPlugins(args: GetPluginsArgs, opts?: pulumi.InvokeOptions): Promise<GetPluginsResult> & GetPluginsResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetPluginsResult> = pulumi.runtime.invoke("cloudamqp:index/getPlugins:getPlugins", {
        "instanceId": args.instanceId,
        "plugins": args.plugins,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getPlugins.
 */
export interface GetPluginsArgs {
    readonly instanceId: number;
    readonly plugins?: inputs.GetPluginsPlugin[];
}

/**
 * A collection of values returned by getPlugins.
 */
export interface GetPluginsResult {
    readonly instanceId: number;
    readonly plugins?: outputs.GetPluginsPlugin[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}