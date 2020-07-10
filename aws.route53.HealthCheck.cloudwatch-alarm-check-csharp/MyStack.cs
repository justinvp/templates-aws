using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foobar = new Aws.CloudWatch.MetricAlarm("foobar", new Aws.CloudWatch.MetricAlarmArgs
        {
            AlarmDescription = "This metric monitors ec2 cpu utilization",
            ComparisonOperator = "GreaterThanOrEqualToThreshold",
            EvaluationPeriods = 2,
            MetricName = "CPUUtilization",
            Namespace = "AWS/EC2",
            Period = 120,
            Statistic = "Average",
            Threshold = 80,
        });
        var foo = new Aws.Route53.HealthCheck("foo", new Aws.Route53.HealthCheckArgs
        {
            CloudwatchAlarmName = foobar.Name,
            CloudwatchAlarmRegion = "us-west-2",
            InsufficientDataHealthStatus = "Healthy",
            Type = "CLOUDWATCH_METRIC",
        });
    }

}

