import pulumi
import pulumi_awsconf
import pulumi_aws

config = pulumi.Config()
region = config.require('region')
profile = config.require('profile')

providers = pulumi_awsconf.Configurer("configurer",
    aws_region=region, aws_profile=profile)

aws_provider = providers.aws_provider()

bucket = pulumi_aws.s3.Bucket(
    'my-bucket-12709-py',
    opts=pulumi.ResourceOptions(provider=aws_provider))

pulumi.export("bucket_id", bucket.id)
pulumi.export("region", aws_provider.region)
