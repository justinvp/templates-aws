import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("example", {
    retentionInDays: 14,
});
const ad_log_policyPolicyDocument = exampleLogGroup.arn.apply(arn => aws.iam.getPolicyDocument({
    statements: [{
        actions: [
            "logs:CreateLogStream",
            "logs:PutLogEvents",
        ],
        effect: "Allow",
        principals: [{
            identifiers: ["ds.amazonaws.com"],
            type: "Service",
        }],
        resources: [arn],
    }],
}, { async: true }));
const ad_log_policyLogResourcePolicy = new aws.cloudwatch.LogResourcePolicy("ad-log-policy", {
    policyDocument: ad_log_policyPolicyDocument.json,
    policyName: "ad-log-policy",
});
const exampleLogService = new aws.directoryservice.LogService("example", {
    directoryId: aws_directory_service_directory_example.id,
    logGroupName: exampleLogGroup.name,
});

