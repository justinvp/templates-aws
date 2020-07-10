import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucket = new aws.s3.Bucket("example", {});
const exampleFlowLog = new aws.ec2.FlowLog("example", {
    logDestination: exampleBucket.arn,
    logDestinationType: "s3",
    trafficType: "ALL",
    vpcId: aws_vpc_example.id,
});

