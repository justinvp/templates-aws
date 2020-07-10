import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.autoscaling.Group("bar", {
    availabilityZones: ["us-east-1a"],
    forceDelete: true,
    healthCheckGracePeriod: 300,
    healthCheckType: "ELB",
    launchConfiguration: aws_launch_configuration_foo.name,
    maxSize: 5,
    minSize: 2,
});
const bat = new aws.autoscaling.Policy("bat", {
    adjustmentType: "ChangeInCapacity",
    autoscalingGroupName: bar.name,
    cooldown: 300,
    scalingAdjustment: 4,
});

