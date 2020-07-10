using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var pool = new Aws.Cognito.UserPool("pool", new Aws.Cognito.UserPoolArgs
        {
        });
        var client = new Aws.Cognito.UserPoolClient("client", new Aws.Cognito.UserPoolClientArgs
        {
            UserPoolId = pool.Id,
        });
    }

}

