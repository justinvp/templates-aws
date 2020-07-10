using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup", new Aws.CloudWatch.LogGroupArgs
        {
            RetentionInDays = 14,
        });
        var exampleJob = new Aws.Glue.Job("exampleJob", new Aws.Glue.JobArgs
        {
            DefaultArguments = 
            {
                { "--continuous-log-logGroup", exampleLogGroup.Name },
                { "--enable-continuous-cloudwatch-log", "true" },
                { "--enable-continuous-log-filter", "true" },
                { "--enable-metrics", "" },
            },
        });
    }

}

