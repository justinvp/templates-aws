import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Vpc("main", {
    cidrBlock: "10.0.0.0/16",
});
const ip_example = new aws.lb.TargetGroup("ip-example", {
    port: 80,
    protocol: "HTTP",
    targetType: "ip",
    vpcId: main.id,
});

