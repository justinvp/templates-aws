using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var organizationAccess = new Aws.CloudWatch.EventPermission("organizationAccess", new Aws.CloudWatch.EventPermissionArgs
        {
            Condition = new Aws.CloudWatch.Inputs.EventPermissionConditionArgs
            {
                Key = "aws:PrincipalOrgID",
                Type = "StringEquals",
                Value = aws_organizations_organization.Example.Id,
            },
            Principal = "*",
            StatementId = "OrganizationAccess",
        });
    }

}

