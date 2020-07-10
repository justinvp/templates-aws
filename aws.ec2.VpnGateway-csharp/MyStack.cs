using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var vpnGw = new Aws.Ec2.VpnGateway("vpnGw", new Aws.Ec2.VpnGatewayArgs
        {
            Tags = 
            {
                { "Name", "main" },
            },
            VpcId = aws_vpc.Main.Id,
        });
    }

}

