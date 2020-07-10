using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleRestApi = new Aws.ApiGateway.RestApi("exampleRestApi", new Aws.ApiGateway.RestApiArgs
        {
        });
        var exampleDocumentationPart = new Aws.ApiGateway.DocumentationPart("exampleDocumentationPart", new Aws.ApiGateway.DocumentationPartArgs
        {
            Location = new Aws.ApiGateway.Inputs.DocumentationPartLocationArgs
            {
                Method = "GET",
                Path = "/example",
                Type = "METHOD",
            },
            Properties = "{\"description\":\"Example description\"}",
            RestApiId = exampleRestApi.Id,
        });
    }

}

