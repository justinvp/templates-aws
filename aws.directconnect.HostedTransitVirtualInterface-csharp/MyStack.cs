using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DirectConnect.HostedTransitVirtualInterface("example", new Aws.DirectConnect.HostedTransitVirtualInterfaceArgs
        {
            AddressFamily = "ipv4",
            BgpAsn = 65352,
            ConnectionId = aws_dx_connection.Example.Id,
            Vlan = 4094,
        });
    }

}

