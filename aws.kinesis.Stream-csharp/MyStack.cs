using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testStream = new Aws.Kinesis.Stream("testStream", new Aws.Kinesis.StreamArgs
        {
            RetentionPeriod = 48,
            ShardCount = 1,
            ShardLevelMetrics = 
            {
                "IncomingBytes",
                "OutgoingBytes",
            },
            Tags = 
            {
                { "Environment", "test" },
            },
        });
    }

}

