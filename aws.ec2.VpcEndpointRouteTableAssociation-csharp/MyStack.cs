using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2.VpcEndpointRouteTableAssociation("example", new Aws.Ec2.VpcEndpointRouteTableAssociationArgs
        {
            RouteTableId = aws_route_table.Example.Id,
            VpcEndpointId = aws_vpc_endpoint.Example.Id,
        });
    }

}

