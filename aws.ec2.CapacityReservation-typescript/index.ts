import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultCapacityReservation = new aws.ec2.CapacityReservation("default", {
    availabilityZone: "eu-west-1a",
    instanceCount: 1,
    instancePlatform: "Linux/UNIX",
    instanceType: "t2.micro",
});

