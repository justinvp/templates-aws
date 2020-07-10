package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/guardduty"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDetector, err := guardduty.NewDetector(ctx, "exampleDetector", &guardduty.DetectorArgs{
			Enable: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = guardduty.NewOrganizationConfiguration(ctx, "exampleOrganizationConfiguration", &guardduty.OrganizationConfigurationArgs{
			AutoEnable: pulumi.Bool(true),
			DetectorId: exampleDetector.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

