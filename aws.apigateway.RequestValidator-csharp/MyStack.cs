using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGateway.RequestValidator("example", new Aws.ApiGateway.RequestValidatorArgs
        {
            RestApi = aws_api_gateway_rest_api.Example.Id,
            ValidateRequestBody = true,
            ValidateRequestParameters = true,
        });
    }

}

