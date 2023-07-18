import * as aws from "@pulumi/aws";
import * as awsconf from "@pulumi/awsconf";
import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let region = config.require("region");
let profile = config.require("profile");

let configurer = new awsconf.Configurer("configurer", {
    profile: profile,
    region: region,
})

configurer.awsProvider.apply(p => {
    let foobar: aws.Provider = p;
    console.log("got provider P", p);
    const bucket = new aws.s3.Bucket("my-bucket-12709-ts", {}, {provider: p});
    console.log("created bucket");
});
