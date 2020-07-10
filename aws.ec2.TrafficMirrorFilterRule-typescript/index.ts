import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const filter = new aws.ec2.TrafficMirrorFilter("filter", {
    description: "traffic mirror filter - example",
    networkServices: ["amazon-dns"],
});
const ruleout = new aws.ec2.TrafficMirrorFilterRule("ruleout", {
    description: "test rule",
    destinationCidrBlock: "10.0.0.0/8",
    ruleAction: "accept",
    ruleNumber: 1,
    sourceCidrBlock: "10.0.0.0/8",
    trafficDirection: "egress",
    trafficMirrorFilterId: filter.id,
});
const rulein = new aws.ec2.TrafficMirrorFilterRule("rulein", {
    description: "test rule",
    destinationCidrBlock: "10.0.0.0/8",
    destinationPortRange: {
        fromPort: 22,
        toPort: 53,
    },
    protocol: 6,
    ruleAction: "accept",
    ruleNumber: 1,
    sourceCidrBlock: "10.0.0.0/8",
    sourcePortRange: {
        fromPort: 0,
        toPort: 10,
    },
    trafficDirection: "ingress",
    trafficMirrorFilterId: filter.id,
});

