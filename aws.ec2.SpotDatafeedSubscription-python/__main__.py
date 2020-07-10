import pulumi
import pulumi_aws as aws

default_bucket = aws.s3.Bucket("defaultBucket")
default_spot_datafeed_subscription = aws.ec2.SpotDatafeedSubscription("defaultSpotDatafeedSubscription",
    bucket=default_bucket.bucket,
    prefix="my_subdirectory")

