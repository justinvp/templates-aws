import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const lbName = config.get("lbName") || "";

const test = pulumi.output(aws.elb.getLoadBalancer({
    name: lbName,
}, { async: true }));

