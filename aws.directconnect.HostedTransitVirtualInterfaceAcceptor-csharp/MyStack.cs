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
        var creator = new Aws.DirectConnect.HostedTransitVirtualInterface("creator", new Aws.DirectConnect.HostedTransitVirtualInterfaceArgs
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
                "aws_dx_gateway.example",
            },
        });
        // Accepter's side of the VIF.
        var example = new Aws.DirectConnect.Gateway("example", new Aws.DirectConnect.GatewayArgs
        {
            AmazonSideAsn = "64512",
        }, new CustomResourceOptions
        {
            Provider = "aws.accepter",
        });
        var accepterHostedTransitVirtualInterfaceAcceptor = new Aws.DirectConnect.HostedTransitVirtualInterfaceAcceptor("accepterHostedTransitVirtualInterfaceAcceptor", new Aws.DirectConnect.HostedTransitVirtualInterfaceAcceptorArgs
        {
            DxGatewayId = example.Id,
            Tags = 
            {
                { "Side", "Accepter" },
            },
            VirtualInterfaceId = creator.Id,
        }, new CustomResourceOptions
        {
            Provider = "aws.accepter",
        });
    }

}

