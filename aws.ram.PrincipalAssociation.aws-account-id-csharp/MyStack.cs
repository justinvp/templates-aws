using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleResourceShare = new Aws.Ram.ResourceShare("exampleResourceShare", new Aws.Ram.ResourceShareArgs
        {
            AllowExternalPrincipals = true,
        });
        var examplePrincipalAssociation = new Aws.Ram.PrincipalAssociation("examplePrincipalAssociation", new Aws.Ram.PrincipalAssociationArgs
        {
            Principal = "111111111111",
            ResourceShareArn = exampleResourceShare.Arn,
        });
    }

}

