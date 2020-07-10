using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var defaultCluster = new Aws.RedShift.Cluster("defaultCluster", new Aws.RedShift.ClusterArgs
        {
            ClusterIdentifier = "default",
            DatabaseName = "default",
        });
        var defaultTopic = new Aws.Sns.Topic("defaultTopic", new Aws.Sns.TopicArgs
        {
        });
        var defaultEventSubscription = new Aws.RedShift.EventSubscription("defaultEventSubscription", new Aws.RedShift.EventSubscriptionArgs
        {
            EventCategories = 
            {
                "configuration",
                "management",
                "monitoring",
                "security",
            },
            Severity = "INFO",
            SnsTopicArn = defaultTopic.Arn,
            SourceIds = 
            {
                defaultCluster.Id,
            },
            SourceType = "cluster",
            Tags = 
            {
                { "Name", "default" },
            },
        });
    }

}

