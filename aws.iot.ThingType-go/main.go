package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iot"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iot.NewThingType(ctx, "foo", nil)
		if err != nil {
			return err
		}
		return nil
	})
}

