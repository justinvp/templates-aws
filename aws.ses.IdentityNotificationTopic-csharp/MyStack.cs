using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.Ses.IdentityNotificationTopic("test", new Aws.Ses.IdentityNotificationTopicArgs
        {
            Identity = aws_ses_domain_identity.Example.Domain,
            IncludeOriginalHeaders = true,
            NotificationType = "Bounce",
            TopicArn = aws_sns_topic.Example.Arn,
        });
    }

}

