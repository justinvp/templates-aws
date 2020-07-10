using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.Deployment("example", new Aws.ApiGatewayV2.DeploymentArgs
        {
            ApiId = aws_apigatewayv2_route.Example.Api_id,
            Description = "Example deployment",
        });
    }

}

