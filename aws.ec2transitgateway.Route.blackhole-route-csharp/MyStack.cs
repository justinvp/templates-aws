using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2TransitGateway.Route("example", new Aws.Ec2TransitGateway.RouteArgs
        {
            Blackhole = true,
            DestinationCidrBlock = "0.0.0.0/0",
            TransitGatewayRouteTableId = aws_ec2_transit_gateway.Example.Association_default_route_table_id,
        });
    }

}

