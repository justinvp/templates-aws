using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.Api("example", new Aws.ApiGatewayV2.ApiArgs
        {
            ProtocolType = "WEBSOCKET",
            RouteSelectionExpression = "$request.body.action",
        });
    }

}

