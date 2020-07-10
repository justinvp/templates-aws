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
            Scopes = 
            {
                new Aws.Cognito.Inputs.ResourceServerScopeArgs
                {
                    ScopeDescription = "a Sample Scope Description",
                    ScopeName = "sample-scope",
                },
            },
            UserPoolId = pool.Id,
        });
    }

}

