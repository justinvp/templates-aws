package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "test.com."
		opt1 := true
		selected, err := route53.LookupZone(ctx, &route53.LookupZoneArgs{
			Name:        &opt0,
			PrivateZone: &opt1,
		}, nil)
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "www", &route53.RecordArgs{
			Name: pulumi.String(fmt.Sprintf("%v%v", "www.", selected.Name)),
			Records: pulumi.StringArray{
				pulumi.String("10.0.0.1"),
			},
			Ttl:    pulumi.Int(300),
			Type:   pulumi.String("A"),
			ZoneId: pulumi.String(selected.ZoneId),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

