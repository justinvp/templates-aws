package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/pinpoint"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := pinpoint.NewApp(ctx, "example", &pinpoint.AppArgs{
			Limits: &pinpoint.AppLimitsArgs{
				MaximumDuration: pulumi.Int(600),
			},
			QuietTime: &pinpoint.AppQuietTimeArgs{
				End:   pulumi.String("06:00"),
				Start: pulumi.String("00:00"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

