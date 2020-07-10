import pulumi
import pulumi_aws as aws

example = aws.organizations.OrganizationalUnit("example", parent_id=aws_organizations_organization["example"]["roots"][0]["id"])

