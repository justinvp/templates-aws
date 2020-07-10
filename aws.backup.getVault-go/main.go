package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.LookupVault(ctx, &backup.LookupVaultArgs{
			Name: "example_backup_vault",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

