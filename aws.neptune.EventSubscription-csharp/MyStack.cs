using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var defaultCluster = new Aws.Neptune.Cluster("defaultCluster", new Aws.Neptune.ClusterArgs
        {
            ApplyImmediately = true,
            BackupRetentionPeriod = 5,
            ClusterIdentifier = "neptune-cluster-demo",
            Engine = "neptune",
            IamDatabaseAuthenticationEnabled = true,
            PreferredBackupWindow = "07:00-09:00",
            SkipFinalSnapshot = true,
        });
        var example = new Aws.Neptune.ClusterInstance("example", new Aws.Neptune.ClusterInstanceArgs
        {
            ApplyImmediately = true,
            ClusterIdentifier = defaultCluster.Id,
            Engine = "neptune",
            InstanceClass = "db.r4.large",
        });
        var defaultTopic = new Aws.Sns.Topic("defaultTopic", new Aws.Sns.TopicArgs
        {
        });
        var defaultEventSubscription = new Aws.Neptune.EventSubscription("defaultEventSubscription", new Aws.Neptune.EventSubscriptionArgs
        {
            EventCategories = 
            {
                "maintenance",
                "availability",
                "creation",
                "backup",
                "restoration",
                "recovery",
                "deletion",
                "failover",
                "failure",
                "notification",
                "configuration change",
                "read replica",
            },
            SnsTopicArn = defaultTopic.Arn,
            SourceIds = 
            {
                example.Id,
            },
            SourceType = "db-instance",
            Tags = 
            {
                { "env", "test" },
            },
        });
    }

}

