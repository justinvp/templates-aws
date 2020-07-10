using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.AppSync.GraphQLApi("example", new Aws.AppSync.GraphQLApiArgs
        {
            AdditionalAuthenticationProviders = 
            {
                new Aws.AppSync.Inputs.GraphQLApiAdditionalAuthenticationProviderArgs
                {
                    AuthenticationType = "AWS_IAM",
                },
            },
            AuthenticationType = "API_KEY",
        });
    }

}

