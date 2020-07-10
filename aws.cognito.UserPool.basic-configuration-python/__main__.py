import pulumi
import pulumi_aws as aws

pool = aws.cognito.UserPool("pool")

