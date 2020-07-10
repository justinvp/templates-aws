import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.NetworkAcl("main", {
    egress: [{
        action: "allow",
        cidrBlock: "10.3.0.0/18",
        fromPort: 443,
        protocol: "tcp",
        ruleNo: 200,
        toPort: 443,
    }],
    ingress: [{
        action: "allow",
        cidrBlock: "10.3.0.0/18",
        fromPort: 80,
        protocol: "tcp",
        ruleNo: 100,
        toPort: 80,
    }],
    tags: {
        Name: "main",
    },
    vpcId: aws_vpc_main.id,
});

