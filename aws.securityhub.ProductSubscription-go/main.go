package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		current, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewProductSubscription(ctx, "exampleProductSubscription", &securityhub.ProductSubscriptionArgs{
			ProductArn: pulumi.String(fmt.Sprintf("%v%v%v", "arn:aws:securityhub:", current.Name, ":733251395267:product/alertlogic/althreatmanagement")),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_securityhub_account.example",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

