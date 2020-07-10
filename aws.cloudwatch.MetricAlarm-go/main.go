package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewMetricAlarm(ctx, "foobar", &cloudwatch.MetricAlarmArgs{
			AlarmDescription:        pulumi.String("This metric monitors ec2 cpu utilization"),
			ComparisonOperator:      pulumi.String("GreaterThanOrEqualToThreshold"),
			EvaluationPeriods:       pulumi.Int(2),
			InsufficientDataActions: []interface{}{},
			MetricName:              pulumi.String("CPUUtilization"),
			Namespace:               pulumi.String("AWS/EC2"),
			Period:                  pulumi.Int(120),
			Statistic:               pulumi.String("Average"),
			Threshold:               pulumi.Float64(80),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

