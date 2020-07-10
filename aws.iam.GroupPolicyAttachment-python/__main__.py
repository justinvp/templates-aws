import pulumi
import pulumi_aws as aws

group = aws.iam.Group("group")
policy = aws.iam.Policy("policy",
    description="A test policy",
    policy="")
# insert policy here
test_attach = aws.iam.GroupPolicyAttachment("test-attach",
    group=group.name,
    policy_arn=policy.arn)

