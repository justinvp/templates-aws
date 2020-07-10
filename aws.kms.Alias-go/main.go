package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		key, err := kms.NewKey(ctx, "key", nil)
		if err != nil {
			return err
		}
		_, err = kms.NewAlias(ctx, "alias", &kms.AliasArgs{
			TargetKeyId: key.KeyId,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

