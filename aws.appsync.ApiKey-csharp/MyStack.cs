using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleGraphQLApi = new Aws.AppSync.GraphQLApi("exampleGraphQLApi", new Aws.AppSync.GraphQLApiArgs
        {
            AuthenticationType = "API_KEY",
        });
        var exampleApiKey = new Aws.AppSync.ApiKey("exampleApiKey", new Aws.AppSync.ApiKeyArgs
        {
            ApiId = exampleGraphQLApi.Id,
            Expires = "2018-05-03T04:00:00Z",
        });
    }

}

