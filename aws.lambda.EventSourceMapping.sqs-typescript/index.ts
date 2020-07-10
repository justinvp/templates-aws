import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.EventSourceMapping("example", {
    eventSourceArn: aws_sqs_queue_sqs_queue_test.arn,
    functionName: aws_lambda_function_example.arn,
});

