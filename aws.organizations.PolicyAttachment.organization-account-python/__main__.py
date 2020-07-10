import pulumi
import pulumi_aws as aws

account = aws.organizations.PolicyAttachment("account",
    policy_id=aws_organizations_policy["example"]["id"],
    target_id="123456789012")

