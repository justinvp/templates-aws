using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGatewayV2.ApiMapping("example", new Aws.ApiGatewayV2.ApiMappingArgs
        {
            ApiId = aws_apigatewayv2_api.Example.Id,
            DomainName = aws_apigatewayv2_domain_name.Example.Id,
            Stage = aws_apigatewayv2_stage.Example.Id,
        });
    }

}

