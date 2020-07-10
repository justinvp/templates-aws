package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ebs.NewVolume(ctx, "example", &ebs.VolumeArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			Size:             pulumi.Int(40),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

