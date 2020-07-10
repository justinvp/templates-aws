using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var root = new Aws.Organizations.PolicyAttachment("root", new Aws.Organizations.PolicyAttachmentArgs
        {
            PolicyId = aws_organizations_policy.Example.Id,
            TargetId = aws_organizations_organization.Example.Roots[0].Id,
        });
    }

}

