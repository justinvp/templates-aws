package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/servicequotas"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "L-F678F1CE"
		_, err := servicequotas.LookupServiceQuota(ctx, &servicequotas.LookupServiceQuotaArgs{
			QuotaCode:   &opt0,
			ServiceCode: "vpc",
		}, nil)
		if err != nil {
			return err
		}
		opt1 := "VPCs per Region"
		_, err = servicequotas.LookupServiceQuota(ctx, &servicequotas.LookupServiceQuotaArgs{
			QuotaName:   &opt1,
			ServiceCode: "vpc",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

