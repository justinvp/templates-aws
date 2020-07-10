import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const secondaryCidr = new aws.ec2.VpcIpv4CidrBlockAssociation("secondary_cidr", {
    cidrBlock: "172.2.0.0/16",
    vpcId: aws_vpc_main.id,
});
const inSecondaryCidr = new aws.ec2.Subnet("in_secondary_cidr", {
    cidrBlock: "172.2.0.0/24",
    vpcId: secondaryCidr.vpcId,
});

