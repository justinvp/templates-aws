using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var awsEc2LocalGatewayRouteTable = config.RequireObject<dynamic>("awsEc2LocalGatewayRouteTable");
        var selected = Output.Create(Aws.Ec2.GetLocalGatewayRouteTable.InvokeAsync(new Aws.Ec2.GetLocalGatewayRouteTableArgs
        {
            LocalGatewayRouteTableId = awsEc2LocalGatewayRouteTable,
        }));
    }

}

