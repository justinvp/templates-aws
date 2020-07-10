import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.VpcEndpointService("example", {
    acceptanceRequired: false,
    networkLoadBalancerArns: [aws_lb_example.arn],
});

