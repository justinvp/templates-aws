using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var routeTableAssociation = new Aws.Ec2.RouteTableAssociation("routeTableAssociation", new Aws.Ec2.RouteTableAssociationArgs
        {
            SubnetId = aws_subnet.Foo.Id,
            RouteTableId = aws_route_table.Bar.Id,
        });
    }

}

