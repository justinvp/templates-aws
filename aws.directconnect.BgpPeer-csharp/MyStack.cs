using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var peer = new Aws.DirectConnect.BgpPeer("peer", new Aws.DirectConnect.BgpPeerArgs
        {
            AddressFamily = "ipv6",
            BgpAsn = 65351,
            VirtualInterfaceId = aws_dx_private_virtual_interface.Foo.Id,
        });
    }

}

