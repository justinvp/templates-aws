package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		yada, err := cloudwatch.NewLogGroup(ctx, "yada", nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogStream(ctx, "foo", &cloudwatch.LogStreamArgs{
			LogGroupName: yada.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

