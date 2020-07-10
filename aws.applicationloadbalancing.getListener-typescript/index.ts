import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const listenerArn = config.require("listenerArn");

const listener = pulumi.output(aws.lb.getListener({
    arn: listenerArn,
}, { async: true }));
const selected = pulumi.output(aws.lb.getLoadBalancer({
    name: "default-public",
}, { async: true }));
const selected443 = selected.apply(selected => aws.lb.getListener({
    loadBalancerArn: selected.arn!,
    port: 443,
}, { async: true }));

