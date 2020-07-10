using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ram.PrincipalAssociation("example", new Aws.Ram.PrincipalAssociationArgs
        {
            Principal = aws_organizations_organization.Example.Arn,
            ResourceShareArn = aws_ram_resource_share.Example.Arn,
        });
    }

}

