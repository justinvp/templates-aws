package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ipset, err := waf.NewIpSet(ctx, "ipset", &waf.IpSetArgs{
			IpSetDescriptors: waf.IpSetIpSetDescriptorArray{
				&waf.IpSetIpSetDescriptorArgs{
					Type:  pulumi.String("IPV4"),
					Value: pulumi.String("192.0.7.0/24"),
				},
			},
		})
		if err != nil {
			return err
		}
		wafrule, err := waf.NewRule(ctx, "wafrule", &waf.RuleArgs{
			MetricName: pulumi.String("tfWAFRule"),
			Predicates: waf.RulePredicateArray{
				&waf.RulePredicateArgs{
					DataId:  ipset.ID(),
					Negated: pulumi.Bool(false),
					Type:    pulumi.String("IPMatch"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_waf_ipset.ipset",
		}))
		if err != nil {
			return err
		}
		_, err = waf.NewWebAcl(ctx, "wafAcl", &waf.WebAclArgs{
			DefaultAction: &waf.WebAclDefaultActionArgs{
				Type: pulumi.String("ALLOW"),
			},
			MetricName: pulumi.String("tfWebACL"),
			Rules: waf.WebAclRuleArray{
				&waf.WebAclRuleArgs{
					Action: &waf.WebAclRuleActionArgs{
						Type: pulumi.String("BLOCK"),
					},
					Priority: pulumi.Int(1),
					RuleId:   wafrule.ID(),
					Type:     pulumi.String("REGULAR"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_waf_ipset.ipset",
			"aws_waf_rule.wafrule",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

