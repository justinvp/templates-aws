import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.sns.Topic("test", {});
const snsTopicPolicy = test.arn.apply(arn => aws.iam.getPolicyDocument({
    policyId: "__default_policy_ID",
    statements: [{
        actions: [
            "SNS:Subscribe",
            "SNS:SetTopicAttributes",
            "SNS:RemovePermission",
            "SNS:Receive",
            "SNS:Publish",
            "SNS:ListSubscriptionsByTopic",
            "SNS:GetTopicAttributes",
            "SNS:DeleteTopic",
            "SNS:AddPermission",
        ],
        conditions: [{
            test: "StringEquals",
            values: [var_account_id],
            variable: "AWS:SourceOwner",
        }],
        effect: "Allow",
        principals: [{
            identifiers: ["*"],
            type: "AWS",
        }],
        resources: [arn],
        sid: "__default_statement_ID",
    }],
}, { async: true }));
const defaultTopicPolicy = new aws.sns.TopicPolicy("default", {
    arn: test.arn,
    policy: snsTopicPolicy.json,
});

