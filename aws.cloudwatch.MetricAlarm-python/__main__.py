import pulumi
import pulumi_aws as aws

foobar = aws.cloudwatch.MetricAlarm("foobar",
    alarm_description="This metric monitors ec2 cpu utilization",
    comparison_operator="GreaterThanOrEqualToThreshold",
    evaluation_periods="2",
    insufficient_data_actions=[],
    metric_name="CPUUtilization",
    namespace="AWS/EC2",
    period="120",
    statistic="Average",
    threshold="80")

