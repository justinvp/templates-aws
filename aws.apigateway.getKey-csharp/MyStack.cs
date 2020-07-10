using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myApiKey = Output.Create(Aws.ApiGateway.GetKey.InvokeAsync(new Aws.ApiGateway.GetKeyArgs
        {
            Id = "ru3mpjgse6",
        }));
    }

}

