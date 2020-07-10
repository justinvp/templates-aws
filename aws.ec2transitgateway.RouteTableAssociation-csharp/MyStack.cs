using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2TransitGateway.RouteTableAssociation("example", new Aws.Ec2TransitGateway.RouteTableAssociationArgs
        {
            TransitGatewayAttachmentId = aws_ec2_transit_gateway_vpc_attachment.Example.Id,
            TransitGatewayRouteTableId = aws_ec2_transit_gateway_route_table.Example.Id,
        });
    }

}

