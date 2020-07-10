package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kms.NewKey(ctx, "key", &kms.KeyArgs{
			DeletionWindowInDays: pulumi.Int(10),
			Description:          pulumi.String("KMS key 1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

