import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const console = new aws.cloudwatch.EventRule("console", {
    description: "Capture all EC2 scaling events",
    eventPattern: `{
  "source": [
    "aws.autoscaling"
  ],
  "detail-type": [
    "EC2 Instance Launch Successful",
    "EC2 Instance Terminate Successful",
    "EC2 Instance Launch Unsuccessful",
    "EC2 Instance Terminate Unsuccessful"
  ]
}
`,
});
const testStream = new aws.kinesis.Stream("test_stream", {
    shardCount: 1,
});
const yada = new aws.cloudwatch.EventTarget("yada", {
    arn: testStream.arn,
    rule: console.name,
    runCommandTargets: [
        {
            key: "tag:Name",
            values: ["FooBar"],
        },
        {
            key: "InstanceIds",
            values: ["i-162058cd308bffec2"],
        },
    ],
});

