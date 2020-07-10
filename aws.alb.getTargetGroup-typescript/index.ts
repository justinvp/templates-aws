import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const lbTgArn = config.get("lbTgArn") || "";
const lbTgName = config.get("lbTgName") || "";

const test = pulumi.output(aws.lb.getTargetGroup({
    arn: lbTgArn,
    name: lbTgName,
}, { async: true }));

