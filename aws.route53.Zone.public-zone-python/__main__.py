import pulumi
import pulumi_aws as aws

primary = aws.route53.Zone("primary")

