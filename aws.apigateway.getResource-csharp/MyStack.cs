using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myRestApi = Output.Create(Aws.ApiGateway.GetRestApi.InvokeAsync(new Aws.ApiGateway.GetRestApiArgs
        {
            Name = "my-rest-api",
        }));
        var myResource = myRestApi.Apply(myRestApi => Output.Create(Aws.ApiGateway.GetResource.InvokeAsync(new Aws.ApiGateway.GetResourceArgs
        {
            Path = "/endpoint/path",
            RestApiId = myRestApi.Id,
        })));
    }

}

