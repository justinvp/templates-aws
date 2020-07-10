import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.AvailabilityZoneGroup("example", {
    groupName: "us-west-2-lax-1",
    optInStatus: "opted-in",
});

