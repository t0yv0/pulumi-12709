import * as aws from "@pulumi/aws";
import * as awsconf from "@pulumi/awsconf";
import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let region = config.require("region");
let profile = config.require("profile");

// This dummy is needed to register AWS provider with the Node SDK. Otherwise awsProviderType becomes string, that is
// provider hydrates as a plain URN which is bad.
let dummyBucket = new aws.s3.Bucket("my-control-bucket-12709-ts", {}, {});

let configurer = new awsconf.Configurer("configurer", {
    profile: profile,
    region: region,
})

export const awsProviderType = configurer.awsProvider.apply(p => typeof(p));

let awsProvider = pulumi.referenceProviderResource("aws", configurer.awsProvider);

const bucket = new aws.s3.Bucket("my-bucket-12709-ts", {}, {
    provider: awsProvider,
});
