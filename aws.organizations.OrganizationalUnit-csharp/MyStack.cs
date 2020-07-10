using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Organizations.OrganizationalUnit("example", new Aws.Organizations.OrganizationalUnitArgs
        {
            ParentId = aws_organizations_organization.Example.Roots[0].Id,
        });
    }

}

