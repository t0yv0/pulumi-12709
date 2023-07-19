import * as aws from "@pulumi/aws";
import * as awsconf from "@pulumi/awsconf";
import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let region = config.require("region");
let profile = config.require("profile");

// This dummy is needed to register AWS provider with the Node SDK. Otherwise awsProviderType becomes string, that is
// provider hydrates as a plain URN which is bad.
let dummyBucket = new aws.s3.Bucket("my-control-bucket-12709-ts", {}, {});

let providers = await new awsconf.Configurer("configurer", {}).configureAwsProviderAsync({
    profile: profile,
    region: region,
    mode: "normal",
});

const bucket = new aws.s3.Bucket("my-bucket-12709-ts", {}, {
    provider: providers.awsProvider,
});

export const awsProvider = {
    typ: typeof(providers.awsProvider),
    urn: providers.awsProvider.urn,
    isOutput: pulumi.Output.isInstance(providers.awsProvider)
};

export const bucketID = bucket.id;
export const someString = providers.someString;
export const someStringType = typeof(providers.someString);
