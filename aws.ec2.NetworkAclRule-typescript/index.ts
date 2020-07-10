import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const barNetworkAcl = new aws.ec2.NetworkAcl("barNetworkAcl", {vpcId: aws_vpc.foo.id});
const barNetworkAclRule = new aws.ec2.NetworkAclRule("barNetworkAclRule", {
    networkAclId: barNetworkAcl.id,
    ruleNumber: 200,
    egress: false,
    protocol: "tcp",
    ruleAction: "allow",
    cidrBlock: aws_vpc.foo.cidr_block,
    fromPort: 22,
    toPort: 22,
});

