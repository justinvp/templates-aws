package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleRule, err := waf.NewRule(ctx, "exampleRule", &waf.RuleArgs{
			MetricName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = waf.NewRuleGroup(ctx, "exampleRuleGroup", &waf.RuleGroupArgs{
			ActivatedRules: waf.RuleGroupActivatedRuleArray{
				&waf.RuleGroupActivatedRuleArgs{
					Action: &waf.RuleGroupActivatedRuleActionArgs{
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

