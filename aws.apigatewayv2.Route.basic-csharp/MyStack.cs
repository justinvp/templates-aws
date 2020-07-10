using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.Route("example", new Aws.ApiGatewayV2.RouteArgs
        {
            ApiId = aws_apigatewayv2_api.Example.Id,
            RouteKey = "$default",
        });
    }

}

