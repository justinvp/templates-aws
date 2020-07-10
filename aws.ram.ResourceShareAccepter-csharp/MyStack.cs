using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var alternate = new Aws.Provider("alternate", new Aws.ProviderArgs
        {
            Profile = "profile1",
        });
        var senderShare = new Aws.Ram.ResourceShare("senderShare", new Aws.Ram.ResourceShareArgs
        {
            AllowExternalPrincipals = true,
            Tags = 
            {
                { "Name", "tf-test-resource-share" },
            },
        }, new CustomResourceOptions
        {
            Provider = "aws.alternate",
        });
        var receiver = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
        var senderInvite = new Aws.Ram.PrincipalAssociation("senderInvite", new Aws.Ram.PrincipalAssociationArgs
        {
            Principal = receiver.Apply(receiver => receiver.AccountId),
            ResourceShareArn = senderShare.Arn,
        }, new CustomResourceOptions
        {
            Provider = "aws.alternate",
        });
        var receiverAccept = new Aws.Ram.ResourceShareAccepter("receiverAccept", new Aws.Ram.ResourceShareAccepterArgs
        {
            ShareArn = senderInvite.ResourceShareArn,
        });
    }

}

