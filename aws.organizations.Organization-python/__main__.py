import pulumi
import pulumi_aws as aws

org = aws.organizations.Organization("org",
    aws_service_access_principals=[
        "cloudtrail.amazonaws.com",
        "config.amazonaws.com",
    ],
    feature_set="ALL")

