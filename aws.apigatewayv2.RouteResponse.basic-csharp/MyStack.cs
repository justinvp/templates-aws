using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.RouteResponse("example", new Aws.ApiGatewayV2.RouteResponseArgs
        {
            ApiId = aws_apigatewayv2_api.Example.Id,
            RouteId = aws_apigatewayv2_route.Example.Id,
            RouteResponseKey = "$default",
        });
    }

}

