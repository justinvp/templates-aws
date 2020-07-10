using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Transfer.GetServer.InvokeAsync(new Aws.Transfer.GetServerArgs
        {
            ServerId = "s-1234567",
        }));
    }

}

