import * as aws from "@pulumi/aws";
import * as awsconf from "@pulumi/awsconf";
import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let region = config.require("region");
let profile = config.require("profile");

// This dummy is needed to register AWS provider with the Node SDK. Otherwise awsProviderType becomes string, that is
// provider hydrates as a plain URN which is bad.
// let dummyBucket = new aws.s3.Bucket("my-control-bucket-12709-ts", {}, {});
//
// Looks like this is not a problem if any aws.s3.Bucket is referenced, before or after the provider configurer.

const providers = new awsconf.Configurer("configurer", {
    awsProfile: profile,
    awsRegion: region,
});

const awsProvider = await providers.awsProvider();

const bucket = new aws.s3.Bucket("my-bucket-12709-ts", {}, {
    provider: awsProvider,
});

export const awsProviderExports = {
    typ: typeof(awsProvider),
    urn: awsProvider.urn,
    isOutput: pulumi.Output.isInstance(awsProvider),
    region: awsProvider.region,
};

export const bucketID = bucket.id;
