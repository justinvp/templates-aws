package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewActiveReceiptRuleSet(ctx, "main", &ses.ActiveReceiptRuleSetArgs{
			RuleSetName: pulumi.String("primary-rules"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

