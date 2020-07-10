using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.DirectConnect.HostedPrivateVirtualInterface("foo", new Aws.DirectConnect.HostedPrivateVirtualInterfaceArgs
        {
            AddressFamily = "ipv4",
            BgpAsn = 65352,
            ConnectionId = "dxcon-zzzzzzzz",
            Vlan = 4094,
        });
    }

}

