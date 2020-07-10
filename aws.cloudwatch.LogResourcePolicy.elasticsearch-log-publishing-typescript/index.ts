import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const elasticsearch_log_publishing_policyPolicyDocument = pulumi.output(aws.iam.getPolicyDocument({
    statements: [{
        actions: [
            "logs:CreateLogStream",
            "logs:PutLogEvents",
            "logs:PutLogEventsBatch",
        ],
        principals: [{
            identifiers: ["es.amazonaws.com"],
            type: "Service",
        }],
        resources: ["arn:aws:logs:*"],
    }],
}, { async: true }));
const elasticsearch_log_publishing_policyLogResourcePolicy = new aws.cloudwatch.LogResourcePolicy("elasticsearch-log-publishing-policy", {
    policyDocument: elasticsearch_log_publishing_policyPolicyDocument.json,
    policyName: "elasticsearch-log-publishing-policy",
});

