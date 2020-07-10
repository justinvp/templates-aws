import pulumi
import pulumi_aws as aws

unit = aws.organizations.PolicyAttachment("unit",
    policy_id=aws_organizations_policy["example"]["id"],
    target_id=aws_organizations_organizational_unit["example"]["id"])

