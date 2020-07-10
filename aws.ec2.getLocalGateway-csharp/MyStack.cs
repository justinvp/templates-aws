using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var localGatewayId = config.RequireObject<dynamic>("localGatewayId");
        var selected = Output.Create(Aws.Ec2.GetLocalGateway.InvokeAsync(new Aws.Ec2.GetLocalGatewayArgs
        {
            Id = localGatewayId,
        }));
    }

}

