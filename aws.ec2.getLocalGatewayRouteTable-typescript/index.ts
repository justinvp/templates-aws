import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const awsEc2LocalGatewayRouteTable = config.require("awsEc2LocalGatewayRouteTable");

const selected = pulumi.output(aws.ec2.getLocalGatewayRouteTable({
    localGatewayRouteTableId: awsEc2LocalGatewayRouteTable,
}, { async: true }));

