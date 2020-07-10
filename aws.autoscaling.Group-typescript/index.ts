import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ec2.PlacementGroup("test", {
    strategy: "cluster",
});
const bar = new aws.autoscaling.Group("bar", {
    desiredCapacity: 4,
    forceDelete: true,
    healthCheckGracePeriod: 300,
    healthCheckType: "ELB",
    initialLifecycleHooks: [{
        defaultResult: "CONTINUE",
        heartbeatTimeout: 2000,
        lifecycleTransition: "autoscaling:EC2_INSTANCE_LAUNCHING",
        name: "foobar",
        notificationMetadata: `{
  "foo": "bar"
}
`,
        notificationTargetArn: "arn:aws:sqs:us-east-1:444455556666:queue1*",
        roleArn: "arn:aws:iam::123456789012:role/S3Access",
    }],
    launchConfiguration: aws_launch_configuration_foobar.name,
    maxSize: 5,
    minSize: 2,
    placementGroup: test.id,
    tags: [
        {
            key: "foo",
            propagateAtLaunch: true,
            value: "bar",
        },
        {
            key: "lorem",
            propagateAtLaunch: false,
            value: "ipsum",
        },
    ],
    vpcZoneIdentifiers: [
        aws_subnet_example1.id,
        aws_subnet_example2.id,
    ],
}, { timeouts: {
    delete: "15m",
} });

