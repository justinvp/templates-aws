package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.NewSnapshotSchedule(ctx, "_default", &redshift.SnapshotScheduleArgs{
			Definitions: pulumi.StringArray{
				pulumi.String("rate(12 hours)"),
			},
			Identifier: pulumi.String("tf-redshift-snapshot-schedule"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

