import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const allowTls = new aws.ec2.SecurityGroup("allowTls", {
    description: "Allow TLS inbound traffic",
    vpcId: aws_vpc.main.id,
    ingress: [{
        description: "TLS from VPC",
        fromPort: 443,
        toPort: 443,
        protocol: "tcp",
        cidrBlocks: [aws_vpc.main.cidr_block],
    }],
    egress: [{
        fromPort: 0,
        toPort: 0,
        protocol: "-1",
        cidrBlocks: ["0.0.0.0/0"],
    }],
    tags: {
        Name: "allow_tls",
    },
});

