import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Vpc("main", {
    cidrBlock: "10.0.0.0/16",
});
const secondaryCidr = new aws.ec2.VpcIpv4CidrBlockAssociation("secondary_cidr", {
    cidrBlock: "172.2.0.0/16",
    vpcId: main.id,
});

