using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var account = new Aws.Organizations.PolicyAttachment("account", new Aws.Organizations.PolicyAttachmentArgs
        {
            PolicyId = aws_organizations_policy.Example.Id,
            TargetId = "123456789012",
        });
    }

}

