using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.Stage("example", new Aws.ApiGatewayV2.StageArgs
        {
            ApiId = aws_apigatewayv2_api.Example.Id,
        });
    }

}

