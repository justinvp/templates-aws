package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
			Actions: glue.TriggerActionArray{
				&glue.TriggerActionArgs{
					JobName: pulumi.Any(aws_glue_job.Example.Name),
				},
			},
			Schedule: pulumi.String("cron(15 12 * * ? *)"),
			Type:     pulumi.String("SCHEDULED"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

