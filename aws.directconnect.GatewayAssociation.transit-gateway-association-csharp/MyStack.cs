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
        var exampleTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("exampleTransitGateway", new Aws.Ec2TransitGateway.TransitGatewayArgs
        {
        });
        var exampleGatewayAssociation = new Aws.DirectConnect.GatewayAssociation("exampleGatewayAssociation", new Aws.DirectConnect.GatewayAssociationArgs
        {
            AllowedPrefixes = 
            {
                "10.255.255.0/30",
                "10.255.255.8/30",
            },
            AssociatedGatewayId = exampleTransitGateway.Id,
            DxGatewayId = exampleGateway.Id,
        });
    }

}

