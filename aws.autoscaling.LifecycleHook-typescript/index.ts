import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foobarGroup = new aws.autoscaling.Group("foobar", {
    availabilityZones: ["us-west-2a"],
    healthCheckType: "EC2",
    tags: [{
        key: "Foo",
        propagateAtLaunch: true,
        value: "foo-bar",
    }],
    terminationPolicies: ["OldestInstance"],
});
const foobarLifecycleHook = new aws.autoscaling.LifecycleHook("foobar", {
    autoscalingGroupName: foobarGroup.name,
    defaultResult: "CONTINUE",
    heartbeatTimeout: 2000,
    lifecycleTransition: "autoscaling:EC2_INSTANCE_LAUNCHING",
    notificationMetadata: `{
  "foo": "bar"
}
`,
    notificationTargetArn: "arn:aws:sqs:us-east-1:444455556666:queue1*",
    roleArn: "arn:aws:iam::123456789012:role/S3Access",
});

