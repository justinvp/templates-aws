import pulumi
import pulumi_aws as aws

cert = aws.acm.Certificate("cert",
    domain_name="example.com",
    tags={
        "Environment": "test",
    },
    validation_method="DNS")

