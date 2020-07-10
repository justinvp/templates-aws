using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.Authorizer("example", new Aws.ApiGatewayV2.AuthorizerArgs
        {
            ApiId = aws_apigatewayv2_api.Example.Id,
            AuthorizerType = "REQUEST",
            AuthorizerUri = aws_lambda_function.Example.Invoke_arn,
            IdentitySources = 
            {
                "route.request.header.Auth",
            },
        });
    }

}

