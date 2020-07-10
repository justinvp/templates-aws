package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foobar, err := cloudwatch.NewMetricAlarm(ctx, "foobar", &cloudwatch.MetricAlarmArgs{
			AlarmDescription:   pulumi.String("This metric monitors ec2 cpu utilization"),
			ComparisonOperator: pulumi.String("GreaterThanOrEqualToThreshold"),
			EvaluationPeriods:  pulumi.Int(2),
			MetricName:         pulumi.String("CPUUtilization"),
			Namespace:          pulumi.String("AWS/EC2"),
			Period:             pulumi.Int(120),
			Statistic:          pulumi.String("Average"),
			Threshold:          pulumi.Float64(80),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewHealthCheck(ctx, "foo", &route53.HealthCheckArgs{
			CloudwatchAlarmName:          foobar.Name,
			CloudwatchAlarmRegion:        pulumi.String("us-west-2"),
			InsufficientDataHealthStatus: pulumi.String("Healthy"),
			Type:                         pulumi.String("CLOUDWATCH_METRIC"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

