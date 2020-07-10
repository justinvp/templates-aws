package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := route53.NewDelegationSet(ctx, "main", &route53.DelegationSetArgs{
			ReferenceName: pulumi.String("DynDNS"),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewZone(ctx, "primary", &route53.ZoneArgs{
			DelegationSetId: main.ID(),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewZone(ctx, "secondary", &route53.ZoneArgs{
			DelegationSetId: main.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

