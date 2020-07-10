using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var subnetId = config.RequireObject<dynamic>("subnetId");
        var selected = Output.Create(Aws.Ec2.GetRouteTable.InvokeAsync(new Aws.Ec2.GetRouteTableArgs
        {
            SubnetId = subnetId,
        }));
        var route = new Aws.Ec2.Route("route", new Aws.Ec2.RouteArgs
        {
            DestinationCidrBlock = "10.0.1.0/22",
            RouteTableId = selected.Apply(selected => selected.Id),
            VpcPeeringConnectionId = "pcx-45ff3dc1",
        });
    }

}

