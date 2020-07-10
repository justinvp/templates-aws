package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/servicequotas"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicequotas.NewServiceQuota(ctx, "example", &servicequotas.ServiceQuotaArgs{
			QuotaCode:   pulumi.String("L-F678F1CE"),
			ServiceCode: pulumi.String("vpc"),
			Value:       pulumi.Float64(75),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

