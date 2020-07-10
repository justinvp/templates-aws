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
		wafrule, err := wafregional.NewRule(ctx, "wafrule", &wafregional.RuleArgs{
			MetricName: pulumi.String("tfWAFRule"),
			Predicates: wafregional.RulePredicateArray{
				&wafregional.RulePredicateArgs{
					DataId:  ipset.ID(),
					Negated: pulumi.Bool(false),
					Type:    pulumi.String("IPMatch"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = wafregional.NewWebAcl(ctx, "wafacl", &wafregional.WebAclArgs{
			DefaultAction: &wafregional.WebAclDefaultActionArgs{
				Type: pulumi.String("ALLOW"),
			},
			MetricName: pulumi.String("tfWebACL"),
			Rules: wafregional.WebAclRuleArray{
				&wafregional.WebAclRuleArgs{
					Action: &wafregional.WebAclRuleActionArgs{
						Type: pulumi.String("BLOCK"),
					},
					Priority: pulumi.Int(1),
					RuleId:   wafrule.ID(),
					Type:     pulumi.String("REGULAR"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

