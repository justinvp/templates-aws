import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Vpc("main", {
    cidrBlock: "10.0.0.0/16",
});
const test = new aws.lb.TargetGroup("test", {
    port: 80,
    protocol: "HTTP",
    vpcId: main.id,
});

