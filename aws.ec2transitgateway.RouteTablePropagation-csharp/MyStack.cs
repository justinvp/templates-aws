using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2TransitGateway.RouteTablePropagation("example", new Aws.Ec2TransitGateway.RouteTablePropagationArgs
        {
            TransitGatewayAttachmentId = aws_ec2_transit_gateway_vpc_attachment.Example.Id,
            TransitGatewayRouteTableId = aws_ec2_transit_gateway_route_table.Example.Id,
        });
    }

}

