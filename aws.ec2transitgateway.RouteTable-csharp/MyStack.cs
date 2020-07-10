using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2TransitGateway.RouteTable("example", new Aws.Ec2TransitGateway.RouteTableArgs
        {
            TransitGatewayId = aws_ec2_transit_gateway.Example.Id,
        });
    }

}

