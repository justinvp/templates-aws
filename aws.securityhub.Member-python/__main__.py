import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
example_member = aws.securityhub.Member("exampleMember",
    account_id="123456789012",
    email="example@example.com",
    invite=True,
    opts=ResourceOptions(depends_on=["aws_securityhub_account.example"]))

