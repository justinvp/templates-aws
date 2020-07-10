package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", &cloudwatch.LogGroupArgs{
			RetentionInDays: pulumi.Int(14),
		})
		if err != nil {
			return err
		}
		_, err = glue.NewJob(ctx, "exampleJob", &glue.JobArgs{
			DefaultArguments: pulumi.StringMap{
				"--continuous-log-logGroup":          exampleLogGroup.Name,
				"--enable-continuous-cloudwatch-log": pulumi.String("true"),
				"--enable-continuous-log-filter":     pulumi.String("true"),
				"--enable-metrics":                   pulumi.String(""),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

