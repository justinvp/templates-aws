import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foobar = new aws.cloudwatch.MetricAlarm("foobar", {
    alarmDescription: "This metric monitors ec2 cpu utilization",
    comparisonOperator: "GreaterThanOrEqualToThreshold",
    evaluationPeriods: 2,
    metricName: "CPUUtilization",
    namespace: "AWS/EC2",
    period: 120,
    statistic: "Average",
    threshold: 80,
});
const foo = new aws.route53.HealthCheck("foo", {
    cloudwatchAlarmName: foobar.alarmName,
    cloudwatchAlarmRegion: "us-west-2",
    insufficientDataHealthStatus: "Healthy",
    type: "CLOUDWATCH_METRIC",
});

