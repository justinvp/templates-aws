import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foobar = new aws.cloudwatch.MetricAlarm("foobar", {
    alarmDescription: "This metric monitors ec2 cpu utilization",
    comparisonOperator: "GreaterThanOrEqualToThreshold",
    evaluationPeriods: 2,
    insufficientDataActions: [],
    metricName: "CPUUtilization",
    namespace: "AWS/EC2",
    period: 120,
    statistic: "Average",
    threshold: 80,
});

