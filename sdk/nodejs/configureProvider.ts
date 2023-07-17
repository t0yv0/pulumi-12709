// *** WARNING: this file was generated by pulumi-gen-awsconf. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

import * as pulumiAws from "@pulumi/aws";

export function configureProvider(args?: ConfigureProviderArgs, opts?: pulumi.InvokeOptions): Promise<ConfigureProviderResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("awsconf:index:ConfigureProvider", {
        "profile": args.profile,
        "region": args.region,
    }, opts);
}

export interface ConfigureProviderArgs {
    profile?: string;
    region?: string;
}

export interface ConfigureProviderResult {
    readonly awsProvider?: pulumiAws.Provider;
}
export function configureProviderOutput(args?: ConfigureProviderOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ConfigureProviderResult> {
    return pulumi.output(args).apply((a: any) => configureProvider(a, opts))
}

export interface ConfigureProviderOutputArgs {
    profile?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}
