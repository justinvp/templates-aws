package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glacier"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glacier.NewVaultLock(ctx, "example", &glacier.VaultLockArgs{
			CompleteLock: pulumi.Bool(true),
			Policy:       pulumi.Any(data.Aws_iam_policy_document.Example.Json),
			VaultName:    pulumi.Any(aws_glacier_vault.Example.Name),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

