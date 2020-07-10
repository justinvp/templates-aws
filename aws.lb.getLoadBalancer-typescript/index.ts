import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const lbArn = config.get("lbArn") || "";
const lbName = config.get("lbName") || "";

const test = pulumi.output(aws.lb.getLoadBalancer({
    arn: lbArn,
    name: lbName,
}, { async: true }));

