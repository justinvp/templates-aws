using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2.LocalGatewayRoute("example", new Aws.Ec2.LocalGatewayRouteArgs
        {
            DestinationCidrBlock = "172.16.0.0/16",
            LocalGatewayRouteTableId = data.Aws_ec2_local_gateway_route_table.Example.Id,
            LocalGatewayVirtualInterfaceGroupId = data.Aws_ec2_local_gateway_virtual_interface_group.Example.Id,
        });
    }

}

