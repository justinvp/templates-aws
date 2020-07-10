import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const filter = new aws.ec2.TrafficMirrorFilter("filter", {
    description: "traffic mirror filter - example",
    networkServices: ["amazon-dns"],
});
const target = new aws.ec2.TrafficMirrorTarget("target", {
    networkLoadBalancerArn: aws_lb_lb.arn,
});
const session = new aws.ec2.TrafficMirrorSession("session", {
    description: "traffic mirror session - example",
    networkInterfaceId: aws_instance_test.primaryNetworkInterfaceId,
    trafficMirrorFilterId: filter.id,
    trafficMirrorTargetId: target.id,
});

