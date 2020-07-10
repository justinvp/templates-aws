package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/globalaccelerator"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := globalaccelerator.NewAccelerator(ctx, "example", &globalaccelerator.AcceleratorArgs{
			Attributes: &globalaccelerator.AcceleratorAttributesArgs{
				FlowLogsEnabled:  pulumi.Bool(true),
				FlowLogsS3Bucket: pulumi.String("example-bucket"),
				FlowLogsS3Prefix: pulumi.String("flow-logs/"),
			},
			Enabled:       pulumi.Bool(true),
			IpAddressType: pulumi.String("IPV4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

