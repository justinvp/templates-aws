using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.DirectConnect.HostedPublicVirtualInterface("foo", new Aws.DirectConnect.HostedPublicVirtualInterfaceArgs
        {
            AddressFamily = "ipv4",
            AmazonAddress = "175.45.176.2/30",
            BgpAsn = 65352,
            ConnectionId = "dxcon-zzzzzzzz",
            CustomerAddress = "175.45.176.1/30",
            RouteFilterPrefixes = 
            {
                "210.52.109.0/24",
                "175.45.176.0/22",
            },
            Vlan = 4094,
        });
    }

}

