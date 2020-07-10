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
    }

}

