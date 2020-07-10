import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ec2.NetworkInterface("test", {
    attachments: [{
        deviceIndex: 1,
        instance: aws_instance_test.id,
    }],
    privateIps: ["10.0.0.50"],
    securityGroups: [aws_security_group_web.id],
    subnetId: aws_subnet_public_a.id,
});

