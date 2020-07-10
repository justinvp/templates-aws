using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var unit = new Aws.Organizations.PolicyAttachment("unit", new Aws.Organizations.PolicyAttachmentArgs
        {
            PolicyId = aws_organizations_policy.Example.Id,
            TargetId = aws_organizations_organizational_unit.Example.Id,
        });
    }

}

