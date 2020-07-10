using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleGateway = new Aws.DirectConnect.Gateway("exampleGateway", new Aws.DirectConnect.GatewayArgs
        {
            AmazonSideAsn = "64512",
        });
        var exampleTransitVirtualInterface = new Aws.DirectConnect.TransitVirtualInterface("exampleTransitVirtualInterface", new Aws.DirectConnect.TransitVirtualInterfaceArgs
        {
            AddressFamily = "ipv4",
            BgpAsn = 65352,
            ConnectionId = aws_dx_connection.Example.Id,
            DxGatewayId = exampleGateway.Id,
            Vlan = 4094,
        });
    }

}

