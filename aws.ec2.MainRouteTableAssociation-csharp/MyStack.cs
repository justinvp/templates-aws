using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var mainRouteTableAssociation = new Aws.Ec2.MainRouteTableAssociation("mainRouteTableAssociation", new Aws.Ec2.MainRouteTableAssociationArgs
        {
            RouteTableId = aws_route_table.Bar.Id,
            VpcId = aws_vpc.Foo.Id,
        });
    }

}

