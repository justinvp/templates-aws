using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Dms.EventSubscription("example", new Aws.Dms.EventSubscriptionArgs
        {
            Enabled = true,
            EventCategories = 
            {
                "creation",
                "failure",
            },
            SnsTopicArn = aws_sns_topic.Example.Arn,
            SourceIds = 
            {
                aws_dms_replication_task.Example.Replication_task_id,
            },
            SourceType = "replication-task",
            Tags = 
            {
                { "Name", "example" },
            },
        });
    }

}

