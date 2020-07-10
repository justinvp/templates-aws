using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2TransitGateway.TransitGateway("example", new Aws.Ec2TransitGateway.TransitGatewayArgs
        {
            Description = "example",
        });
    }

}

