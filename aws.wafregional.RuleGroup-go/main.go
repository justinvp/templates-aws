package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleRule, err := wafregional.NewRule(ctx, "exampleRule", &wafregional.RuleArgs{
			MetricName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = wafregional.NewRuleGroup(ctx, "exampleRuleGroup", &wafregional.RuleGroupArgs{
			ActivatedRules: wafregional.RuleGroupActivatedRuleArray{
				&wafregional.RuleGroupActivatedRuleArgs{
					Action: &wafregional.RuleGroupActivatedRuleActionArgs{
						Type: pulumi.String("COUNT"),
					},
					Priority: pulumi.Int(50),
					RuleId:   exampleRule.ID(),
				},
			},
			MetricName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

