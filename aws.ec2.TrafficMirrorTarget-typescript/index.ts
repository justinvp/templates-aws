import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const nlb = new aws.ec2.TrafficMirrorTarget("nlb", {
    description: "NLB target",
    networkLoadBalancerArn: aws_lb_lb.arn,
});
const eni = new aws.ec2.TrafficMirrorTarget("eni", {
    description: "ENI target",
    networkInterfaceId: aws_instance_test.primaryNetworkInterfaceId,
});

