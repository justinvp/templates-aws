using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var pc = Output.Create(Aws.Ec2.GetVpcPeeringConnection.InvokeAsync(new Aws.Ec2.GetVpcPeeringConnectionArgs
        {
            PeerCidrBlock = "10.0.1.0/22",
            VpcId = aws_vpc.Foo.Id,
        }));
        // Create a route table
        var rt = new Aws.Ec2.RouteTable("rt", new Aws.Ec2.RouteTableArgs
        {
            VpcId = aws_vpc.Foo.Id,
        });
        // Create a route
        var route = new Aws.Ec2.Route("route", new Aws.Ec2.RouteArgs
        {
            DestinationCidrBlock = pc.Apply(pc => pc.PeerCidrBlock),
            RouteTableId = rt.Id,
            VpcPeeringConnectionId = pc.Apply(pc => pc.Id),
        });
    }

}

