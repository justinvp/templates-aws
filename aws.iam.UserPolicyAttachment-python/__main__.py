import pulumi
import pulumi_aws as aws

user = aws.iam.User("user")
policy = aws.iam.Policy("policy",
    description="A test policy",
    policy="")
# insert policy here
test_attach = aws.iam.UserPolicyAttachment("test-attach",
    policy_arn=policy.arn,
    user=user.name)

