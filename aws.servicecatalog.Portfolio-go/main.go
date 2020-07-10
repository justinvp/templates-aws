package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/servicecatalog"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicecatalog.NewPortfolio(ctx, "portfolio", &servicecatalog.PortfolioArgs{
			Description:  pulumi.String("List of my organizations apps"),
			ProviderName: pulumi.String("Brett"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

