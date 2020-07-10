import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const gw = new aws.ec2.InternetGateway("gw", {
    tags: {
        Name: "main",
    },
    vpcId: aws_vpc_main.id,
});

