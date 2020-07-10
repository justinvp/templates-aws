import pulumi
import pulumi_aws as aws

example = aws.organizations.get_organization()
sns_topic = aws.sns.Topic("snsTopic")
sns_topic_policy_policy_document = sns_topic.arn.apply(lambda arn: aws.iam.get_policy_document(statements=[{
    "actions": [
        "SNS:Subscribe",
        "SNS:Publish",
    ],
    "condition": [{
        "test": "StringEquals",
        "values": [example.id],
        "variable": "aws:PrincipalOrgID",
    }],
    "effect": "Allow",
    "principals": [{
        "identifiers": ["*"],
        "type": "AWS",
    }],
    "resources": [arn],
}]))
sns_topic_policy_topic_policy = aws.sns.TopicPolicy("snsTopicPolicyTopicPolicy",
    arn=sns_topic.arn,
    policy=sns_topic_policy_policy_document.json)

