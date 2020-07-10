import pulumi
import pulumi_aws as aws

token = aws.ecr.get_authorization_token()

