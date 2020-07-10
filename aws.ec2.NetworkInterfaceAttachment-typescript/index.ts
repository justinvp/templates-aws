import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ec2.NetworkInterfaceAttachment("test", {
    deviceIndex: 0,
    instanceId: aws_instance_test.id,
    networkInterfaceId: aws_network_interface_test.id,
});

