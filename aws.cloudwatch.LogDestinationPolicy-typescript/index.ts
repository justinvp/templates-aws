import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testDestination = new aws.cloudwatch.LogDestination("test_destination", {
    roleArn: aws_iam_role_iam_for_cloudwatch.arn,
    targetArn: aws_kinesis_stream_kinesis_for_cloudwatch.arn,
});
const testDestinationPolicyPolicyDocument = testDestination.arn.apply(arn => aws.iam.getPolicyDocument({
    statements: [{
        actions: ["logs:PutSubscriptionFilter"],
        effect: "Allow",
        principals: [{
            identifiers: ["123456789012"],
            type: "AWS",
        }],
        resources: [arn],
    }],
}, { async: true }));
const testDestinationPolicyLogDestinationPolicy = new aws.cloudwatch.LogDestinationPolicy("test_destination_policy", {
    accessPolicy: testDestinationPolicyPolicyDocument.json,
    destinationName: testDestination.name,
});

