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
            InsufficientDataActions = {},
            MetricName = "CPUUtilization",
            Namespace = "AWS/EC2",
            Period = 120,
            Statistic = "Average",
            Threshold = 80,
        });
    }

}

