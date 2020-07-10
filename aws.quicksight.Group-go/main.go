package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewGroup(ctx, "example", &quicksight.GroupArgs{
			GroupName: pulumi.String("tf-example"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

