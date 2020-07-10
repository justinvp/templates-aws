package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := autoscaling.LookupGroup(ctx, &autoscaling.LookupGroupArgs{
			Name: "foo",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

