package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elb"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := elb.GetHostedZoneId(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "www", &route53.RecordArgs{
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					EvaluateTargetHealth: pulumi.Bool(true),
					Name:                 pulumi.Any(aws_elb.Main.Dns_name),
					ZoneId:               pulumi.String(main.Id),
				},
			},
			Name:   pulumi.String("example.com"),
			Type:   pulumi.String("A"),
			ZoneId: pulumi.Any(aws_route53_zone.Primary.Zone_id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

