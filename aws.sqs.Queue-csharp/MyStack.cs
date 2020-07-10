using System.Collections.Generic;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var queue = new Aws.Sqs.Queue("queue", new Aws.Sqs.QueueArgs
        {
            DelaySeconds = 90,
            MaxMessageSize = 2048,
            MessageRetentionSeconds = 86400,
            ReceiveWaitTimeSeconds = 10,
            RedrivePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                { "deadLetterTargetArn", aws_sqs_queue.Queue_deadletter.Arn },
                { "maxReceiveCount", 4 },
            }),
            Tags = 
            {
                { "Environment", "production" },
            },
        });
    }

}

