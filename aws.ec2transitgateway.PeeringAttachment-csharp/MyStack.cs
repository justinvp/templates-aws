using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var local = new Aws.Provider("local", new Aws.ProviderArgs
        {
            Region = "us-east-1",
        });
        var peer = new Aws.Provider("peer", new Aws.ProviderArgs
        {
            Region = "us-west-2",
        });
        var peerRegion = Output.Create(Aws.GetRegion.InvokeAsync());
        var localTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("localTransitGateway", new Aws.Ec2TransitGateway.TransitGatewayArgs
        {
            Tags = 
            {
                { "Name", "Local TGW" },
            },
        }, new CustomResourceOptions
        {
            Provider = aws.Local,
        });
        var peerTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("peerTransitGateway", new Aws.Ec2TransitGateway.TransitGatewayArgs
        {
            Tags = 
            {
                { "Name", "Peer TGW" },
            },
        }, new CustomResourceOptions
        {
            Provider = aws.Peer,
        });
        var example = new Aws.Ec2TransitGateway.PeeringAttachment("example", new Aws.Ec2TransitGateway.PeeringAttachmentArgs
        {
            PeerAccountId = peerTransitGateway.OwnerId,
            PeerRegion = peerRegion.Apply(peerRegion => peerRegion.Name),
            PeerTransitGatewayId = peerTransitGateway.Id,
            TransitGatewayId = localTransitGateway.Id,
            Tags = 
            {
                { "Name", "TGW Peering Requestor" },
            },
        });
    }

}

