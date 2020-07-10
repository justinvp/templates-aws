using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2TransitGateway.GetRouteTable.InvokeAsync(new Aws.Ec2TransitGateway.GetRouteTableArgs
        {
            Id = "tgw-rtb-12345678",
        }));
    }

}

