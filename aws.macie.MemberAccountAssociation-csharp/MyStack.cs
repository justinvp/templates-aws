using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Macie.MemberAccountAssociation("example", new Aws.Macie.MemberAccountAssociationArgs
        {
            MemberAccountId = "123456789012",
        });
    }

}

