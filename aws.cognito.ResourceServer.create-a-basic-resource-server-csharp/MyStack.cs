using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var pool = new Aws.Cognito.UserPool("pool", new Aws.Cognito.UserPoolArgs
        {
        });
        var resource = new Aws.Cognito.ResourceServer("resource", new Aws.Cognito.ResourceServerArgs
        {
            Identifier = "https://example.com",
            UserPoolId = pool.Id,
        });
    }

}

