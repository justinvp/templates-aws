using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("exampleTransitGateway", new Aws.Ec2TransitGateway.TransitGatewayArgs
        {
        });
        var exampleCustomerGateway = new Aws.Ec2.CustomerGateway("exampleCustomerGateway", new Aws.Ec2.CustomerGatewayArgs
        {
            BgpAsn = 65000,
            IpAddress = "172.0.0.1",
            Type = "ipsec.1",
        });
        var exampleVpnConnection = new Aws.Ec2.VpnConnection("exampleVpnConnection", new Aws.Ec2.VpnConnectionArgs
        {
            CustomerGatewayId = exampleCustomerGateway.Id,
            TransitGatewayId = exampleTransitGateway.Id,
            Type = exampleCustomerGateway.Type,
        });
    }

}

