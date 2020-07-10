using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var subnetId = config.RequireObject<dynamic>("subnetId");
        var @default = Output.Create(Aws.Ec2.GetNatGateway.InvokeAsync(new Aws.Ec2.GetNatGatewayArgs
        {
            SubnetId = aws_subnet.Public.Id,
        }));
    }

}

