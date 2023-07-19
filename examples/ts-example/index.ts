import * as aws from "@pulumi/aws";
import * as awsconf from "@pulumi/awsconf";
import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let region = config.require("region");
let profile = config.require("profile");

// This dummy is needed to register AWS provider with the Node SDK. Otherwise awsProviderType becomes string, that is
// provider hydrates as a plain URN which is bad.
let dummyBucket = new aws.s3.Bucket("my-control-bucket-12709-ts", {}, {});

let configurer = new awsconf.Configurer("configurer", {})

let configuredAwsProvider = configurer.configureAwsProvider({
    profile: profile,
    region: region,
}).apply(p => p.awsProvider);

let awsProvider = pulumi.referenceProviderResource("aws", configuredAwsProvider);

export const awsProviderType = configuredAwsProvider.apply(pr => typeof(pr));

const bucket = new aws.s3.Bucket("my-bucket-12709-ts", {}, {
    provider: awsProvider,
});

export const bucketID = bucket.id;
