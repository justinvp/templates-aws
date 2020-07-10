import pulumi
import pulumi_aws as aws

example = aws.cognito.UserPool("example")
main = aws.cognito.UserPoolDomain("main",
    domain="example-domain",
    user_pool_id=example.id)

