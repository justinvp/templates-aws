package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		dada, err := cloudwatch.NewLogGroup(ctx, "dada", nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogMetricFilter(ctx, "yada", &cloudwatch.LogMetricFilterArgs{
			LogGroupName: dada.Name,
			MetricTransformation: &cloudwatch.LogMetricFilterMetricTransformationArgs{
				Name:      pulumi.String("EventCount"),
				Namespace: pulumi.String("YourNamespace"),
				Value:     pulumi.String("1"),
			},
			Pattern: pulumi.String(""),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

