// *** WARNING: this file was generated by pulumi-gen-awsconf. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

import * as pulumiAws from "@pulumi/aws";

export class Configurer extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'awsconf:index:Configurer';

    /**
     * Returns true if the given object is an instance of Configurer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Configurer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Configurer.__pulumiType;
    }


    /**
     * Create a Configurer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigurerArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.awsProfile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'awsProfile'");
            }
            if ((!args || args.awsRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'awsRegion'");
            }
            resourceInputs["awsProfile"] = args ? args.awsProfile : undefined;
            resourceInputs["awsRegion"] = args ? args.awsRegion : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
        } else {
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Configurer.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }

    awsProvider(): Promise<pulumiAws.Provider> {
        return pulumi.runtime.callAsync("awsconf:index:Configurer/awsProvider", {
            "__self__": this,
        }, this, {plainResourceField: "resource"});
    }
}

/**
 * The set of arguments for constructing a Configurer resource.
 */
export interface ConfigurerArgs {
    awsProfile: pulumi.Input<string>;
    awsRegion: pulumi.Input<string>;
    mode?: pulumi.Input<string>;
}

export namespace Configurer {
    /**
     * The results of the Configurer.awsProvider method.
     */
    export interface AwsProviderResult {
        readonly resource: pulumiAws.Provider;
    }

}
