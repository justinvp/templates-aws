import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foobarGroup = new aws.autoscaling.Group("foobar", {
    availabilityZones: ["us-west-2a"],
    forceDelete: true,
    healthCheckGracePeriod: 300,
    healthCheckType: "ELB",
    maxSize: 1,
    minSize: 1,
    terminationPolicies: ["OldestInstance"],
});
const foobarSchedule = new aws.autoscaling.Schedule("foobar", {
    autoscalingGroupName: foobarGroup.name,
    desiredCapacity: 0,
    endTime: "2016-12-12T06:00:00Z",
    maxSize: 1,
    minSize: 0,
    scheduledActionName: "foobar",
    startTime: "2016-12-11T18:00:00Z",
});

