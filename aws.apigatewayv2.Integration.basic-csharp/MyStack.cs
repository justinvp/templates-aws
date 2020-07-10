using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.Integration("example", new Aws.ApiGatewayV2.IntegrationArgs
        {
            ApiId = aws_apigatewayv2_api.Example.Id,
            IntegrationType = "MOCK",
        });
    }

}

