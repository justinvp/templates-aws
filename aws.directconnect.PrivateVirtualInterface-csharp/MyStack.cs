using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.DirectConnect.PrivateVirtualInterface("foo", new Aws.DirectConnect.PrivateVirtualInterfaceArgs
        {
            AddressFamily = "ipv4",
            BgpAsn = 65352,
            ConnectionId = "dxcon-zzzzzzzz",
            Vlan = 4094,
        });
    }

}

