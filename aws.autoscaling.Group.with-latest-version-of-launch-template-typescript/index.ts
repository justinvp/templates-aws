import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foobar = new aws.ec2.LaunchTemplate("foobar", {
    imageId: "ami-1a2b3c",
    instanceType: "t2.micro",
    namePrefix: "foobar",
});
const bar = new aws.autoscaling.Group("bar", {
    availabilityZones: ["us-east-1a"],
    desiredCapacity: 1,
    launchTemplate: {
        id: foobar.id,
        version: "$Latest",
    },
    maxSize: 1,
    minSize: 1,
});

