import pulumi
import pulumi_aws as aws

test = aws.ses.IdentityNotificationTopic("test",
    identity=aws_ses_domain_identity["example"]["domain"],
    include_original_headers=True,
    notification_type="Bounce",
    topic_arn=aws_sns_topic["example"]["arn"])

