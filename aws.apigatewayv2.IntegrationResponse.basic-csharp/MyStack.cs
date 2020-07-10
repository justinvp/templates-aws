using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.IntegrationResponse("example", new Aws.ApiGatewayV2.IntegrationResponseArgs
        {
            ApiId = aws_apigatewayv2_api.Example.Id,
            IntegrationId = aws_apigatewayv2_integration.Example.Id,
            IntegrationResponseKey = "/200/",
        });
    }

}

