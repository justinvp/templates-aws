package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iot"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iot.NewThing(ctx, "example", &iot.ThingArgs{
			Attributes: pulumi.StringMap{
				"First": pulumi.String("examplevalue"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

