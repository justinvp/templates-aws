using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var clusterName = config.Get("clusterName") ?? "example";
        var exampleCluster = new Aws.Eks.Cluster("exampleCluster", new Aws.Eks.ClusterArgs
        {
            EnabledClusterLogTypes = 
            {
                "api",
                "audit",
            },
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_cloudwatch_log_group.example",
            },
        });
        var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup", new Aws.CloudWatch.LogGroupArgs
        {
            RetentionInDays = 7,
        });
    }

}

