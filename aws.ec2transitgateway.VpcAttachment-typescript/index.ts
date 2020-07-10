import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.VpcAttachment("example", {
    subnetIds: [aws_subnet_example.id],
    transitGatewayId: aws_ec2_transit_gateway_example.id,
    vpcId: aws_vpc_example.id,
});

