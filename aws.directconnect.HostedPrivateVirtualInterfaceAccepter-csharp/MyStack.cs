using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var accepter = new Aws.Provider("accepter", new Aws.ProviderArgs
        {
        });
        var accepterCallerIdentity = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
        // Creator's side of the VIF
        var creator = new Aws.DirectConnect.HostedPrivateVirtualInterface("creator", new Aws.DirectConnect.HostedPrivateVirtualInterfaceArgs
        {
            AddressFamily = "ipv4",
            BgpAsn = 65352,
            ConnectionId = "dxcon-zzzzzzzz",
            OwnerAccountId = accepterCallerIdentity.Apply(accepterCallerIdentity => accepterCallerIdentity.AccountId),
            Vlan = 4094,
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_vpn_gateway.vpn_gw",
            },
        });
        // Accepter's side of the VIF.
        var vpnGw = new Aws.Ec2.VpnGateway("vpnGw", new Aws.Ec2.VpnGatewayArgs
        {
        }, new CustomResourceOptions
        {
            Provider = "aws.accepter",
        });
        var accepterHostedPrivateVirtualInterfaceAccepter = new Aws.DirectConnect.HostedPrivateVirtualInterfaceAccepter("accepterHostedPrivateVirtualInterfaceAccepter", new Aws.DirectConnect.HostedPrivateVirtualInterfaceAccepterArgs
        {
            Tags = 
            {
                { "Side", "Accepter" },
            },
            VirtualInterfaceId = creator.Id,
            VpnGatewayId = vpnGw.Id,
        }, new CustomResourceOptions
        {
            Provider = "aws.accepter",
        });
    }

}

