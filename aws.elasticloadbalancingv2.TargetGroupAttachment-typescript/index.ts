import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testTargetGroup = new aws.lb.TargetGroup("test", {});
const testInstance = new aws.ec2.Instance("test", {});
const testTargetGroupAttachment = new aws.lb.TargetGroupAttachment("test", {
    port: 80,
    targetGroupArn: testTargetGroup.arn,
    targetId: testInstance.id,
});

