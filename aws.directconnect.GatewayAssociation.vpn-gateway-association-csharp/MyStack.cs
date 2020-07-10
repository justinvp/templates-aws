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
        var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.255.255.0/28",
        });
        var exampleVpnGateway = new Aws.Ec2.VpnGateway("exampleVpnGateway", new Aws.Ec2.VpnGatewayArgs
        {
            VpcId = exampleVpc.Id,
        });
        var exampleGatewayAssociation = new Aws.DirectConnect.GatewayAssociation("exampleGatewayAssociation", new Aws.DirectConnect.GatewayAssociationArgs
        {
            AssociatedGatewayId = exampleVpnGateway.Id,
            DxGatewayId = exampleGateway.Id,
        });
    }

}

