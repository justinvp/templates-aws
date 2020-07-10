import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLocalGatewayRouteTable = aws.ec2.getLocalGatewayRouteTable({
    outpostArn: "arn:aws:outposts:us-west-2:123456789012:outpost/op-1234567890abcdef",
});
const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.0.0.0/16"});
const exampleLocalGatewayRouteTableVpcAssociation = new aws.ec2.LocalGatewayRouteTableVpcAssociation("exampleLocalGatewayRouteTableVpcAssociation", {
    localGatewayRouteTableId: exampleLocalGatewayRouteTable.then(exampleLocalGatewayRouteTable => exampleLocalGatewayRouteTable.id),
    vpcId: exampleVpc.id,
});

