using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var routeTable = new Aws.Ec2.RouteTable("routeTable", new Aws.Ec2.RouteTableArgs
        {
            Routes = 
            {
                new Aws.Ec2.Inputs.RouteTableRouteArgs
                {
                    CidrBlock = "10.0.1.0/24",
                    GatewayId = aws_internet_gateway.Main.Id,
                },
                new Aws.Ec2.Inputs.RouteTableRouteArgs
                {
                    EgressOnlyGatewayId = aws_egress_only_internet_gateway.Foo.Id,
                    Ipv6CidrBlock = "::/0",
                },
            },
            Tags = 
            {
                { "Name", "main" },
            },
            VpcId = aws_vpc.Default.Id,
        });
    }

}

