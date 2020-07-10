package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "example"
		_, err := secretsmanager.LookupSecretVersion(ctx, &secretsmanager.LookupSecretVersionArgs{
			SecretId:     data.Aws_secretsmanager_secret.Example.Id,
			VersionStage: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

