import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const us_east_1 = new aws.Provider("us-east-1", {
    region: "us-east-1",
});
const exampleComZone = new aws.route53.Zone("example_com", {});
const awsRoute53ExampleCom = new aws.cloudwatch.LogGroup("aws_route53_example_com", {
    retentionInDays: 30,
}, { provider: us_east_1 });
const route53_query_logging_policyPolicyDocument = pulumi.output(aws.iam.getPolicyDocument({
    statements: [{
        actions: [
            "logs:CreateLogStream",
            "logs:PutLogEvents",
        ],
        principals: [{
            identifiers: ["route53.amazonaws.com"],
            type: "Service",
        }],
        resources: ["arn:aws:logs:*:*:log-group:/aws/route53/*"],
    }],
}, { async: true }));
const route53_query_logging_policyLogResourcePolicy = new aws.cloudwatch.LogResourcePolicy("route53-query-logging-policy", {
    policyDocument: route53_query_logging_policyPolicyDocument.json,
    policyName: "route53-query-logging-policy",
}, { provider: us_east_1 });
const exampleComQueryLog = new aws.route53.QueryLog("example_com", {
    cloudwatchLogGroupArn: awsRoute53ExampleCom.arn,
    zoneId: exampleComZone.zoneId,
}, { dependsOn: [route53_query_logging_policyLogResourcePolicy] });

