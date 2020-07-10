package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewReceiptRuleSet(ctx, "main", &ses.ReceiptRuleSetArgs{
			RuleSetName: pulumi.String("primary-rules"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

