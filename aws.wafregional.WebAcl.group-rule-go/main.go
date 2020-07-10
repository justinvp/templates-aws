package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafregional.NewWebAcl(ctx, "example", &wafregional.WebAclArgs{
			DefaultAction: &wafregional.WebAclDefaultActionArgs{
				Type: pulumi.String("ALLOW"),
			},
			MetricName: pulumi.String("example"),
			Rules: wafregional.WebAclRuleArray{
				&wafregional.WebAclRuleArgs{
					OverrideAction: &wafregional.WebAclRuleOverrideActionArgs{
						Type: pulumi.String("NONE"),
					},
					Priority: pulumi.Int(1),
					RuleId:   pulumi.Any(aws_wafregional_rule_group.Example.Id),
					Type:     pulumi.String("GROUP"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

