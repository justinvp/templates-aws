import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.organizations.getOrganization({ async: true }));
const snsTopic = new aws.sns.Topic("sns_topic", {});
const snsTopicPolicyPolicyDocument = pulumi.all([example, snsTopic.arn]).apply(([example, arn]) => aws.iam.getPolicyDocument({
    statements: [{
        actions: [
            "SNS:Subscribe",
            "SNS:Publish",
        ],
        conditions: [{
            test: "StringEquals",
            values: [example.id],
            variable: "aws:PrincipalOrgID",
        }],
        effect: "Allow",
        principals: [{
            identifiers: ["*"],
            type: "AWS",
        }],
        resources: [arn],
    }],
}, { async: true }));
const snsTopicPolicyTopicPolicy = new aws.sns.TopicPolicy("sns_topic_policy", {
    arn: snsTopic.arn,
    policy: snsTopicPolicyPolicyDocument.json,
});

