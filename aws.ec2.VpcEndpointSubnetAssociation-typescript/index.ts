import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const snEc2 = new aws.ec2.VpcEndpointSubnetAssociation("sn_ec2", {
    subnetId: aws_subnet_sn.id,
    vpcEndpointId: aws_vpc_endpoint_ec2.id,
});

