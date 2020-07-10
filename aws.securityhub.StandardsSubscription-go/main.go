package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityhub.NewAccount(ctx, "example", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewStandardsSubscription(ctx, "cis", &securityhub.StandardsSubscriptionArgs{
			StandardsArn: pulumi.String("arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0"),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_securityhub_account.example",
		}))
		if err != nil {
			return err
		}
		_, err = securityhub.NewStandardsSubscription(ctx, "pci321", &securityhub.StandardsSubscriptionArgs{
			StandardsArn: pulumi.String("arn:aws:securityhub:us-east-1::standards/pci-dss/v/3.2.1"),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_securityhub_account.example",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

