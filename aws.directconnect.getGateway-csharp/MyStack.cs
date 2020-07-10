using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.DirectConnect.GetGateway.InvokeAsync(new Aws.DirectConnect.GetGatewayArgs
        {
            Name = "example",
        }));
    }

}

