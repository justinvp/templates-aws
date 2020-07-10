using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleLocalGatewayRouteTable = Output.Create(Aws.Ec2.GetLocalGatewayRouteTable.InvokeAsync(new Aws.Ec2.GetLocalGatewayRouteTableArgs
        {
            OutpostArn = "arn:aws:outposts:us-west-2:123456789012:outpost/op-1234567890abcdef",
        }));
        var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.0.0.0/16",
        });
        var exampleLocalGatewayRouteTableVpcAssociation = new Aws.Ec2.LocalGatewayRouteTableVpcAssociation("exampleLocalGatewayRouteTableVpcAssociation", new Aws.Ec2.LocalGatewayRouteTableVpcAssociationArgs
        {
            LocalGatewayRouteTableId = exampleLocalGatewayRouteTable.Apply(exampleLocalGatewayRouteTable => exampleLocalGatewayRouteTable.Id),
            VpcId = exampleVpc.Id,
        });
    }

}

