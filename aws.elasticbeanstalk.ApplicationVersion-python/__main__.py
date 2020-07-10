import pulumi
import pulumi_aws as aws

default_bucket = aws.s3.Bucket("defaultBucket")
default_bucket_object = aws.s3.BucketObject("defaultBucketObject",
    bucket=default_bucket.id,
    key="beanstalk/go-v1.zip",
    source=pulumi.FileAsset("go-v1.zip"))
default_application = aws.elasticbeanstalk.Application("defaultApplication", description="tf-test-desc")
default_application_version = aws.elasticbeanstalk.ApplicationVersion("defaultApplicationVersion",
    application="tf-test-name",
    bucket=default_bucket.id,
    description="application version",
    key=default_bucket_object.id)

