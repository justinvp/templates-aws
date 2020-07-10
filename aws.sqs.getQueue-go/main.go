package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sqs.LookupQueue(ctx, &sqs.LookupQueueArgs{
			Name: "queue",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

