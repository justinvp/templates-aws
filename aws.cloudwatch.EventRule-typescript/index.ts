import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const console = new aws.cloudwatch.EventRule("console", {
    description: "Capture each AWS Console Sign In",
    eventPattern: `{
  "detail-type": [
    "AWS Console Sign In via CloudTrail"
  ]
}
`,
});
const awsLogins = new aws.sns.Topic("aws_logins", {});
const sns = new aws.cloudwatch.EventTarget("sns", {
    arn: awsLogins.arn,
    rule: console.name,
});
const snsTopicPolicy = awsLogins.arn.apply(arn => aws.iam.getPolicyDocument({
    statements: [{
        actions: ["SNS:Publish"],
        effect: "Allow",
        principals: [{
            identifiers: ["events.amazonaws.com"],
            type: "Service",
        }],
        resources: [arn],
    }],
}, { async: true }));
const defaultTopicPolicy = new aws.sns.TopicPolicy("default", {
    arn: awsLogins.arn,
    policy: snsTopicPolicy.json,
});

