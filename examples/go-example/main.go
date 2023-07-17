package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	awsconf "github.com/t0yv0/pulumi-12709/sdk/go/awsconf"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf, err := awsconf.NewConfigurer(ctx, "configurer", &awsconf.ConfigurerArgs{
			Profile: pulumi.String("myawsprofile"),
			Region:  pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}

		conf.AwsProvider.ApplyT(func(p *aws.Provider) (int, error) {

			// Create an AWS resource (S3 Bucket)
			bucket, err := s3.NewBucket(ctx, "my-bucket", nil)
			if err != nil {
				return 0, err
			}

			ctx.Export("bucketID", bucket.ID())

			return 0, nil
		})

		return nil
	})
}
