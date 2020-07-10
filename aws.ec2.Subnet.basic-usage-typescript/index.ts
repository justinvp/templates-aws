import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Subnet("main", {
    cidrBlock: "10.0.1.0/24",
    tags: {
        Name: "Main",
    },
    vpcId: aws_vpc_main.id,
});

