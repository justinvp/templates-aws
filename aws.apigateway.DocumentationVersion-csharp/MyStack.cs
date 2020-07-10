using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleRestApi = new Aws.ApiGateway.RestApi("exampleRestApi", new Aws.ApiGateway.RestApiArgs
        {
        });
        var exampleDocumentationVersion = new Aws.ApiGateway.DocumentationVersion("exampleDocumentationVersion", new Aws.ApiGateway.DocumentationVersionArgs
        {
            Description = "Example description",
            RestApiId = exampleRestApi.Id,
            Version = "example_version",
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_api_gateway_documentation_part.example",
            },
        });
        var exampleDocumentationPart = new Aws.ApiGateway.DocumentationPart("exampleDocumentationPart", new Aws.ApiGateway.DocumentationPartArgs
        {
            Location = new Aws.ApiGateway.Inputs.DocumentationPartLocationArgs
            {
                Type = "API",
            },
            Properties = "{\"description\":\"Example\"}",
            RestApiId = exampleRestApi.Id,
        });
    }

}

