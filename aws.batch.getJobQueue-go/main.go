package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/batch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := batch.LookupJobQueue(ctx, &batch.LookupJobQueueArgs{
			Name: "tf-test-batch-job-queue",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

