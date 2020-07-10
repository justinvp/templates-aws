import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.SecurityGroupRule("example", {
    type: "ingress",
    fromPort: 0,
    toPort: 65535,
    protocol: "tcp",
    cidrBlocks: aws_vpc.example.cidr_block,
    securityGroupId: "sg-123456",
});

