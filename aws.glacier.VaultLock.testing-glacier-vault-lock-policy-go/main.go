package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glacier"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleVault, err := glacier.NewVault(ctx, "exampleVault", nil)
		if err != nil {
			return err
		}
		_, err = glacier.NewVaultLock(ctx, "exampleVaultLock", &glacier.VaultLockArgs{
			CompleteLock: pulumi.Bool(false),
			Policy: examplePolicyDocument.ApplyT(func(examplePolicyDocument iam.GetPolicyDocumentResult) (string, error) {
				return examplePolicyDocument.Json, nil
			}).(pulumi.StringOutput),
			VaultName: exampleVault.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

