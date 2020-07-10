using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.Authorizer("example", new Aws.ApiGatewayV2.AuthorizerArgs
        {
            ApiId = aws_apigatewayv2_api.Example.Id,
            AuthorizerType = "JWT",
            IdentitySources = 
            {
                "$request.header.Authorization",
            },
            JwtConfiguration = new Aws.ApiGatewayV2.Inputs.AuthorizerJwtConfigurationArgs
            {
                Audience = 
                {
                    "example",
                },
                Issuer = $"https://{aws_cognito_user_pool.Example.Endpoint}",
            },
        });
    }

}

