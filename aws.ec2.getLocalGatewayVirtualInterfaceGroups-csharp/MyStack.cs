using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var all = Output.Create(Aws.Ec2.GetLocalGatewayVirtualInterfaceGroups.InvokeAsync());
    }

}

