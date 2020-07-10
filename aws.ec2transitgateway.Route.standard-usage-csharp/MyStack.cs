using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2TransitGateway.Route("example", new Aws.Ec2TransitGateway.RouteArgs
        {
            DestinationCidrBlock = "0.0.0.0/0",
            TransitGatewayAttachmentId = aws_ec2_transit_gateway_vpc_attachment.Example.Id,
            TransitGatewayRouteTableId = aws_ec2_transit_gateway.Example.Association_default_route_table_id,
        });
    }

}

