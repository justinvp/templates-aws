using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
        });
        var topic = new Aws.Sns.Topic("topic", new Aws.Sns.TopicArgs
        {
            Policy = bucket.Arn.Apply(arn => @$"{{
    ""Version"":""2012-10-17"",
    ""Statement"":[{{
        ""Effect"": ""Allow"",
        ""Principal"": {{""AWS"":""*""}},
        ""Action"": ""SNS:Publish"",
        ""Resource"": ""arn:aws:sns:*:*:s3-event-notification-topic"",
        ""Condition"":{{
            ""ArnLike"":{{""aws:SourceArn"":""{arn}""}}
        }}
    }}]
}}

"),
        });
        var bucketNotification = new Aws.S3.BucketNotification("bucketNotification", new Aws.S3.BucketNotificationArgs
        {
            Bucket = bucket.Id,
            Topics = 
            {
                new Aws.S3.Inputs.BucketNotificationTopicArgs
                {
                    Events = 
                    {
                        "s3:ObjectCreated:*",
                    },
                    FilterSuffix = ".log",
                    TopicArn = topic.Arn,
                },
            },
        });
    }

}

