package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewLogGroup(ctx, "yada", &cloudwatch.LogGroupArgs{
			Tags: pulumi.StringMap{
				"Application": pulumi.String("serviceA"),
				"Environment": pulumi.String("production"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

