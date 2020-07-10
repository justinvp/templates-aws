using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ec2TransitGateway.GetTransitGateway.InvokeAsync(new Aws.Ec2TransitGateway.GetTransitGatewayArgs
        {
            Id = "tgw-12345678",
        }));
    }

}

