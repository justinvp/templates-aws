import pulumi
import pulumi_aws as aws

example = aws.iam.get_policy(arn="arn:aws:iam::123456789012:policy/UsersManageOwnCredentials")

