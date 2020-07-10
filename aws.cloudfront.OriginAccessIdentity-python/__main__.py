import pulumi
import pulumi_aws as aws

origin_access_identity = aws.cloudfront.OriginAccessIdentity("originAccessIdentity", comment="Some comment")

