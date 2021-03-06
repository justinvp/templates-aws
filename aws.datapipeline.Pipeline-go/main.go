package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/datapipeline"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datapipeline.NewPipeline(ctx, "_default", nil)
		if err != nil {
			return err
		}
		return nil
	})
}

