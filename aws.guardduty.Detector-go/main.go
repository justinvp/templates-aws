package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/guardduty"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := guardduty.NewDetector(ctx, "myDetector", &guardduty.DetectorArgs{
			Enable: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

