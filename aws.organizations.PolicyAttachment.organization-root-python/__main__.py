import pulumi
import pulumi_aws as aws

root = aws.organizations.PolicyAttachment("root",
    policy_id=aws_organizations_policy["example"]["id"],
    target_id=aws_organizations_organization["example"]["roots"][0]["id"])

