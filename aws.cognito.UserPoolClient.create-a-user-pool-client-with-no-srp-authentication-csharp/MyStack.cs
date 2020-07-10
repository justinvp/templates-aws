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
            ExplicitAuthFlows = 
            {
                "ADMIN_NO_SRP_AUTH",
            },
            GenerateSecret = true,
            UserPoolId = pool.Id,
        });
    }

}

