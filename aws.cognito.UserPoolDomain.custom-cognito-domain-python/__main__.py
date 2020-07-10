import pulumi
import pulumi_aws as aws

example = aws.cognito.UserPool("example")
main = aws.cognito.UserPoolDomain("main",
    certificate_arn=aws_acm_certificate["cert"]["arn"],
    domain="example-domain.example.com",
    user_pool_id=example.id)

