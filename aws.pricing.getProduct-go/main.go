package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/pricing"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := pricing.GetProduct(ctx, &pricing.GetProductArgs{
			Filters: []pricing.GetProductFilter{
				pricing.GetProductFilter{
					Field: "instanceType",
					Value: "c5.xlarge",
				},
				pricing.GetProductFilter{
					Field: "operatingSystem",
					Value: "Linux",
				},
				pricing.GetProductFilter{
					Field: "location",
					Value: "US East (N. Virginia)",
				},
				pricing.GetProductFilter{
					Field: "preInstalledSw",
					Value: "NA",
				},
				pricing.GetProductFilter{
					Field: "licenseModel",
					Value: "No License required",
				},
				pricing.GetProductFilter{
					Field: "tenancy",
					Value: "Shared",
				},
				pricing.GetProductFilter{
					Field: "capacitystatus",
					Value: "Used",
				},
			},
			ServiceCode: "AmazonEC2",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

