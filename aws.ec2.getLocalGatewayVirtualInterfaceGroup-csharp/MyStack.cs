using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2.GetLocalGatewayVirtualInterfaceGroup.InvokeAsync(new Aws.Ec2.GetLocalGatewayVirtualInterfaceGroupArgs
        {
            LocalGatewayId = data.Aws_ec2_local_gateway.Example.Id,
        }));
    }

}

