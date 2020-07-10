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
        var route = Output.Create(Aws.Ec2.GetRoute.InvokeAsync(new Aws.Ec2.GetRouteArgs
        {
            DestinationCidrBlock = "10.0.1.0/24",
            RouteTableId = aws_route_table.Selected.Id,
        }));
        var @interface = Output.Create(Aws.Ec2.GetNetworkInterface.InvokeAsync(new Aws.Ec2.GetNetworkInterfaceArgs
        {
            NetworkInterfaceId = route.Apply(route => route.NetworkInterfaceId),
        }));
    }

}

