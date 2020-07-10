using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2.VpnGatewayRoutePropagation("example", new Aws.Ec2.VpnGatewayRoutePropagationArgs
        {
            RouteTableId = aws_route_table.Example.Id,
            VpnGatewayId = aws_vpn_gateway.Example.Id,
        });
    }

}

