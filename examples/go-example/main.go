package main

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	awsconf "github.com/t0yv0/pulumi-12709/sdk/go/awsconf"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := config.New(ctx, "")
		profile := conf.Require("profile")
		region := conf.Require("region")

		serverRandomPet, err := random.NewRandomPet(ctx, "serverRandomPet", &random.RandomPetArgs{})
		if err != nil {
			return err
		}

		profileX := serverRandomPet.ID().ApplyT(func(id string) (string, error) {
			fmt.Println("IGNORING", id)
			return profile, nil
		}).(pulumi.StringOutput)

		configurer, err := awsconf.NewConfigurer(ctx, "configurer", &awsconf.ConfigurerArgs{
			AwsRegion:  pulumi.String(region),
			AwsProfile: profileX,
		})
		if err != nil {
			return err
		}

		awsProv, err := configurer.AwsProvider(ctx)
		if err != nil {
			return err
		}

		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucket(ctx, "my-bucket-12709", nil, pulumi.Provider(awsProv))
		if err != nil {
			return err
		}

		ctx.Export("bucketID", bucket.ID())

		return nil
	})
}
