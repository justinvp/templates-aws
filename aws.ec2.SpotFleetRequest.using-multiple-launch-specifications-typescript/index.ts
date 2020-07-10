import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ec2.SpotFleetRequest("foo", {
    iamFleetRole: "arn:aws:iam::12345678:role/spot-fleet",
    launchSpecifications: [
        {
            ami: "ami-d06a90b0",
            availabilityZone: "us-west-2a",
            instanceType: "m1.small",
            keyName: "my-key",
        },
        {
            ami: "ami-d06a90b0",
            availabilityZone: "us-west-2a",
            instanceType: "m5.large",
            keyName: "my-key",
        },
    ],
    spotPrice: "0.005",
    targetCapacity: 2,
    validUntil: "2019-11-04T20:44:20Z",
});

