package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	awsconf "github.com/t0yv0/pulumi-12709/sdk/go/awsconf"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := config.New(ctx, "")
		profile := conf.Require("profile")
		region := conf.Require("region")

		configurer, err := awsconf.NewConfigurer(ctx, "configurer", &awsconf.ConfigurerArgs{
			Profile: pulumi.String(profile),
			Region:  pulumi.String(region),
		})
		if err != nil {
			return err
		}

		configurer.AwsProvider.ApplyT(func(p *aws.Provider) (int, error) {

			// Create an AWS resource (S3 Bucket)
			bucket, err := s3.NewBucket(ctx, "my-bucket-12709", nil, pulumi.Provider(p))
			if err != nil {
				return 0, err
			}

			ctx.Export("bucketID", bucket.ID())

			return 0, nil
		})

		return nil
	})
}
