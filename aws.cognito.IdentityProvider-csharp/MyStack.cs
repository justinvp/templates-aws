using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Cognito.UserPool("example", new Aws.Cognito.UserPoolArgs
        {
            AutoVerifiedAttributes = 
            {
                "email",
            },
        });
        var exampleProvider = new Aws.Cognito.IdentityProvider("exampleProvider", new Aws.Cognito.IdentityProviderArgs
        {
            AttributeMapping = 
            {
                { "email", "email" },
                { "username", "sub" },
            },
            ProviderDetails = 
            {
                { "authorize_scopes", "email" },
                { "client_id", "your client_id" },
                { "client_secret", "your client_secret" },
            },
            ProviderName = "Google",
            ProviderType = "Google",
            UserPoolId = example.Id,
        });
    }

}

