package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewReceiptRule(ctx, "store", &ses.ReceiptRuleArgs{
			AddHeaderActions: ses.ReceiptRuleAddHeaderActionArray{
				&ses.ReceiptRuleAddHeaderActionArgs{
					HeaderName:  pulumi.String("Custom-Header"),
					HeaderValue: pulumi.String("Added by SES"),
					Position:    pulumi.Int(1),
				},
			},
			Enabled: pulumi.Bool(true),
			Recipients: pulumi.StringArray{
				pulumi.String("karen@example.com"),
			},
			RuleSetName: pulumi.String("default-rule-set"),
			S3Actions: ses.ReceiptRuleS3ActionArray{
				&ses.ReceiptRuleS3ActionArgs{
					BucketName: pulumi.String("emails"),
					Position:   pulumi.Int(2),
				},
			},
			ScanEnabled: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

