package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafv2.NewIpSet(ctx, "example", &wafv2.IpSetArgs{
			Addresses: pulumi.StringArray{
				pulumi.String("1.2.3.4/32"),
				pulumi.String("5.6.7.8/32"),
			},
			Description:      pulumi.String("Example IP set"),
			IpAddressVersion: pulumi.String("IPV4"),
			Scope:            pulumi.String("REGIONAL"),
			Tags: pulumi.StringMap{
				"Tag1": pulumi.String("Value1"),
				"Tag2": pulumi.String("Value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

