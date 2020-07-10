import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.EventSourceMapping("example", {
    eventSourceArn: aws_kinesis_stream_example.arn,
    functionName: aws_lambda_function_example.arn,
    startingPosition: "LATEST",
});

