import pulumi
import pulumi_aws as aws

console = aws.cloudwatch.EventRule("console",
    description="Capture each AWS Console Sign In",
    event_pattern="""{
  "detail-type": [
    "AWS Console Sign In via CloudTrail"
  ]
}

""")
aws_logins = aws.sns.Topic("awsLogins")
sns = aws.cloudwatch.EventTarget("sns",
    arn=aws_logins.arn,
    rule=console.name)
sns_topic_policy = aws_logins.arn.apply(lambda arn: aws.iam.get_policy_document(statements=[{
    "actions": ["SNS:Publish"],
    "effect": "Allow",
    "principals": [{
        "identifiers": ["events.amazonaws.com"],
        "type": "Service",
    }],
    "resources": [arn],
}]))
default = aws.sns.TopicPolicy("default",
    arn=aws_logins.arn,
    policy=sns_topic_policy.json)

