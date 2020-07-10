using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var defaultInstance = new Aws.Rds.Instance("defaultInstance", new Aws.Rds.InstanceArgs
        {
            AllocatedStorage = 10,
            DbSubnetGroupName = "my_database_subnet_group",
            Engine = "mysql",
            EngineVersion = "5.6.17",
            InstanceClass = "db.t2.micro",
            Name = "mydb",
            ParameterGroupName = "default.mysql5.6",
            Password = "bar",
            Username = "foo",
        });
        var defaultTopic = new Aws.Sns.Topic("defaultTopic", new Aws.Sns.TopicArgs
        {
        });
        var defaultEventSubscription = new Aws.Rds.EventSubscription("defaultEventSubscription", new Aws.Rds.EventSubscriptionArgs
        {
            EventCategories = 
            {
                "availability",
                "deletion",
                "failover",
                "failure",
                "low storage",
                "maintenance",
                "notification",
                "read replica",
                "recovery",
                "restoration",
            },
            SnsTopic = defaultTopic.Arn,
            SourceIds = 
            {
                defaultInstance.Id,
            },
            SourceType = "db-instance",
        });
    }

}

