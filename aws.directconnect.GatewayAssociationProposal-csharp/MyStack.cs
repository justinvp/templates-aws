using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DirectConnect.GatewayAssociationProposal("example", new Aws.DirectConnect.GatewayAssociationProposalArgs
        {
            AssociatedGatewayId = aws_vpn_gateway.Example.Id,
            DxGatewayId = aws_dx_gateway.Example.Id,
            DxGatewayOwnerAccountId = aws_dx_gateway.Example.Owner_account_id,
        });
    }

}

