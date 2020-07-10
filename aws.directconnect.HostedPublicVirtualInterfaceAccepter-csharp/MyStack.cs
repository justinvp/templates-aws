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
        var creator = new Aws.DirectConnect.HostedPublicVirtualInterface("creator", new Aws.DirectConnect.HostedPublicVirtualInterfaceArgs
        {
            AddressFamily = "ipv4",
            AmazonAddress = "175.45.176.2/30",
            BgpAsn = 65352,
            ConnectionId = "dxcon-zzzzzzzz",
            CustomerAddress = "175.45.176.1/30",
            OwnerAccountId = accepterCallerIdentity.Apply(accepterCallerIdentity => accepterCallerIdentity.AccountId),
            RouteFilterPrefixes = 
            {
                "210.52.109.0/24",
                "175.45.176.0/22",
            },
            Vlan = 4094,
        });
        // Accepter's side of the VIF.
        var accepterHostedPublicVirtualInterfaceAccepter = new Aws.DirectConnect.HostedPublicVirtualInterfaceAccepter("accepterHostedPublicVirtualInterfaceAccepter", new Aws.DirectConnect.HostedPublicVirtualInterfaceAccepterArgs
        {
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

