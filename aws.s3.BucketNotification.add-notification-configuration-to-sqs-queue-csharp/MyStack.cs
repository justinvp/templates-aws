using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
        });
        var queue = new Aws.Sqs.Queue("queue", new Aws.Sqs.QueueArgs
        {
            Policy = bucket.Arn.Apply(arn => @$"{{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Effect"": ""Allow"",
      ""Principal"": ""*"",
      ""Action"": ""sqs:SendMessage"",
	  ""Resource"": ""arn:aws:sqs:*:*:s3-event-notification-queue"",
      ""Condition"": {{
        ""ArnEquals"": {{ ""aws:SourceArn"": ""{arn}"" }}
      }}
    }}
  ]
}}

"),
        });
        var bucketNotification = new Aws.S3.BucketNotification("bucketNotification", new Aws.S3.BucketNotificationArgs
        {
            Bucket = bucket.Id,
            Queues = 
            {
                new Aws.S3.Inputs.BucketNotificationQueueArgs
                {
                    Events = 
                    {
                        "s3:ObjectCreated:*",
                    },
                    FilterSuffix = ".log",
                    QueueArn = queue.Arn,
                },
            },
        });
    }

}

