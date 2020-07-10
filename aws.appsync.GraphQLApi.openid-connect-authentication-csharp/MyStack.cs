using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.AppSync.GraphQLApi("example", new Aws.AppSync.GraphQLApiArgs
        {
            AuthenticationType = "OPENID_CONNECT",
            OpenidConnectConfig = new Aws.AppSync.Inputs.GraphQLApiOpenidConnectConfigArgs
            {
                Issuer = "https://example.com",
            },
        });
    }

}

