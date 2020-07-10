using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ApiGateway.RestApi("example", new Aws.ApiGateway.RestApiArgs
        {
            EndpointConfiguration = new Aws.ApiGateway.Inputs.RestApiEndpointConfigurationArgs
            {
                Types = "REGIONAL",
            },
        });
    }

}

