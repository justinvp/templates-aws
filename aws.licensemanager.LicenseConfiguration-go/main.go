package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/licensemanager"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := licensemanager.NewLicenseConfiguration(ctx, "example", &licensemanager.LicenseConfigurationArgs{
			Description:           pulumi.String("Example"),
			LicenseCount:          pulumi.Int(10),
			LicenseCountHardLimit: pulumi.Bool(true),
			LicenseCountingType:   pulumi.String("Socket"),
			LicenseRules: pulumi.StringArray{
				pulumi.String("#minimumSockets=2"),
			},
			Tags: pulumi.StringMap{
				"foo": pulumi.String("barr"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

