import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ses.IdentityNotificationTopic("test", {
    identity: aws_ses_domain_identity_example.domain,
    includeOriginalHeaders: true,
    notificationType: "Bounce",
    topicArn: aws_sns_topic_example.arn,
});

