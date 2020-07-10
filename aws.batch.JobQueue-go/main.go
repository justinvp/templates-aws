package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/batch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := batch.NewJobQueue(ctx, "testQueue", &batch.JobQueueArgs{
			ComputeEnvironments: pulumi.StringArray{
				pulumi.Any(aws_batch_compute_environment.Test_environment_1.Arn),
				pulumi.Any(aws_batch_compute_environment.Test_environment_2.Arn),
			},
			Priority: pulumi.Int(1),
			State:    pulumi.String("ENABLED"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

