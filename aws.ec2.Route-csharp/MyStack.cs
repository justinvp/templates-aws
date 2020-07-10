using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var route = new Aws.Ec2.Route("route", new Aws.Ec2.RouteArgs
        {
            RouteTableId = "rtb-4fbb3ac4",
            DestinationCidrBlock = "10.0.1.0/22",
            VpcPeeringConnectionId = "pcx-45ff3dc1",
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_route_table.testing",
            },
        });
    }

}

