using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = Output.Create(Aws.Ec2.GetNetworkInterface.InvokeAsync(new Aws.Ec2.GetNetworkInterfaceArgs
        {
            Id = "eni-01234567",
        }));
    }

}

