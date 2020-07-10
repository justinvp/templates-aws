using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var originAccessIdentity = new Aws.CloudFront.OriginAccessIdentity("originAccessIdentity", new Aws.CloudFront.OriginAccessIdentityArgs
        {
            Comment = "Some comment",
        });
    }

}

