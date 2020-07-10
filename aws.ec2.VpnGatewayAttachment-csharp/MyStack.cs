using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var network = new Aws.Ec2.Vpc("network", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.0.0.0/16",
        });
        var vpn = new Aws.Ec2.VpnGateway("vpn", new Aws.Ec2.VpnGatewayArgs
        {
            Tags = 
            {
                { "Name", "example-vpn-gateway" },
            },
        });
        var vpnAttachment = new Aws.Ec2.VpnGatewayAttachment("vpnAttachment", new Aws.Ec2.VpnGatewayAttachmentArgs
        {
            VpcId = network.Id,
            VpnGatewayId = vpn.Id,
        });
    }

}

