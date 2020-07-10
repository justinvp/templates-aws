import pulumi
import pulumi_aws as aws

organization_access = aws.cloudwatch.EventPermission("organizationAccess",
    condition={
        "key": "aws:PrincipalOrgID",
        "type": "StringEquals",
        "value": aws_organizations_organization["example"]["id"],
    },
    principal="*",
    statement_id="OrganizationAccess")

