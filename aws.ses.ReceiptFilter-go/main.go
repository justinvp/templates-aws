package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewReceiptFilter(ctx, "filter", &ses.ReceiptFilterArgs{
			Cidr:   pulumi.String("10.10.10.10"),
			Policy: pulumi.String("Block"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

