import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const gw = new aws.ec2.NatGateway("gw", {
    allocationId: aws_eip_nat.id,
    subnetId: aws_subnet_example.id,
});

