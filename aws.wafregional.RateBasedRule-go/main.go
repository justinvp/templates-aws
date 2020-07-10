package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ipset, err := wafregional.NewIpSet(ctx, "ipset", &wafregional.IpSetArgs{
			IpSetDescriptors: wafregional.IpSetIpSetDescriptorArray{
				&wafregional.IpSetIpSetDescriptorArgs{
					Type:  pulumi.String("IPV4"),
					Value: pulumi.String("192.0.7.0/24"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = wafregional.NewRateBasedRule(ctx, "wafrule", &wafregional.RateBasedRuleArgs{
			MetricName: pulumi.String("tfWAFRule"),
			Predicates: wafregional.RateBasedRulePredicateArray{
				&wafregional.RateBasedRulePredicateArgs{
					DataId:  ipset.ID(),
					Negated: pulumi.Bool(false),
					Type:    pulumi.String("IPMatch"),
				},
			},
			RateKey:   pulumi.String("IP"),
			RateLimit: pulumi.Int(100),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_wafregional_ipset.ipset",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

