import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Request a spot instance at $0.03
const cheapWorker = new aws.ec2.SpotInstanceRequest("cheap_worker", {
    ami: "ami-1234",
    instanceType: "c4.xlarge",
    spotPrice: "0.03",
    tags: {
        Name: "CheapWorker",
    },
});

