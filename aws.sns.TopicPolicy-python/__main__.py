import pulumi
import pulumi_aws as aws

test = aws.sns.Topic("test")
sns_topic_policy = test.arn.apply(lambda arn: aws.iam.get_policy_document(policy_id="__default_policy_ID",
    statements=[{
        "actions": [
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
        "condition": [{
            "test": "StringEquals",
            "values": [var["account-id"]],
            "variable": "AWS:SourceOwner",
        }],
        "effect": "Allow",
        "principals": [{
            "identifiers": ["*"],
            "type": "AWS",
        }],
        "resources": [arn],
        "sid": "__default_statement_ID",
    }]))
default = aws.sns.TopicPolicy("default",
    arn=test.arn,
    policy=sns_topic_policy.json)

