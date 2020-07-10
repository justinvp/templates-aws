using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var vpc = new Aws.Ec2.Vpc("vpc", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.0.0.0/16",
        });
        var vpnGateway = new Aws.Ec2.VpnGateway("vpnGateway", new Aws.Ec2.VpnGatewayArgs
        {
            VpcId = vpc.Id,
        });
        var customerGateway = new Aws.Ec2.CustomerGateway("customerGateway", new Aws.Ec2.CustomerGatewayArgs
        {
            BgpAsn = 65000,
            IpAddress = "172.0.0.1",
            Type = "ipsec.1",
        });
        var main = new Aws.Ec2.VpnConnection("main", new Aws.Ec2.VpnConnectionArgs
        {
            CustomerGatewayId = customerGateway.Id,
            StaticRoutesOnly = true,
            Type = "ipsec.1",
            VpnGatewayId = vpnGateway.Id,
        });
        var office = new Aws.Ec2.VpnConnectionRoute("office", new Aws.Ec2.VpnConnectionRouteArgs
        {
            DestinationCidrBlock = "192.168.10.0/24",
            VpnConnectionId = main.Id,
        });
    }

}

