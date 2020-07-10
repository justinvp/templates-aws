package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewRecord(ctx, "www_dev", &route53.RecordArgs{
			Name: pulumi.String("www"),
			Records: pulumi.StringArray{
				pulumi.String("dev.example.com"),
			},
			SetIdentifier: pulumi.String("dev"),
			Ttl:           pulumi.Int(5),
			Type:          pulumi.String("CNAME"),
			WeightedRoutingPolicies: route53.RecordWeightedRoutingPolicyArray{
				&route53.RecordWeightedRoutingPolicyArgs{
					Weight: pulumi.Int(10),
				},
			},
			ZoneId: pulumi.Any(aws_route53_zone.Primary.Zone_id),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "www_live", &route53.RecordArgs{
			Name: pulumi.String("www"),
			Records: pulumi.StringArray{
				pulumi.String("live.example.com"),
			},
			SetIdentifier: pulumi.String("live"),
			Ttl:           pulumi.Int(5),
			Type:          pulumi.String("CNAME"),
			WeightedRoutingPolicies: route53.RecordWeightedRoutingPolicyArray{
				&route53.RecordWeightedRoutingPolicyArgs{
					Weight: pulumi.Int(90),
				},
			},
			ZoneId: pulumi.Any(aws_route53_zone.Primary.Zone_id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

