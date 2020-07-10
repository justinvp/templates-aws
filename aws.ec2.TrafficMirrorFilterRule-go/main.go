package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		filter, err := ec2.NewTrafficMirrorFilter(ctx, "filter", &ec2.TrafficMirrorFilterArgs{
			Description: pulumi.String("traffic mirror filter - example"),
			NetworkServices: pulumi.StringArray{
				pulumi.String("amazon-dns"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewTrafficMirrorFilterRule(ctx, "ruleout", &ec2.TrafficMirrorFilterRuleArgs{
			Description:           pulumi.String("test rule"),
			DestinationCidrBlock:  pulumi.String("10.0.0.0/8"),
			RuleAction:            pulumi.String("accept"),
			RuleNumber:            pulumi.Int(1),
			SourceCidrBlock:       pulumi.String("10.0.0.0/8"),
			TrafficDirection:      pulumi.String("egress"),
			TrafficMirrorFilterId: filter.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewTrafficMirrorFilterRule(ctx, "rulein", &ec2.TrafficMirrorFilterRuleArgs{
			Description:          pulumi.String("test rule"),
			DestinationCidrBlock: pulumi.String("10.0.0.0/8"),
			DestinationPortRange: &ec2.TrafficMirrorFilterRuleDestinationPortRangeArgs{
				FromPort: pulumi.Int(22),
				ToPort:   pulumi.Int(53),
			},
			Protocol:        pulumi.Int(6),
			RuleAction:      pulumi.String("accept"),
			RuleNumber:      pulumi.Int(1),
			SourceCidrBlock: pulumi.String("10.0.0.0/8"),
			SourcePortRange: &ec2.TrafficMirrorFilterRuleSourcePortRangeArgs{
				FromPort: pulumi.Int(0),
				ToPort:   pulumi.Int(10),
			},
			TrafficDirection:      pulumi.String("ingress"),
			TrafficMirrorFilterId: filter.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

